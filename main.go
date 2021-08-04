package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mniak/Alkanoid/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterAllRoutes(r)
	r.Run()
}
