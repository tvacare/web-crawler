package main

import (
	"fmt"

	"github.com/tvacare/web-crawler/database"
	"github.com/tvacare/web-crawler/util"
)

var (
	// LINK with the data about bovespa papers
	LINK = util.GetenvRequired("LINK")

	// URLBASE for fetching paper info
	URLBASE = util.GetenvRequired("URLBASE")
)

func init() {
	// Create database connection
	_ = database.NewDB()
}

func main() {

	// Bovespa papers crawler
	papers := crawler()

	// Show 10 biggest market values
	fmt.Printf(`
	As 10 maiores empresas em capitais de mercado são:`)

	for i := 0; i < 10; i++ {
		database.CreatePaper(papers[i])

		fmt.Printf(`
		%d - 
		Sigla: %v
		Empresa: - %v
		Valor de Mercado: - $%v
		Variaçao Diária: - %v
		URL: - %v`, i+1, papers[i].Name, papers[i].Company, papers[i].MarketValue,
			papers[i].DailyVariation, papers[i].URL)
	}
	fmt.Printf(`

	--> Total de Resultados: %d`, len(papers))
}
