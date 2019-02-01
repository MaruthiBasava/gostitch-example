package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/jmoiron/sqlx"
	sq "github.com/masterminds/squirrel"
	"github.com/smotes/purse"
)

const (
	dbhost     = "localhost"
	dbport     = "5432"
	dbname     = "pet_store"
	dbuser     = "GEN_DB_USER"
	dbpassword = "GEN_DB_PASS"
)

var psMem = getSQLFiles()
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func presetup(db *sqlx.DB, needed bool) {
	if needed {
		execute("sql/create_tables.sql", db)
	}
}

func execute(fileName string, db *sqlx.DB) {
	contents := getContents(fileName)
	_, err := db.Exec(contents)

	if err != nil {
		fmt.Println(err)
		panic(fmt.Sprintf("Error has occured executing SQL statements from %s", fileName))
	}
}

func getContents(fileName string) string {
	contents, ok := psMem.Get(fileName)

	if !ok {
		panic("Error has occured with fetching contents")
	}

	return contents
}

func getSQLFiles() *purse.MemoryPurse {

	ps, err := purse.New(filepath.Join(".", "sql"))

	if err != nil {
		debug.PrintStack()
		fmt.Println(err)
		panic("Error has occured with fetching SQL Files.")
	}

	return ps
}

func InitializeDb() *sqlx.DB {

	psql := formatedDbInfo()

	fmt.Println(psql)

	db, err := sqlx.Open("postgres", psql)

	if err != nil {
		panic("Error")
	}

	err = db.Ping()

	if err != nil {
		panic("Error")
	}

	return db
}

func formatedDbInfo() string {
	config := configureDb()
	target := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	return fmt.Sprintf(target, dbhost, dbport, config[dbuser], config[dbpassword], dbname)
}

func configureDb() map[string]string {

	config := make(map[string]string)

	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("Sorry, something went wrong")
	}

	pass, ok := os.LookupEnv(dbpassword)
	if !ok {
		panic("Sorry, something went wrong")
	}

	config[dbuser] = user
	config[dbpassword] = pass

	return config
}
