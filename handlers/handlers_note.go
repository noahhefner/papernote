package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noahhefner/notes/models"
	"os"
	"time"
    "io/ioutil"
	"strings"
)

type errorMessage struct {
	Message string `json:"message"`
}

type fileName struct {
	FileName string `json:"filename"`
}

type notesPageContext struct {
	Username string
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

	c.Header("Cache-Control", "max-age=3600")  // Cache response for 1 hour
    expires := time.Now().Add(time.Hour)
    c.Header("Expires", expires.Format(time.RFC1123))  // Example expiration date

	c.HTML(http.StatusOK, "notePreview.html", singleNote)
	//c.HTML(http.StatusOK, "editor.html", singleNote)
}

func GetEditor(c *gin.Context) {

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
		Content:  strings.TrimSpace(string(content)),
	}

	c.HTML(http.StatusOK, "editor.html", singleNote)
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

	context := notesPageContext {
		Username: c.GetString("username"),
		Names: filenames,
	}

	c.HTML(http.StatusOK, "notes.html", context)

}

/*
  Update an existing note.
*/
func UpdateNote(c *gin.Context) {

	path := c.GetString("username") + "/" + c.Param("filename")
	content := strings.TrimSpace(c.PostForm("editor"))

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
		Content:  strings.TrimSpace(string(content)),
	}

	c.HTML(http.StatusOK, "editor.html", singleNote)

}

/*
  Delete a note.
*/
func DeleteNote(c *gin.Context) {

	path := c.GetString("username") + "/" + c.Param("filename")

	err := os.Remove(path)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to delete note." },
		)
		return
	}

	c.IndentedJSON(http.StatusOK, "Deleted note.")

}

/*
When the user clicks the edit icon next to the file name in the editor, return
an input box where they can set a new file name.
*/
func EditNoteName(c *gin.Context) {

	context := fileName {
		FileName: c.Param("filename"),
	}

	c.HTML(http.StatusOK, "rename.html", context)

}

/*
Rename the file on the server.
*/
func RenameNote(c *gin.Context) {

	pathOld := c.GetString("username") + "/" + c.Param("filenameOld")
	pathNew := c.GetString("username") + "/" + c.Param("filenameNew")


	e := os.Rename(pathOld, pathNew) 
    if e != nil { 
        panic(e)
    }

	context := fileName {
		FileName: c.Param("filename"),
	}

	c.HTML(http.StatusOK, "newFilename.html", context)

}