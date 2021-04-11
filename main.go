package main

import (
	"database/sql"
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

var DB *sql.DB
var files []File

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage lmao")
}

func fileBrowser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Filebrowser brrr lmao")
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmp, _ := template.ParseFiles("./static/login.html")
		tmp.Execute(w, nil)
	} else {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if username == "doggo" && password == "maracuja" {
			data := map[string]interface{}{
				"err": "Invalid credentials!",
			}
			tmp, _ := template.ParseFiles("./static/login.html")
			tmp.Execute(w, data)
		} else {
			http.Redirect(w, r, "/fileBrowser", http.StatusSeeOther)
		}

	}

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmp, _ := template.ParseFiles("./static/register.html")
		tmp.Execute(w, nil)
	} else {
		r.ParseForm()
		err := addUserDB(r.FormValue("username"), r.FormValue("password"), DB)
		if err != nil {
			tmp, _ := template.ParseFiles("./static/register.html")
			tmp.Execute(w, nil)
		}
		data := map[string]interface{}{
			"suc": "Succesfully created account! Please login.",
		}
		tmp, _ := template.ParseFiles("./static/register.html")
		tmp.Execute(w, data)
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/fileBrowser", fileBrowser)
	http.HandleFunc("/api/getFiles", getFiles)
	http.HandleFunc("/api/getFile", getFile)
	http.Handle("/static/css/style.css", http.StripPrefix("/static/css", http.FileServer(http.Dir("static/css"))))
	http.Handle("/lol/test.mkv", http.StripPrefix("/lol", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	createDB()
	DB = openDB()
	files = append(files, File{Filename: "test.mkv", Path: "/static/", Type: "mkv", Date: "11/04/2021", Size: 10003})
	files = append(files, File{Filename: "malaka.mkv", Path: "/static/", Type: "mkv", Date: "11/04/2021", Size: 10003})
	handleRequests()
	DB.Close()
}
