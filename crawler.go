package main

import (
	"fmt"

	"github.com/tvacare/web-crawler/paper"
	"github.com/tvacare/web-crawler/website"
)

func crawler() []paper.Paper {

	// Get all papers URL
	papersLinks := website.GetLinks(LINK, URLBASE)
	papers := make([]paper.Paper, 0)

	//  cn := make(chan Paper)
	//  Get slice of papers and remove papers without name or marketValue
	for _, link := range papersLinks {

		// ***!!Evitei usar Go Routines porque a API estava retornando muitos timed outs
		// var p paper.Paper
		p := paper.GetInformation(link)
		// time.Sleep(100 * time.Millisecond)
		// go getInformation(link, cn)

		if p.Name != "" || p.MarketValue != 0 {
			papers = append(papers, p)
		}
	}

	paper.By(paper.DescMarketValue).Sort(papers)
	fmt.Print("\n\n\n\n\n")
	return papers
}
