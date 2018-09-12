package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/tvacare/web-crawler/util"
)

var properties = []string{"Papel", "Empresa", "Valor de mercado"}

// Paper represents a company Paper
type Paper struct {
	name           string
	company        string
	marketValue    float64
	dailyVariation string
	url            string
}

// Get Paper information such as: name, company, marketValue, dailyVariation
// @Returning [] Paper {struct} || error err
// func getInformation(url string, c chan Paper) {
func getInformation(url string) Paper {
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// return http.ErrUseLastResponse
			return errors.New(req.Response.Status)
		},
	}

	p := &Paper{}

	resp, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(err)
		p.url = url
		// c <- *p
		return *p
	}

	paperParsed, _ := util.MatchPage(resp, `<td.*?>(<span.*?>(.*?)</span>)?</.*td>`)
	paperVariation, _ := util.MatchPage(resp, `<td.*?>(<span.*?>(<font.*?>)?(.*?)(</.*font>)?</span>)?</.*td>`)

	p = p.filterInfo(paperParsed)
	p = p.filterVariation(paperVariation)

	p.url = url

	defer resp.Body.Close()
	// c <- *p
	fmt.Printf("Papel: %v\n", p)
	return *p
}

// filterVariation filters Paper variation
// @Returning [] Paper {struct}
func (p *Paper) filterVariation(paperParsed [][]string) *Paper {
	isDone := false
	for i, pp := range paperParsed {
		contain, prop := util.SliceContains(pp[3], []string{"Dia"})

		// Break for-loop if Paper data is fullfilled
		if isDone {
			return p
		}

		if contain {
			switch prop {
			case "Dia":
				if strings.Contains(paperParsed[i+1][3], "%") {
					p.dailyVariation = paperParsed[i+1][3]
					isDone = true
					break
				}
			}
		}
	}

	return p
}

// filterInfo filters Paper info
// @Returning [] Paper {struct}
func (p *Paper) filterInfo(paperParsed [][]string) *Paper {
	isDone := false
	for i, pp := range paperParsed {
		contain, prop := util.SliceContains(pp[2], properties)

		// Break for-loop if Paper data is fullfilled
		if isDone {
			return p
		}

		if contain {
			switch prop {
			case "Papel":
				p.name = paperParsed[i+1][2]
				break
			case "Empresa":
				p.company = paperParsed[i+1][2]
				break
			case "Valor de mercado":
				s, _ := strconv.ParseFloat(strings.Replace(paperParsed[i+1][2], ".", "", -1), 64)
				p.marketValue = s
				isDone = true
				break
			}
		}
	}
	return p
}

// Get all Papers from base site
// @Returning [] string PapersUrls
func getPapersLinks() []string {
	resp, e := http.Get(LINK)
	if e != nil {
		fmt.Println("Err", e)
	}

	papers, _ := util.MatchPage(resp, `<td><a.*?href="?(=?.*?)\s?"?>.*?</a></td>`)
	papersUrls := make([]string, 0)

	for _, Paper := range papers {
		papersUrls = append(papersUrls, URLBASE+string(Paper[1]))
	}
	return papersUrls
}
