package main

import (
	"backend-app/config"
	"backend-app/db"
	"fmt"
)

func main() {
	var mongo = db.NewMongoHandler()

	if err := config.Setup(); err != nil {
		fmt.Println(err)
	}

	mongo.Disconnect()
}
