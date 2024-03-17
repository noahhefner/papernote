package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/static"
  _ "github.com/mattn/go-sqlite3"

  "noahhefner/notes/database"
  "noahhefner/notes/handlers"
  "noahhefner/notes/middlewares"

)

func main() {

  // Connect to SQLite database
	var err error
	err = database.Init("./users.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

  router := gin.Default()

  // Serve Vue.js frontend
  router.Use(static.Serve("/", static.LocalFile("/usr/src/app/dist", false)))

  // Note routes requiring auth
  authorized := router.Group("/")

  authorized.Use(middlewares.AuthMiddleware()) 
  {
    
    authorized.POST("/notes", handlers.CreateNote)
    authorized.GET("/notes", handlers.GetNoteByFilename)
    authorized.PATCH("/notes", handlers.UpdateNote)
    authorized.DELETE("/notes", handlers.DeleteNote)

    authorized.GET("/users/:username", handlers.GetUser)

  }

  // No auth required
  router.POST("/login", handlers.Login)
	router.POST("/users", handlers.AddUser)
	

  router.Run("0.0.0.0:8080")

}