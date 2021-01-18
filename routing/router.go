package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskmanagerAPI/objects"
)

func Route() {

	router := mux.NewRouter()
	router.HandleFunc("/posts", objects.GetPosts).Methods("GET")
	router.HandleFunc("/posts", objects.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", objects.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", objects.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", objects.DeletePost).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
