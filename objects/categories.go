package objects

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Category struct {
	Category string `json:"category"`
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categories []Category

	result, err := db.Query("SELECT category from categories")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var category Category
		err := result.Scan(&category.Category)
		if err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
	}

	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO tasks(categories) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	category_cat := keyVal["category"]

	_, err = stmt.Exec(category_cat)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New category was created")
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT category FROM categories WHERE category = ?", params["Category"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	var category Category
	for result.Next() {
		err := result.Scan(&category.Category)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE categories SET category = ? WHERE category = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)

	json.Unmarshal(body, &keyVal)

	newtask_name := keyVal["category"]
	_, err = stmt.Exec(newtask_name, params["category"])

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Category %s was updated", params["category"])
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM categories WHERE category = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["category"])

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Category %s was deleted", params["category"])
}
