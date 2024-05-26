package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
	"noahhefner/notes/database"
	"noahhefner/notes/handlers"
	"noahhefner/notes/filesystem"
	"noahhefner/notes/middlewares"
)

var dbDir string = getDataDir() + "/db"

func main() {

	// Connect to SQLite database
	var err error
	err = database.Init(dbDir + "/users.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	middlewares.InitJWTSecret()
	handlers.InitFieldValidator()
	filesystem.InitFileSystemModule()

	router := gin.Default()

	router.Static("/static", "./public")

	router.GET("/favicon.ico", func(c *gin.Context) {
        c.File("./public/img/favicon.ico")
    })

	router.LoadHTMLGlob("templates/**/*")

	authorized := router.Group("/")

	authorized.Use(middlewares.AuthMiddleware())
	{

		authorized.POST("/notes/create", handlers.CreateNote)
		authorized.GET("/notes", handlers.GetNotesPage)
		authorized.GET("/clickedNewNote", func(c *gin.Context) {
			c.HTML(http.StatusOK, "clickedNewNote.html", gin.H{
				"name": "clickedNewNote",
			})
		})
		authorized.GET("/clickedCancelNewNote", func(c *gin.Context) {
			c.HTML(http.StatusOK, "newNoteButton.html", gin.H{
				"name": "clickedCancelNewNote",
			})
		})
		authorized.GET("/notes/:filename", handlers.GetNoteRenderedMarkdown)
		authorized.GET("/notes/fullpagenoteview/:filename", handlers.GetFullPageNoteView)
		authorized.GET("/notes/edit/:filename", handlers.GetEditor)
		authorized.POST("/notes/:filename", handlers.UpdateNote)
		authorized.DELETE("/notes/:filename", handlers.DeleteNote)

		authorized.GET("/search", handlers.Search)

		authorized.GET("/users/:username", handlers.GetUser)

	}

	// Create new user
	router.POST("/users", handlers.AddUser)
	
	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"name": "signup",
		})
	})

	router.POST("/login", handlers.Login)
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"name": "login",
		})
	})
	router.POST("/logout", handlers.Logout)

	router.GET("/", func(c *gin.Context) {
		// If the user is authenticated, they will be redirected to /notes
		// Otherwise, the auth middleware will redirect to /login
		c.Redirect(http.StatusTemporaryRedirect, "/notes")
	})

	router.Run(":8080")

}

func getDataDir() string {
	dataDir, ok := os.LookupEnv("DATA_DIR")
	if ok {
		return dataDir
	} else {
		return "/data"
	}
}