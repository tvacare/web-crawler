package database

import (
	"database/sql"
	"log"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/tvacare/web-crawler/paper"
	"github.com/tvacare/web-crawler/util"
)

var (
	// DBCon is the connection handle for the database
	DBCon *sql.DB

	mysqlHost     = util.GetenvRequired("MYSQL_HOST")
	mysqlUser     = util.GetenvRequired("MYSQL_USER")
	mysqlPassword = util.GetenvRequired("MYSQL_PASSWORD")
	mysqlDatabase = util.GetenvRequired("MYSQL_DATABASE")
)

const ()

// NewDB Create a new db connection
func NewDB() error {
	var err error
	DBCon, err = sql.Open("mysql", ""+mysqlUser+":"+mysqlPassword+"@tcp("+mysqlHost+")/"+mysqlDatabase+"")
	checkErr(err)

	err = DBCon.Ping()
	checkErr(err)

	err = runMigrations(DBCon)
	checkErr(err)

	return err
}

// CreatePaper create new paper
func CreatePaper(paper paper.Paper) {
	stmt, _ := DBCon.Prepare("INSERT INTO papers(name, company, marketValue, dailyVariation, url) values (?, ?, ?, ?, ?)")
	stmt.Exec(paper.Name, paper.Company, paper.MarketValue, paper.DailyVariation, paper.URL)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
