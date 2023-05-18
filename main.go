package main

import (
	"qaweb/database"
	"qaweb/router"
)

func main() {
	database.Initdb()
	router.Initrouter()
}
