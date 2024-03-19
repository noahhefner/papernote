package main

import (
  "time"
  "github.com/gin-gonic/gin"
  _ "github.com/mattn/go-sqlite3"
  "github.com/gin-contrib/cors"

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

  config := cors.DefaultConfig()
  config.AllowAllOrigins = true
  config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
  config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
  config.ExposeHeaders = []string{"Content-Length"}
  config.AllowCredentials = true
  config.MaxAge = 12 * time.Hour

  router.Use(cors.New(config))

  // Routes requiring auth
  authorized := router.Group("/")

  authorized.Use(middlewares.AuthMiddleware()) 
  {
    
    authorized.POST("/notes", handlers.CreateNote)
    authorized.GET("/notes", handlers.GetAllNotesForUser)
    authorized.GET("/notes/:filename", handlers.GetNoteByFilename)
    authorized.PATCH("/notes", handlers.UpdateNote)
    authorized.DELETE("/notes", handlers.DeleteNote)

    authorized.GET("/users/:username", handlers.GetUser)

  }

  // Routes not requiring auth
  router.POST("/login", handlers.Login)
	router.POST("/users", handlers.AddUser)
	

  router.Run("localhost:8080")

}