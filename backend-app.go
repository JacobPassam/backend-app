package main

import (
	"backend-app/config"
	"backend-app/db"
	"backend-app/http"
	"fmt"
)

func main() {
	if err := config.Setup(); err != nil {
		fmt.Println(err)
	}

	db.NewMongoHandler()
	http.Run()

	defer db.Disconnect()
}
