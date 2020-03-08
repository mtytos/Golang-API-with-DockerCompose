package postgres

import (
	"database/sql"
	"fmt"
	"os"
	// "syscall"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// DbHandler for DB connections
type DbHandler struct {
	DB *sql.DB
}

func getDBHadler() *DbHandler {
	connStr := getDsn()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatal("Unable to open database")
	}
	DbHandler := &DbHandler{
		DB: db,
	}

	return DbHandler
}

func InitDBTables() {
	DbHandler := getDBHadler()
	defer DbHandler.DB.Close()

	query := "CREATE SCHEMA IF NOT EXISTS register_journal"
	_, err := DbHandler.DB.Query(query)
	if err != nil {
		log.WithFields(log.Fields{
			"SCHEMA": "register_journal",
			"err":    err,
		}).Fatal("SCHEMA CREATE ERROR")
	}

	query = "CREATE TABLE IF NOT EXISTS register_journal.register_info (id SERIAL PRIMARY KEY, employee_name varchar(255) NOT NULL, location_id INTEGER NOT NULL, terminal_id INTEGER NOT NULL, usage_begin_time VARCHAR, usage_end_time VARCHAR)"
	_, err = DbHandler.DB.Query(query)
	if err != nil {
		log.WithFields(log.Fields{
			"db":  "tasks",
			"err": err,
		}).Fatal("DB CREATE ERROR")
	}
}

func RegisterNewUser(userName, locationId, terminalId string) {

	DbHandler := getDBHadler()
	defer DbHandler.DB.Close()

	giveTerm := time.Now()
	tGive := giveTerm.Format("01-02-2006 15:04:05")
	tBack := "Using"
	query := "INSERT INTO register_journal.register_info(employee_name, location_id, terminal_id, usage_begin_time, usage_end_time) VALUES ($1, $2, $3, $4, $5)"
	_, err := DbHandler.DB.Exec(query, userName, locationId, terminalId, tGive, tBack)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "INSERT InsertNewDataCSV",
			"err":  err,
		}).Fatal("DB ERROR")
	}
}

func UnregisterUser(terminalId string) {

	DbHandler := getDBHadler()
	defer DbHandler.DB.Close()

	giveTerm := time.Now()
	tGive := giveTerm.Format("01-02-2006 15:04:05")
	thirdArg := "Using"
	query := "UPDATE register_journal.register_info set usage_end_time = $1 WHERE terminal_id = $2 AND usage_end_time = $3"
	_, err := DbHandler.DB.Exec(query, tGive, terminalId, thirdArg)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "INSERT update",
			"err":  err,
		}).Fatal("DB ERROR")
	}
}

func getDsn() string {

	host, err := os.LookupEnv("SQL_HOST")
	port, _ := os.LookupEnv("SQL_PORT")
	user, _ := os.LookupEnv("SQL_USER")
	password, _ := os.LookupEnv("SQL_PASSWORD")
	database, _ := os.LookupEnv("SQL_DB")
	if err != true {
		log.Fatal("Env error: ")
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database)

	return dsn
}


//select employee_name, location_id, terminal_id, usage_begin_time, usage_end_time from register_journal.register_info;