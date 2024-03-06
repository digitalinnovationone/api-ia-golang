package main

import (
	"api_golang_ia/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	homeController := controllers.HomeController{}
	r.GET("/", homeController.Index)

	palavrasController := controllers.PalavrasController{}
	r.GET("/palavras", palavrasController.Index)

	r.Run() // por padrão, escuta na porta 8080
}
