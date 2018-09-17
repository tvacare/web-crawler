package paper

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
	Name           string
	Company        string
	MarketValue    float64
	DailyVariation string
	URL            string
}

// GetInformation Get Paper information such as: name, company, marketValue, dailyVariation
// @Returning [] Paper {struct} || error err
// func getInformation(url string, c chan Paper) {
func GetInformation(url string) Paper {
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("%s - %v\n", url, req.Response.Status)
			return errors.New(req.Response.Status)
		},
	}

	p := &Paper{}

	resp, err := httpClient.Get(url)
	if err != nil {
		p.URL = url
		// c <- *p
		return *p
	}

	paperParsed, _ := util.MatchPage(resp, `<td.*?>(<span.*?>(.*?)</span>)?</.*td>`)
	paperVariation, _ := util.MatchPage(resp, `<td.*?>(<span.*?>(<font.*?>)?(.*?)(</.*font>)?</span>)?</.*td>`)

	p = p.filterInfo(paperParsed)
	p = p.filterVariation(paperVariation)

	p.URL = url

	fmt.Printf("Papel: %v\n", p)
	// c <- *p
	defer resp.Body.Close()
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
					p.DailyVariation = paperParsed[i+1][3]
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
				p.Name = paperParsed[i+1][2]
				break
			case "Empresa":
				p.Company = paperParsed[i+1][2]
				break
			case "Valor de mercado":
				s, _ := strconv.ParseFloat(strings.Replace(paperParsed[i+1][2], ".", "", -1), 64)
				p.MarketValue = s
				isDone = true
				break
			}
		}
	}
	return p
}
