package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"golang.org/x/net/websocket"
)

func main() {
	log.Println("[INFO] Start server")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	s := newServer()
	//go s.ping()

	// new DB
	db := newDB()
	s.db = db
	defer db.Close()

	http.Handle("/ws", websocket.Handler(s.handlerWS))

	// FS() is created by esc and returns a http.Filesystem.
	http.Handle("/static/", http.FileServer(FS(false)))

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
