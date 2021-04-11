package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type File struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Type     string `json:"filetype"`
	Date     string `json:"date"`
	Size     int    `json:"size"`
}

var files []File

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage lmao")
}

func login(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./static/login.html")
	tmp.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./static/register.html")
	tmp.Execute(w, nil)
}

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

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/api/getFiles", getFiles)
	http.HandleFunc("/api/getFile", getFile)
	http.Handle("/static/css/style.css", http.StripPrefix("/static/css", http.FileServer(http.Dir("static/css"))))
	http.Handle("/lol/test.mkv", http.StripPrefix("/lol", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	createDB()
	files = append(files, File{Filename: "test.mkv", Path: "/static/", Type: "mkv", Date: "11/04/2021", Size: 10003})
	files = append(files, File{Filename: "malaka.mkv", Path: "/static/", Type: "mkv", Date: "11/04/2021", Size: 10003})
	handleRequests()
}
