package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/mark", markHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {

	data := map[string]interface{}{
		"attendees": getAttendeesCount(),
	}

	b, err := ioutil.ReadFile("index.html")
	if err != nil {
		serveErr(rw, req, err)
		return
	}

	t := template.Must(template.New("index").Parse(string(b)))
	t.Execute(rw, data)
}

func markHandler(rw http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	if email == "" {
		rw.Write([]byte("Invalid request"))
		return
	}
	err := markAttendance(email)
	if err != nil {
		serveErr(rw, req, err)
		return
	}
	http.Redirect(rw, req, "/", 301)
}

func serveErr(rw http.ResponseWriter, req *http.Request, err error) {
	rw.WriteHeader(505)
	rw.Write([]byte(fmt.Sprint("Internal Server Error", err)))
}
