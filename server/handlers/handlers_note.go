package handlers

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"noahhefner/notes/models"
)

type errorMessage struct {
	Message string `json:"message"`
  }

type fileName struct {
	FileName string `json:"filename"`
}

/*
Create a new note.
*/
func CreateNote(c *gin.Context) {
  
	var createdNote models.Note

	err := c.BindJSON(&createdNote)
	if err != nil {
	  c.IndentedJSON(
		http.StatusInternalServerError, 
		errorMessage{Message: "Unmarshalling request data failed."},
	  )
	  return
	}

	err = os.Chdir(c.GetString("username"))
	if err != nil {
	  c.IndentedJSON(
		http.StatusNotFound, 
		errorMessage{Message: "User not found: " + c.Param("user")},
	  )
	  return
	}
  
	err = os.WriteFile(createdNote.FileName, []byte(createdNote.Content), 0666)
	if err != nil {
	  c.IndentedJSON(
		http.StatusInternalServerError, 
		errorMessage{Message: "Failed to create file: " + c.Param("filename")},
	  )
	  return
	}
  
	c.IndentedJSON(http.StatusCreated, createdNote)
  
  }
  
  /*
  Get a single note by id.
  */
  func GetNoteByFilename(c *gin.Context){

	err := os.Chdir(c.GetString("username"))
	if err != nil {
	  c.IndentedJSON(
		http.StatusNotFound, 
		errorMessage{Message: "User not found: " + c.Param("user")},
	  )
	  return
	}
  
	var fileNameRequested fileName

	err = c.BindJSON(&fileNameRequested)
	if err != nil {
	  c.IndentedJSON(
		http.StatusInternalServerError, 
		errorMessage{Message: "Unmarshalling request data failed."},
	  )
	  return
	}

	content, err := os.ReadFile(fileNameRequested.FileName)
	if err != nil {
	  c.IndentedJSON(
		http.StatusNotFound, 
		errorMessage{Message: "File not found: " + fileNameRequested.FileName},
	  )
	  return
	}
  
	aNote := models.Note {
	  FileName: fileNameRequested.FileName,
	  Content: string(content),
	}
  
	c.IndentedJSON(http.StatusOK, aNote)
  }
  
  /*
  Update an existing note.
  */
  func UpdateNote(c *gin.Context) {
  
	err := os.Chdir(c.GetString("username"))
	if err != nil {
	  c.IndentedJSON(
		http.StatusNotFound, 
		errorMessage{Message: "User not found: " + c.Param("user")},
	  )
	  return
	}
  
	var noteToUpdate models.Note
  
	err = c.BindJSON(&noteToUpdate)
	if err != nil {
	  c.IndentedJSON(
		http.StatusInternalServerError, 
		errorMessage{Message: "Unmarshalling request data failed."},
	  )
	  return
	}
	
	err = os.WriteFile(noteToUpdate.FileName, []byte(noteToUpdate.Content), 0666)
	if err != nil {
	  c.IndentedJSON(
		http.StatusInternalServerError, 
		errorMessage{Message: "Failed to update file: " + noteToUpdate.FileName},
	  )
	  return
	}
	
	c.IndentedJSON(http.StatusOK, noteToUpdate)
  }
  
  /*
  Delete a note.
  */
  func DeleteNote(c *gin.Context) {
  
	err := os.Chdir(c.GetString("username"))
	if err != nil {
	  c.IndentedJSON(
		http.StatusNotFound, 
		errorMessage{Message: "User not found: " + c.Param("user")},
	  )
	  return
	}

	var fileNameRequested fileName

	err = c.BindJSON(&fileNameRequested)
	if err != nil {
	  c.IndentedJSON(
		http.StatusInternalServerError, 
		errorMessage{Message: "Unmarshalling request data failed."},
	  )
	  return
	}
  
	err = os.Remove(fileNameRequested.FileName)
	if err != nil {
	  c.IndentedJSON(
		http.StatusInternalServerError, 
		errorMessage{Message: "Failed to delete note: " + fileNameRequested.FileName},
	  )
	  return
	}
  
	c.IndentedJSON(http.StatusOK, "Deleted note.")
  
  }