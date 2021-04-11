package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

/*
Returns all files in json format.
*/
func getFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
	log.Print(files)
}

/*
Returns a file(s) in json format which can be filtered on date, type, filename.
Can only query 1 attribute at a time, otherwise a random attribute will be returned.
*/
func getFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filename := r.URL.Query()["filename"]
	filetype := r.URL.Query()["type"]
	filedate := r.URL.Query()["date"]

	var result []File

	for _, object := range files {
		if len(filename) == 0 {
			if len(filetype) == 0 {
				if len(filedate) == 0 {
					break
				} else if object.Date == filedate[0] {
					result = append(result, object)
				}
			} else if object.Type == filetype[0] {
				result = append(result, object)
			}
		} else if object.Filename == filename[0] {
			result = append(result, object)
		}
	}

	json.NewEncoder(w).Encode(result)
}

/*
func removeFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fileid := strconv.Atoi(r.URL.Query()["fileid"].FormValue())

	for _, object := range files {
		if object.
	}
}*/
