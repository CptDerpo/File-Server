package main

import (
	"errors"
	"fmt"
	"os"
	pathpkg "path"

	_ "github.com/mattn/go-sqlite3"
)

//createFile makes a new file at location given to this function.
//On success a pointer is returned to this new file, or an error if it already exists.
func createFileOS(path string, size int64) (f *os.File, err error) {
	f, err = os.OpenFile(pathpkg.Join("home", path), os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
	fmt.Println(pathpkg.Join("home", path))
	if errors.Is(err, os.ErrExist) || err != nil {
		return nil, err
	}
	return f, nil
}

//renameFile moves the old path to the new path.
//As a result, files will be renamed or moved.
//Returns an error of type *os.LinkError if moving fails, or DB error.
func moveFileOS(oldPath string, newPath string) error {
	oldPath = pathpkg.Join("home", oldPath)
	newPath = pathpkg.Join("home", newPath)
	err := os.Rename(oldPath, newPath)
	return err
}

//deleteFile deletes the file which corresponds to the given fileid.
//Also removes the entry from the database.
//Returns error if any DB operation fails.
func deleteFileOS(path string) error {
	err := os.Remove(path)
	return err
}
