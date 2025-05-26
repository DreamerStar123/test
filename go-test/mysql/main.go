package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Set up the connection string
	// dsn := "admin:Password2024@tcp(nucleus-db.cx6cu628207f.us-east-2.rds.amazonaws.com:3306)/webshop_merged"
	dsn := "admin:LqGxTLiIdCuI2vcExkmM@tcp(wallet-root-db.cryo0ss8qnz0.ap-southeast-2.rds.amazonaws.com:3306)/webshop_merged"
	// Replace username, password, host, port, and dbname with your values

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}
