package internal

import (
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Path      string `json:"path"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

func LoadFile(path string) File {
	openedFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer openedFile.Close()

	file := File{
		Path:      path,
		Name:      filepath.Base(openedFile.Name()),
		Extension: filepath.Ext(openedFile.Name()),
	}

	return file
}
