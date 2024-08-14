package main

import (
	"card-bridge/database"
	"card-bridge/routers"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// input
// go get -u "github.com/gin-gonic/gin"
// go get -u "github.com/lib/pq"
// go get -u "github.com/rubenv/sql-migrate"
// go get -u "github.com/joho/godotenv"

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "wkwkwkwk"
// 	dbName   = "praktikum_mcs_bab6"
// )

var (
	DB  *sql.DB
	err error
)

func main() {
	// var PORT = ":8080"

	var PORT = ":" + os.Getenv("PORT")

	// psqlInfo := fmt.Sprintf(
	// 	`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
	// 	host, port, user, password, dbName,
	// )

	// psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// )

	psqlInfo := os.Getenv("DATABASE_URL")

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error Open DB: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging Database: %v\n", err)
	}

	database.DBMigrate(DB)

	defer DB.Close()

	routers.StartServer().Run(PORT)
	fmt.Print("Success Connetced")
}
