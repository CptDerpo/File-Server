package main

import (
	"errors"
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

//deleteFile deletes the file which corresponds to the given fileid.
//Also removes the entry from the database.
func deleteFileOS(path string) error {
	err := os.Remove(path)
	return err
}
