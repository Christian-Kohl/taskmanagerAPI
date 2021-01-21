package objects

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"taskmanagerAPI/dto"
)

type Post struct {
	Task_id   string `json:"task_id"`
	Task_name string `json:"task_name"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []dto.Post

	result, err := db.Query("SELECT task_id, task_name from tasks")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {
		var post dto.Post
		err := result.Scan(&post.Task_id, &post.Task_name)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}

	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO tasks(task_name) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	task_name := keyVal["task_name"]

	_, err = stmt.Exec(task_name)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New post was created")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT task_id, task_name FROM tasks WHERE task_id = ?", params["Task_id"])
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()
	var post dto.Post
	for result.Next() {
		err := result.Scan(&post.Task_id, &post.Task_name)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE tasks SET task_name = ? WHERE task_id = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)

	json.Unmarshal(body, &keyVal)

	newtask_name := keyVal["task_name"]
	_, err = stmt.Exec(newtask_name, params["task_id"])

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with task_id = %s was updated", params["task_id"])
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM tasks WHERE task_id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["task_id"])

	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post with task_id = %s was deleted", params["task_id"])
}
