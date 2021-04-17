package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

// Persistence
var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

type allTasks []task

func getTasks(w http.ResponseWriter, r *http.Request) { //esto sirve para mostar todos los datos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) { // esto sirve para crear tareas
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body) //funcion para recibir los datos del body que nos envia el cliente
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &newTask) //asiganmos lo que nos viene a newtask
	newTask.ID = len(tasks) + 1       //aumentamos el ID
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bienvenido a mi api")
}

func main() {
	fmt.Println("helloworld")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
