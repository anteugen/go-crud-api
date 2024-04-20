package main

import (
	"github.com/anteugen/go-crud-api/db"
   	"github.com/anteugen/go-crud-api/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}