package resources

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/contrib/testutils"
	"testing"
	"time"
)

func TestUnmarshalJSON(t *testing.T) {
	input := []byte("\"2017-02-28T08:54:42.123-03:00\"")
	var date DateTime8601
	date.UnmarshalJSON(input)

	expectedTime := time.Date(2017, time.February, 28, 8, 54, 42, 123*1000000, time.FixedZone("-03", -3*60*60))

	testutils.Assert(t, date.Time.Equal(expectedTime), "Times are not equal")
}

func TestMarshalJSON(t *testing.T) {
	expected := []byte("\"2017-02-28T08:54:42.123-03:00\"")

	date := DateTime8601{Time: time.Date(2017, time.February, 28, 8, 54, 42, 123*1000000, time.FixedZone("-03", -3*60*60))}
	result, err := date.MarshalJSON()

	testutils.Ok(t, err)
	testutils.Equals(t, expected, result)
}

func TestFromStringFactoryReturnsExpectedValue(t *testing.T) {
	input := "2017-02-28T16:54:42.123-03:00"

	result, err := DateTime8601FromString(input)
	testutils.Ok(t, err)

	expected := DateTime8601{}
	err = expected.FromString(input)

	testutils.Equals(t, expected, result)
}

func TestDateTimeFromStringMethodReturnsExpectedValue(t *testing.T) {
	input := "2017-02-28T08:54:42.123-03:00"

	expectedTime, err := time.Parse(ISO8601Layout, input)
	testutils.Ok(t, err)

	expected := DateTime8601{Time: expectedTime}

	result := DateTime8601{}
	err = result.FromString(input)
	testutils.Ok(t, err)

	testutils.Equals(t, expected, result)
}
