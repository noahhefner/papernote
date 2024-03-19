package handlers

import (
	"os"
	"fmt"
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
		errorMessage{Message: "User not found"},
	  )
	  return
	}
  
	var fileNameRequested = c.Param("filename")

	content, err := os.ReadFile(fileNameRequested)
	if err != nil {
	  c.IndentedJSON(
		http.StatusNotFound, 
		errorMessage{Message: "File not found: " + fileNameRequested},
	  )
	  return
	}
  
	aNote := models.Note {
	  FileName: fileNameRequested,
	  Content: string(content),
	}
  
	c.IndentedJSON(http.StatusOK, aNote)
  }

  func GetAllNotesForUser(c *gin.Context) {

	err := os.Chdir(c.GetString("username"))
	if err != nil {
	  c.IndentedJSON(
		http.StatusNotFound, 
		errorMessage{Message: "User not found: " + c.Param("user")},
	  )
	  return
	}

	 // Open the current directory
	 dir, err := os.Open(".")
	 if err != nil {
		c.IndentedJSON(
			http.StatusNotFound, 
			errorMessage{Message: "Failed to open user directory."},
		  )
		 return
	 }
	 defer dir.Close()
 
	 // Read all files in the directory
	 files, err := dir.Readdir(0)
	 if err != nil {
		c.IndentedJSON(
			http.StatusNotFound, 
			errorMessage{Message: "Failed to read files in user directoy."},
		  )
		 return
	 }
 
	 var fileContents []models.Note
 
	 // Iterate over the files
	 for _, fileInfo := range files {
		 // Check if the file is not a directory
		 if !fileInfo.IsDir() {
	 
			 // Read the entire content of the file
			 content, err := os.ReadFile(fileInfo.Name())
			 if err != nil {
				 fmt.Println("Error reading file:", err)
				 continue
			 }
 
			 // Create a FileContent struct and append it to the slice
			 fileContents = append(fileContents, models.Note{
				 FileName:  fileInfo.Name(),
				 Content: 	string(content),
			 })
		 }
	 }

	 c.IndentedJSON(http.StatusOK, fileContents)

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