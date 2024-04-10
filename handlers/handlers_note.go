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

type errorMessage struct {
	Message string `json:"message"`
}

type fileName struct {
	FileName string `json:"filename"`
}

type notesPageContext struct {
	Username string
	Names    []string
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

	filename := c.PostForm("newNoteName")
	path := c.GetString("username") + "/" + filename

	err := os.WriteFile(path, []byte(""), 0666)
	if err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			errorMessage{Message: "Failed to create file: " + c.Param("filename")},
		)
		return
	}

	c.Redirect(http.StatusFound, "/notes")

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
}

/*
Full page view of a single note, no editor.
*/
func GetFullPageNoteView(c *gin.Context) {

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

	c.HTML(http.StatusFound, "fullpagenoteview.html", singleNote)

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

	c.HTML(http.StatusFound, "editor.html", singleNote)
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

	context := notesPageContext{
		Username: c.GetString("username"),
		Names:    filenames,
	}

	c.HTML(http.StatusFound, "notes.html", context)

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

	c.Redirect(http.StatusFound, "/notes/fullpagenoteview/" + c.Param("filename"))

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
			errorMessage{Message: "Failed to delete note."},
		)
		return
	}

	c.Redirect(http.StatusOK, "/notes")

}

/*
When the user clicks the edit icon next to the file name in the editor, return
an input box where they can set a new file name.
*/
func EditNoteName(c *gin.Context) {

	context := fileName{
		FileName: c.Param("filename"),
	}

	c.HTML(http.StatusOK, "rename.html", context)

}

/*
Rename the file on the server.
*/
func RenameNote(c *gin.Context) {

	pathOld := c.GetString("username") + "/" + c.PostForm("oldFileName")
	pathNew := c.GetString("username") + "/" + c.PostForm("newFileName")

	fmt.Print(pathOld)
	fmt.Print(pathNew)

	e := os.Rename(pathOld, pathNew)
	if e != nil {
		panic(e)
	}

	c.Redirect(http.StatusFound, "/notes/edit/"+c.PostForm("newFileName"))

}
