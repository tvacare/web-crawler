package main

import (
	"fmt"
)

func crawler() []Paper {
	// func crawler(cn chan []Paper) {
	// Get all papers URL
	papersLinks := getPapersLinks()
	papers := make([]Paper, 0)

	//  Get slice of papers and remove papers without name or marketValue
	for _, link := range papersLinks {
		// !!Evitei usar Go Routines porque a API estava retornando muitos timed outs
		// go getInformation(link, cn)
		p := getInformation(link)
		if p.name != "" || p.marketValue != 0 {
			papers = append(papers, p)
		}
	}

	// Sort papers by marketValue
	marketValue := func(p1, p2 *Paper) bool {
		return p1.marketValue < p2.marketValue
	}
	// Sort papers by marketValue in descending order
	descMarketValue := func(p1, p2 *Paper) bool {
		return marketValue(p2, p1)
	}

	By(descMarketValue).Sort(papers)
	fmt.Print("\n\n\n\n\n")
	return papers
}
