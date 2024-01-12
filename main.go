package main

import (
	"apiuser/database"

	"apiuser/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.ConnectDatabase()

	routes.Api(r, db)

	r.Run(":9888")
}
