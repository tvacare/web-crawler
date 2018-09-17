package website

import (
	"log"
	"net/http"

	"github.com/tvacare/web-crawler/util"
)

// GetLinks Get all links from base site
// @Returning []string links
func GetLinks(link string, urlbase string) []string {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	websiteParsed, _ := util.MatchPage(resp, `<td><a.*?href="?(=?.*?)\s?"?>.*?</a></td>`)
	urls := make([]string, 0)

	for _, w := range websiteParsed {
		urls = append(urls, urlbase+string(w[1]))
	}
	return urls
}
