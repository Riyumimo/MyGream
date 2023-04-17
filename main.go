package main

import (
	"MyGream/database"
	"MyGream/router"
)

func main() {
	database.StartDb()
	r := router.StartApp()
	r.Run(":8080")
}
