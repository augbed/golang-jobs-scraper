package mysql

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

func init() {
	initDB()
	migrateDB()
}

func initDB() {
	mySQLConnection := os.Getenv("MYSQL_CONNECTION")
	if mySQLConnection == "" {
		log.Panic("MYSQL_CONNECTION must be set")
	}

	db, err := sql.Open("mysql", mySQLConnection)
	if err != nil {
		log.Panic("Error connecting to database")
	}

	DB = db
}

func migrateDB() {
	driver, err := mysql.WithInstance(DB, &mysql.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/mysql", "mysql", driver)
	if err != nil {
		log.Panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
