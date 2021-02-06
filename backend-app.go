package main

import (
	"backend-app/config"
	"backend-app/db"
	"fmt"
)

func main() {
	if err := config.Setup(); err != nil {
		fmt.Println(err)
	}

	var mongo = db.NewMongoHandler()

	if err := mongo.Disconnect(); err != nil {
		fmt.Println(err)
	}
}
