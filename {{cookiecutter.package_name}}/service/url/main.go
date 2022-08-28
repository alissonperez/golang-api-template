package url

import (
	"github.com/gorilla/mux"
	"log"
	"net/url"
)

type UrlService struct {
	router *mux.Router
}

func handleRouterUrl(urlObj *url.URL, err error) string {
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return urlObj.Path
}

func (s UrlService) Hello(name string) string {
	localUrl, err := s.router.Get("Hello").URL("myName", name)
	return handleRouterUrl(localUrl, err)
}

func NewService(router *mux.Router) UrlService {
	return UrlService{router: router}
}
