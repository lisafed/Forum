package main

import (
	"log"
	"net/http"

	"Forum/core"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := core.InitDatabase("data")
	defer db.Close()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		core.CreateUser(db, w, r)
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		core.CreatePost(db, w, r)
	})

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		core.CreateComment(db, w, r)
	})

	http.HandleFunc("/likes", func(w http.ResponseWriter, r *http.Request) {
		core.CreateLike(db, w, r)
	})

	http.HandleFunc("/dislikes", func(w http.ResponseWriter, r *http.Request) {
		core.CreateDislike(db, w, r)
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
