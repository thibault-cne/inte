package main

import (
	"backend/db"
	"backend/server"
)

func main() {
	db.InitDatabase()
	server.InitServer()
}
