package main

import (
	"dashboard/internal/dashboard"
	"dashboard/internal/router"
	"log"
)

func main() {
	dashboard, err := dashboard.Init()
	if err != nil {
		log.Fatal("Couldn't init dashboard service", err)
	}
	router := router.SetupRouter(dashboard)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
