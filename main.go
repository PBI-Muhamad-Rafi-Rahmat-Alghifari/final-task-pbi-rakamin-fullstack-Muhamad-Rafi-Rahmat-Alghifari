package main

import (
	"FinalTaskAPI/database"
	"FinalTaskAPI/router"
	"os"
)

func main() {
	database.StartDB()
	var PORT = os.Getenv("PORT")

	r := router.StartApp()
	r.Run(":" + PORT)
}
