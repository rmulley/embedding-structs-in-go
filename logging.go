package main

import (
	"database/sql"
	"log"
	"net/url"
	"os"
)

type DB struct {
	*sql.DB
	*log.Logger
}

func NewDB() *DB {
	var db = new(DB)

	db.DB, _ = sql.Open("mysql", "user:pass@tcp(localhost:3306)/db_name?"+url.QueryEscape("charset=utf8mb4,utf8&loc=America/New_York"))
	db.Logger = log.New(os.Stderr, "mysql: ", log.LstdFlags)

	return db
}

type Redis struct {
	// Embed some Redis struct
	*log.Logger
}

func NewRedis() *Redis {
	return &Redis{
		log.New(os.Stderr, "redis: ", log.LstdFlags),
	}
}

func main() {
	var (
		db    = NewDB()
		redis = NewRedis()
	)

	log.Println("Some generic error here.")
	db.Println("Insert statement failed")
	redis.Println("Connection error")
}
