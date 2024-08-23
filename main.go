package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getPosts(w, r)
		case http.MethodPost:
			createPost(w, r)
		}
	})

	http.HandleFunc("/posts/get", getPost)
	http.HandleFunc("/posts/update", updatePost)
	http.HandleFunc("/posts/delete", deletePost)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
