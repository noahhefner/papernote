package main

import (
  "os"
  "net/http"
  "github.com/gin-gonic/gin"
)

type note struct {
  ID string `json:"id"`
  FileName string `json:"filename"`
  Content string `json:"content"`
}

// Test data
var notes = []note{
  {ID: "1", FileName: "NoteOne", Content:"Some text"},
  {ID: "2", FileName: "NoteTwo", Content:"Some two text"},
  {ID: "3", FileName: "NoteThree", Content:"Some three text"},
}

func main() {
  router := gin.Default()
  router.GET("/notes", getNotes)
  router.GET("/notes/:id", getNoteByID)
  router.POST("/notes/:id", createNote)
  router.PATCH("/notes/:id", updateNote)
  router.DELETE("/notes/:id", deleteNote)

  router.Run("localhost:8080")
}

/*
Delete a note.
*/
func deleteNote(c *gin.Context) {
  id := c.Param("id")
  for i, n := range notes {
    if n.ID == id {
      notes = append(notes[:i], notes[i+1:]...)
      c.IndentedJSON(http.StatusOK, n)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}

/*
Update an existing note.
*/
func updateNote(c *gin.Context) {
  id := c.Param("id")
  for i, n := range notes {
    if n.ID == id {
      var updatedNote note
      if err := c.BindJSON(&updatedNote); err != nil {
        return
      }
      notes[i] = updatedNote
      c.IndentedJSON(http.StatusOK, updatedNote)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
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
func getNoteByID(c *gin.Context){
  id := c.Param("id")
  for _, n := range notes {
    if n.ID == id {
      c.IndentedJSON(http.StatusOK, n)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}
