package main

import (
  "os"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/static"
)

type note struct {
  FileName string `json:"filename"`
  Content string `json:"content"`
}

// Test data
var notes = []note{
  {FileName: "NoteOne", Content:"Some text"},
  {FileName: "NoteTwo", Content:"Some two text"},
  {FileName: "NoteThree", Content:"Some three text"},
}

func main() {
  router := gin.Default()

  // Serve Vue.js frontend
  router.Use(static.Serve("/", static.LocalFile("/usr/src/app/dist", false)))

  router.POST("/notes/", createNote)
  router.GET("/notes/", getNote)
  router.PATCH("/notes/", updateNote)
  router.DELETE("/notes/", deleteNote)

  router.Run("0.0.0.0:8080")
}

/*
Delete a note.
*/
func deleteNote(c *gin.Context) {
  if err := os.Remove(c.Param("filename")); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, "Failed to delete note.")
    return
  }

  c.IndentedJSON(http.StatusOK, "Deleted note.")
  return
}

/*
Update an existing note.
*/
func updateNote(c *gin.Context) {

  // Build note struct from request
  var updatedNote note
  if err := c.BindJSON(&updatedNote); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, updatedNote)
    return
  }
  
  // Write to file
  if err := os.WriteFile(updatedNote.FileName, []byte(updatedNote.Content), 0666); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, updatedNote)
    return
  }
  
  c.IndentedJSON(http.StatusOK, updatedNote)
}

/*
Create a new note.
*/
func createNote(c *gin.Context) {

  var newNote note

  if err := c.BindJSON(&newNote); err != nil {
    return
  }

  if err := os.WriteFile(newNote.FileName, []byte(newNote.Content), 0666); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, newNote)
    return
  }

  c.IndentedJSON(http.StatusCreated, newNote)

}

/*
Get all notes.
*/
func getNotes(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, notes)
}

/*
Get a single note by id.
*/
func getNote(c *gin.Context){

  // Grab content from file
  content, err := os.ReadFile(c.Param("filename"))
  if err != nil {
    return
  }

  // Build note struct
  aNote := note {
    FileName: c.Param("filename"),
    Content: string(content),
  }

  c.IndentedJSON(http.StatusOK, aNote)
}
