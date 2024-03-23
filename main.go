package main

import (
  "github.com/gin-gonic/gin"
  _ "github.com/mattn/go-sqlite3"
  "net/http" 
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

  // Load HTML templates
  router.LoadHTMLGlob("templates/*")

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

  router.GET("/login", func(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", gin.H{
      "name": "login",
    })
  })

  router.Run("localhost:8080")

}