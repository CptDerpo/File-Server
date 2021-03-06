package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//createFile makes a new file at location given to this function.
//On success a pointer is returned to this new file,
//or an error of type *PathError if it already exists.
func createFileOS(path string, size int64) (f *os.File, err error) {
	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	if errors.Is(err, os.ErrExist) || err != nil {
		return nil, err
	}
	return f, nil
}

//renameFile moves the old path to the new path.
//As a result, files will be renamed or moved.
//Returns an error of type *os.LinkError if moving fails.
func moveFileOS(oldPath string, newPath string) error {
	err := os.Rename(oldPath, newPath)
	return err
}

//deleteFile deletes the file which is located at the given path.
//Returns an error of type *PathError on failure
func deleteFileOS(path string) error {
	err := os.Remove(path)
	return err
}

type Tree struct {
	Filename string `json:"name"`
	IsDir    bool   `json:"isdir"`
	Filesize int64  `json:"size"`
	ModDate  string `json:"date"`
}

func dirTreeOS(path string) ([]Tree, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var tree []Tree

	for _, f := range files {
		fileinfo, err := f.Info()
		if err != nil {
			log.Println(err)
			continue
		}

		var entry Tree
		entry.Filename = fileinfo.Name()
		entry.IsDir = fileinfo.IsDir()
		entry.Filesize = fileinfo.Size()
		entry.ModDate = fileinfo.ModTime().String()
		tree = append(tree, entry)
	}

	result, err := json.Marshal(tree)

	if err != nil {
		return nil, err
	}
	log.Printf(string([]byte((result))))

	return tree, nil
}
