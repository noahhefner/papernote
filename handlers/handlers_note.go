package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noahhefner/notes/models"
	"noahhefner/notes/filesystem"
	"github.com/go-playground/validator/v10"
	"os"
)

var validate *validator.Validate

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

func InitFieldValidator() {

	validate = validator.New(validator.WithRequiredStructEnabled())

}

/*
Create a new note.
*/
func CreateNote(c *gin.Context) {

	noteName := c.PostForm("newNoteName")

	if !(validateFilename(noteName)) {
		c.IndentedJSON(
			http.StatusBadRequest,
			errorMessage{Message: "Invalid filename"},
		)
		return
	}

	username := c.GetString("username")

	err := filesystem.CreateNewEmptyNote(username, noteName)

	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to create new empty note."},
		)
		return
	}

	noteTiles, err := filesystem.GetUsersNoteTitles(username)

	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to get note titles for user."},
		)
		return
	}

	context := noteList{
		FileNames: noteTiles,
	}

	c.HTML(http.StatusOK, "noteList.html", context)

}

func GetNoteRenderedMarkdown(c *gin.Context) {

	username := c.GetString("username")
	noteName := c.Param("filename")

	content, err := filesystem.GetNoteContent(username, noteName)

	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			errorMessage{Message: "File not found."},
		)
		return
	}

	var singleNote = models.Note{
		FileName: noteName,
		Content:  content,
	}

	c.HTML(http.StatusOK, "notePreview.html", singleNote)
}

func GetFullPageNoteView(c *gin.Context) {

	username := c.GetString("username")
	noteName := c.Param("filename")

	content, err := filesystem.GetNoteContent(username, noteName)

	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			errorMessage{Message: "File not found."},
		)
		return
	}

	var singleNote = models.Note{
		FileName: noteName,
		Content:  content,
	}

	c.HTML(http.StatusOK, "fullpagenoteview.html", singleNote)

}

func GetEditor(c *gin.Context) {

	username := c.GetString("username")
	noteName := c.Param("filename")

	content, err := filesystem.GetNoteContent(username, noteName)

	if err != nil {
		c.IndentedJSON(
			http.StatusNotFound,
			errorMessage{Message: "File not found."},
		)
		return
	}

	var singleNote = models.Note{
		FileName: noteName,
		Content:  content,
	}

	c.HTML(http.StatusOK, "editor.html", singleNote)

}

func GetNotesPage(c *gin.Context) {

	username := c.GetString("username")
	
	noteTiles, err := filesystem.GetUsersNoteTitles(username)

	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to get note titles for user."},
		)
		return
	}

	context := notesPageContext{
		Username:  username,
		FileNames: noteTiles,
	}

	c.HTML(http.StatusFound, "newNotePage.html", context)

}

func UpdateNote(c *gin.Context) {

	username := c.GetString("username")
	noteName := c.Param("filename")
	content := c.PostForm("editor")

	err := filesystem.UpdateNoteContent(username, noteName, content)

	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to update note."},
		)
		return
	}

	c.Redirect(http.StatusSeeOther, "/notes/fullpagenoteview/" + noteName)

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


func getDataDir() string {
	dataDir, ok := os.LookupEnv("DATA_DIR")
	if ok {
		return dataDir
	} else {
		return "/data"
	}
}

func validateFilename(filename string) bool {

	// user did not include .md
	errs := validate.Var(filename, "required,alphanum")
	return errs == nil
}
