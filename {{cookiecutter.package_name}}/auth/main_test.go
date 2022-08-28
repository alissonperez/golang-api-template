package auth

import (
	"fmt"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/config"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/contrib/testutils"
	"net/http"
	"testing"
)

var (
	validToken   string
	expiredToken string
	validKey     string
)

func init() {
	validKey = "my-secret-key-for-tests"
	validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6NDJ9.69FwyKQqg_L9DPzYcsWY-U4B5ZUxFZSwhKzopWfKtTw"
	expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6NDIsImV4cCI6MTYxMjI3OTk0M30.YcaODJ19jO2jfaR12j8zN7JBWBG8pTcTeKKN3Br4Qq8"
}

func createAuth(t *testing.T, jwtkey, token string) (*http.Request, Auth) {
	localConfig := config.CreateConfigFromMap(map[string]interface{}{
		"jwt_key": jwtkey,
	})
	auth := CreateAuth(localConfig)

	request, err := http.NewRequest("POST", "http://my-api.com/foo/bar", http.NoBody)
	testutils.Ok(t, err)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	return request, auth
}

func TestFromRequestReturnExpectedClaim(t *testing.T) {
	request, auth := createAuth(
		t,
		validKey,
		validToken,
	)

	result, err := auth.FromRequest(request)
	testutils.Ok(t, err)

	testutils.Equals(t, result.ClientId, 42)
}

func TestFromRequestWithEmptyTokenReturnUnexpectedTokenError(t *testing.T) {
	request, auth := createAuth(
		t,
		validKey,
		"",
	)

	result, err := auth.FromRequest(request)
	testutils.Equals(t, fmt.Errorf("Invalid Authorization header"), err)
	testutils.Equals(t, (*Claims)(nil), result)
}

func TestFromRequestWithInvalidTokenReturnMalformedTokenError(t *testing.T) {
	request, auth := createAuth(
		t,
		validKey,
		"foooobaaarrr",
	)

	result, err := auth.FromRequest(request)
	testutils.Equals(t, fmt.Errorf("Malformed token"), err)
	testutils.Equals(t, (*Claims)(nil), result)
}

func TestFromRequestWithInvalidSignatureReturnMalformedTokenError(t *testing.T) {
	request, auth := createAuth(
		t,
		"any-other-key",
		validToken,
	)

	result, err := auth.FromRequest(request)
	testutils.Equals(t, fmt.Errorf("Unexpected token"), err)
	testutils.Equals(t, (*Claims)(nil), result)
}

func TestFromRequestWithExpiredTokenReturnExpiredTokenError(t *testing.T) {
	request, auth := createAuth(
		t,
		validKey,
		expiredToken,
	)

	result, err := auth.FromRequest(request)
	testutils.Equals(t, fmt.Errorf("Expired token"), err)
	testutils.Equals(t, (*Claims)(nil), result)
}

func TestFromRequestWithNoAuthorizationReturnExpectedError(t *testing.T) {
	request, auth := createAuth(
		t,
		validKey,
		validToken,
	)

	request.Header.Del("Authorization")

	result, err := auth.FromRequest(request)
	testutils.Equals(t, fmt.Errorf("Invalid Authorization header"), err)
	testutils.Equals(t, (*Claims)(nil), result)
}

func TestFromRequestWithEmptyBearerReturnExpectedError(t *testing.T) {
	request, auth := createAuth(
		t,
		validKey,
		validToken,
	)

	request.Header.Set("Authorization", "Bearer")

	result, err := auth.FromRequest(request)
	testutils.Equals(t, fmt.Errorf("Invalid Authorization header"), err)
	testutils.Equals(t, (*Claims)(nil), result)
}

func TestFromRequestWithInvalidAuthorizationHeaderReturnExpectedError(t *testing.T) {
	request, auth := createAuth(
		t,
		validKey,
		validToken,
	)

	request.Header.Set("Authorization", "foooobaaarr")

	result, err := auth.FromRequest(request)
	testutils.Equals(t, fmt.Errorf("Authorization header must begging with 'Bearer'"), err)
	testutils.Equals(t, (*Claims)(nil), result)
}
