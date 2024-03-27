package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noahhefner/notes/models"
	"os"
    "io/ioutil"
)

type errorMessage struct {
	Message string `json:"message"`
}

type fileName struct {
	FileName string `json:"filename"`
}

type fileNameList struct {
	Names []string
}

type FileNode struct {
    Name     string      `json:"name"`
    IsDir    bool        `json:"isDir"`
    Children []*FileNode `json:"children,omitempty"`
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
func GetNoteByFilename(c *gin.Context) {

	var path = c.GetString("username") + "/" + c.Param("filename")

	content, err := os.ReadFile(path)
	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			errorMessage{Message: "File not found."},
		)
		return
	}

	var singleNote = models.Note{
		FileName: c.Param("filename"),
		Content:  string(content),
	}

	c.HTML(http.StatusOK, "notePreview.html", singleNote)
	//c.HTML(http.StatusOK, "editor.html", singleNote)
}

/*
  Returns a list of the titles of a single users notes.
*/
func GetAllNotesForUser(c *gin.Context) {

	// TODO: Validate user dir before proceeding
	files, err := ioutil.ReadDir("./" + c.GetString("username"))
    if err != nil {
        panic(err)
    }

	var filenames []string

    for _, file := range files {
		// TODO: Validate file type is .md
		filenames = append(filenames, file.Name())
    }

	c.HTML(http.StatusOK, "notes.html", fileNameList{Names:filenames})

}

/*
  Update an existing note.
*/
func UpdateNote(c *gin.Context) {

	path := c.GetString("username") + "/" + c.Param("filename")
	content := c.PostForm("editor")

	err := os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to update note."},
		)
		return
	}

	var singleNote = models.Note{
		FileName: c.Param("filename"),
		Content:  string(content),
	}

	c.HTML(http.StatusOK, "editor.html", singleNote)

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
