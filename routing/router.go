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

	router.HandleFunc("/categories", objects.GetCategories).Methods("GET")
	router.HandleFunc("/categories", objects.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{category_name}", objects.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{category_name}", objects.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{category_name}", objects.DeleteCategory).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
