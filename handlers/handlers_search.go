package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"regexp"
)

type SearchContext struct {
	FileNames []string
}

func Search(c *gin.Context) {

	query, ok := c.GetQuery("query")

	if !ok {
		fmt.Println("Failed to find query parameter")
	}

	var path = notesDir + "/" + c.GetString("username")

	files, err := os.ReadDir(path)
    if err != nil {
        fmt.Println("Failed to open user dir")
    }
 
	var filenames []string

	var filenameMatch = regexp.MustCompile(query)

    for _, file := range files {
        if filenameMatch.MatchString(file.Name()) {
			filenames = append(filenames, file.Name())
		}
    }

	searchContext := SearchContext{
		FileNames: filenames,
	}

	c.HTML(http.StatusOK, "noteList.html", searchContext)

}
