package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/static"
  _ "github.com/mattn/go-sqlite3"

  "noahhefner/notes/database"
  "noahhefner/notes/handlers"

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

  // Note routes
  router.POST("/notes", handlers.CreateNote)
  router.GET("/notes", handlers.GetNoteByFilename)
  router.PATCH("/notes", handlers.UpdateNote)
  router.DELETE("/notes", handlers.DeleteNote)

  // User routes
  router.POST("/login", handlers.Login)
	router.POST("/users", handlers.AddUser)
	router.GET("/users/:username", handlers.GetUser)

  router.Run("0.0.0.0:8080")

}