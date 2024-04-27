package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"noahhefner/notes/models"
	"os"
	"strings"
)

var notesDir string = getDataDir() + "/notes"

type errorMessage struct {
	Message string `json:"message"`
}

type notesPageContext struct {
	Username  string
	FileNames []string
}

type noteList struct {
	FileNames []string
}

/*
Create a new note.
*/
func CreateNote(c *gin.Context) {

	filename := c.PostForm("newNoteName")
	path := notesDir + "/" + c.GetString("username") + "/" + filename

	err := os.WriteFile(path, []byte(""), 0666)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to create file: " + c.Param("filename")},
		)
		return
	}

	context := noteList{
		FileNames: getFileListForUser(c.GetString("username")),
	}

	c.HTML(http.StatusOK, "noteList.html", context)

}

/*
Get a single note by id.
*/
func GetNoteByFilename(c *gin.Context) {

	var path = notesDir + "/" + c.GetString("username") + "/" + c.Param("filename")

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
}

/*
Full page view of a single note, no editor.
*/
func GetFullPageNoteView(c *gin.Context) {

	var path = notesDir + "/" + c.GetString("username") + "/" + c.Param("filename")

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

	c.HTML(http.StatusFound, "fullpagenoteview.html", singleNote)

}

/*
Get the editor view and populate it with data from filename.
*/
func GetEditor(c *gin.Context) {

	var path = notesDir + "/" + c.GetString("username") + "/" + c.Param("filename")

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

	c.HTML(http.StatusFound, "editor.html", singleNote)
}

/*
Returns a list of the titles of a single users notes.
*/
func GetAllNotesForUser(c *gin.Context) {

	context := notesPageContext{
		Username:  c.GetString("username"),
		FileNames: getFileListForUser(c.GetString("username")),
	}

	c.HTML(http.StatusFound, "notes.html", context)

}

/*
Update an existing note.
*/
func UpdateNote(c *gin.Context) {

	path := notesDir + "/" + c.GetString("username") + "/" + c.Param("filename")
	content := strings.TrimSpace(c.PostForm("editor"))

	err := os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to update note."},
		)
		return
	}

	c.Redirect(http.StatusFound, "/notes/fullpagenoteview/"+c.Param("filename"))

}

/*
Delete a note.
*/
func DeleteNote(c *gin.Context) {

	path := notesDir + "/" + c.GetString("username") + "/" + c.Param("filename")

	err := os.Remove(path)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to delete note."},
		)
		return
	}

	c.String(http.StatusOK, "")

}

// Helpers ---------------------------------------------------------------------

/*
Get filenames of all a given users notes.
*/
func getFileListForUser(username string) []string {

	fmt.Print(notesDir + "/" + username)

	files, err := ioutil.ReadDir(notesDir + "/" + username)
	if err != nil {
		panic(err)
	}

	var filenames []string

	fmt.Print(files)

	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	return filenames
}

func getDataDir() string {
	dataDir, ok := os.LookupEnv("DATA_DIR")
	if ok {
		return dataDir
	} else {
		return "/data"
	}
}
