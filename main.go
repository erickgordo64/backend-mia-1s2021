package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type cn struct{
	db *sql.DB
}

func newCn() *cn {
	return &cn{db: nil}
}

func (db *cn) abrir(){
	db.db, _ =sql.Open("godror", "HR/1234@localhost:1521/xe")

}

func (db *cn)cerrar(){
	defer db.db.Close()
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

func datos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pol := newCn()
	pol.abrir()
	rows, err := pol.db.Query("select nombre from usuario")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
		fmt.Printf("The date is: %s\n", thedate,thedate)
		json.NewEncoder(w).Encode(thedate)
	}
	pol.cerrar()
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
	router.HandleFunc("/data", datos).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
