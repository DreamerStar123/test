package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
	schema   = "public"

	// host     = "193.56.23.149"
	// port     = 5432
	// user     = "postgres"
	// password = "peak202394"
	// dbname   = "postgres"
	// schema   = "schema"
)

func postgresQuery(query string) string {
	pqQuery, n := "", 1
	for _, v := range query {
		if v != '?' {
			pqQuery += string(v)
		} else {
			pqQuery += fmt.Sprintf("$%d", n)
			n++
		}
	}
	return pqQuery
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
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

	query := postgresQuery(`select * from checkout_session where currency=?`)
	_, err = db.Exec(query, "usd")
	if err != nil {
		log.Fatal("Query exec failed:", query, err)
	}
	fmt.Println("Query exec success!")
}
