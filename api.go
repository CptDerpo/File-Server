package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	pathpkg "path"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func downloadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Query().Get("path")

	if _, err := os.Stat(pathpkg.Join("home", path)); errors.Is(err, os.ErrNotExist) {
		w.WriteHeader(400)
		return
	}

	_, filename := pathpkg.Split(path)

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(filename))

	log.Println("Served the file: " + pathpkg.Join("home", path))
	http.ServeFile(w, r, pathpkg.Join("home", path))
	w.WriteHeader(200)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(4 << 30) //4 * 1024^3 bytes = 4GiB max in ram, rest on disk, add max file size
	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		return
	}
	defer file.Close()

	tempFile, err := createFileOS(pathpkg.Join("home", handler.Filename), handler.Size)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	_, err = tempFile.Write(fileBytes)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	log.Println("File uploaded: " + handler.Filename)
	w.WriteHeader(200)
}

func removeFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filepath := r.URL.Query().Get("path")

	if filepath == "" {
		w.WriteHeader(400)
		return
	}

	log.Println(pathpkg.Join("home", filepath))

	err := deleteFileOS(pathpkg.Join("home", filepath))

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	w.WriteHeader(200)
}

func moveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	oldPath := r.URL.Query().Get("oldPath")
	newPath := r.URL.Query().Get("newPath")

	if _, err := os.Stat(pathpkg.Join("home", oldPath)); errors.Is(err, os.ErrNotExist) {
		w.WriteHeader(400)
		return
	}

	err := moveFileOS(pathpkg.Join("home", oldPath), pathpkg.Join("home", newPath))

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

func getDir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dirpath := r.URL.Query().Get("path")
	result, err := dirTreeOS(pathpkg.Join("home", dirpath))

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

	json.NewEncoder(w).Encode(result)
}
