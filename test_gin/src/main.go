package main

import (
	"os"
	"log"

	// ローカルパッケージパス
	"app/model/todo"
	"app/web"

	"github.com/gin-gonic/gin"
)

func main() {
	todo.Migrate()

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")

	// [START todo web]
	todo := router.Group("/todo")
	{
		todo.GET("/", web.TodoRoot)
		todo.POST("/new", web.TodoPost)
		todo.GET("/detail/:id", web.TodoDetail)
		todo.POST("/update/:id", web.TodoUpdate)
		todo.GET("/delete/check/:id", web.TodoDeleteCheck)
		todo.POST("/delete/:id", web.TodoDelete)
	}
	// [END todo web]

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	// [END setting_port]

	router.Run(":" + port)
}
