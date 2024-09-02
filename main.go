package main

import (
	"fmt"
	"vrs/api"
	"vrs/database"
	"vrs/utils"
)

func main() {

	//-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_ Start Database -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-
	database.InitConnectionDatabase()
	//-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_ Start APP -_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_
	app := api.Routes()
	err := app.Listen(fmt.Sprint(":", utils.Dotenv("PORT_APPLICATION")))
	if err != nil {
		utils.CreateFileDay(utils.Message{File: "main.go", Error: err.Error()})
	}
}
