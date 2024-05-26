package filesystem

import (
	"os"
	"io/ioutil"
)

var notesDir string

func InitFileSystemModule () {

	dataDir, ok := os.LookupEnv("DATA_DIR")

	if ok {
		notesDir = dataDir + "/notes"
	} else {
		notesDir = "/data/notes"
	}

}

func CreateNewEmptyNote (username string, noteName string) error {

	path := notesDir + "/" + username + "/" + noteName

	err := os.WriteFile(path, []byte(""), 0666)
	if err != nil {
		return err
	}
	return nil

}

func GetNoteContent (username string, noteName string) (string, error) {

	path := notesDir + "/" + username + "/" + noteName

	content, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil

}

func GetUsersNoteTitles (username string) ([]string, error) {

	var filenames []string
	
	files, err := ioutil.ReadDir(notesDir + "/" + username)
	if err != nil {
		return filenames, err
	}	

	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	return filenames, nil
}

func UpdateNoteContent (username string, noteName string, content string) error {

	path := notesDir + "/" + username + "/" + noteName

	err := os.WriteFile(path, []byte(content), 0666)
	if err != nil {
		return err
	}

	return nil

}