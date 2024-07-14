package main

import (
	routing "belajar_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies(nil)
	routing.SetupRoute(router)

	router.Run(":8080")
}
