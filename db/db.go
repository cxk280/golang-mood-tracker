package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

// Define the DB struct
type DB struct {
	*sql.DB
}

const (
	DbUser = "postgres"
	DbPassword = "postgres"
	DbName = "golang_gin_db"
)

var db *gorp.DbMap

// Initialize the database
func Init() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

// Connect to the database
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

// Get the database
func GetDB() *gorp.DbMap {
	return db
}
