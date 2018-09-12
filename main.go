package main

import (
	"fmt"
)

const (
	// LINK with the data about bovespa papers
	LINK = "https://www.fundamentus.com.br/detalhes.php"

	// URLBASE for fetching paper info
	URLBASE = "http://www.fundamentus.com.br/"
)

func main() {

	// Bovespa papers crawler
	papers := crawler()

	// Show 10 biggest market values
	fmt.Printf(`
	As 10 maiores empresas em capitais de mercado são:`)

	for i := 0; i < 10; i++ {
		fmt.Printf(`
		%d - 
		Sigla: %v
		Empresa: - %v
		Valor de Mercado: - $%v
		Variaçao Diária: - %v
		URL: - %v`, i+1, papers[i].name, papers[i].company, papers[i].marketValue,
			papers[i].dailyVariation, papers[i].url)
	}
	fmt.Printf(`

	--> Total de Resultados: %d`, len(papers))
}
