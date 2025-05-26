package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "nucleus-postgres.cx6cu628207f.us-east-2.rds.amazonaws.com"
	port     = 5432
	user     = "root"
	password = "IImbo4jlqXx8ec1BuFdX"
	dbname   = "postgres"
	schema   = "public"

	// host     = "193.56.23.149"
	// port     = 5432
	// user     = "postgres"
	// password = "peak202394"
	// dbname   = "postgres"
	// schema   = "schema"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s",
		host, port, user, password, dbname, schema)

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open connection:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	fmt.Println("Connected to PostgreSQL!")
}
