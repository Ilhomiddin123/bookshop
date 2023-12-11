package main

import (
	"bookShop/db"
	"bookShop/routes"
)

func main() {
	db.StartDbConnection()

	routes.RunRoutes()
}
