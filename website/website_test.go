package website_test

import (
	"net/url"
	"testing"

	"github.com/tvacare/web-crawler/website"
)

const (
	LINK    = "https://www.fundamentus.com.br/detalhes.php"
	URLBASE = "http://www.fundamentus.com.br/"
)

func TestGetLinks(t *testing.T) {
	links := website.GetLinks(LINK, URLBASE)

	for _, l := range links {
		_, err := url.ParseRequestURI(l)
		if err != nil {
			t.Errorf("URL request might be invalid %v", err)
		}
	}
}
