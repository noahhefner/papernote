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

type content struct {
  Content string `json:"content"`
}

func main() {
  router := gin.Default()

  // Serve Vue.js frontend
  router.Use(static.Serve("/", static.LocalFile("/usr/src/app/dist", false)))

  router.POST("/:user/notes/:filename", createNote)
  router.GET("/:user/notes/:filename", getNoteByFilename)
  router.PATCH("/:user/notes/:filename", updateNote)
  router.DELETE("/:user/notes/:filename", deleteNote)

  router.Run("0.0.0.0:8080")
}

/*
Delete a note.
*/
func deleteNote(c *gin.Context) {
  var filePath = c.Param("user") + "/" + c.Param("filename")
  if err := os.Remove(filePath); err != nil {
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

  var noteContent content

  if err := c.BindJSON(&noteContent); err != nil {
    return
  }

  if err := os.Chdir(c.Param("user")); err != nil {
    c.IndentedJSON(http.StatusNotFound, c.Param("user"))
    return
  }

  if err := os.WriteFile(c.Param("filename"), []byte(noteContent.Content), 0666); err != nil {
    c.IndentedJSON(http.StatusInternalServerError, noteContent)
    return
  }

  c.IndentedJSON(http.StatusCreated, noteContent)

}

/*
Get all notes.
*/
//func getNotes(c *gin.Context) {
//  c.IndentedJSON(http.StatusOK, notes)
//}

/*
Get a single note by id.
*/
func getNoteByFilename(c *gin.Context){

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
