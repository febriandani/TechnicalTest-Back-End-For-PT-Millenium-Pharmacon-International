package main

import (
	"fmt"
	"technical-test-02-10-22/database"
	"technical-test-02-10-22/router"
)

func main() {
	database.StartDB()
	fmt.Println("Hello World")
	router.SetupRoutes()
}
