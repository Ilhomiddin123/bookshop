package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var database *sql.DB

func initDB() *sql.DB {
	psqlInfo := fmt.Sprintf(`
				host=%s 
				port=%d 
				user=%s 
				password=%s 
				dbname=%s 
				sslmode=disable`,
		"localhost", 5432, "postgres", "12345", "bookshop_db")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

// StartDbConnection Creates connection to database
func StartDbConnection() {
	database = initDB()
}

// GetDBConn func for getting db conn globally
func GetDBConn() *sql.DB {
	return database
}

func CloseDBConn() {
	database.Close()
}
