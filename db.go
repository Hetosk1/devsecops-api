package main

import (
	"database/sql"
	"log"
	"os"
	"github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB 

func ConnectDB(){
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = "postgres://root:root@postgres:5432/ecommerce?sslmode=disable"
	}

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal("Failed to open DB: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	DB = db
	log.Println("Connected to PostgreSQL")
}
