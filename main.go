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

type cn struct {
	db *sql.DB
}

func newCn() *cn {
	return &cn{db: nil}
}

func (db *cn) abrir() {
	db.db, _ = sql.Open("godror", "HR/1234@localhost:1521/xe")

}

func (db *cn) cerrar() {
	defer db.db.Close()
}

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

type dato struct {
	ID              int     `json:"ID"`
	Estado          string  `json:"Estado"`
	Nombre          string  `json:"Nombre"`
	Apellido        string  `json:"Apellido"`
	Correo          string  `json:"Correo"`
	Contrasena      string  `json:"Contrasena"`
	Fechanacimiento string  `json:"Fechanacimiento"`
	Pais            string  `json:"Pais"`
	Foto            string  `json:"Foto"`
	Creditos        float32 `json:"Creditos"`
}

//Persistence

type alldatos []dato

type categoria struct {
	ID        int    `json:"ID"`
	CATEGORIA string `json:"CATEGORIA"`
}
type allcategorias []categoria

func getCategorias(w http.ResponseWriter, r *http.Request) { //esto sirve para mostar todos los datos
	w.Header().Set("Content-Type", "application/json")

	var categorias = allcategorias{}

	var Cat categoria
	pol := newCn()
	pol.abrir()
	rows, err := pol.db.Query("select * from categoria")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&Cat.ID, &Cat.CATEGORIA)
		if err != nil {
			log.Fatalln(err)
		}
		categorias = append(categorias, Cat)
	}

	json.NewEncoder(w).Encode(categorias)
	pol.cerrar()
}

func getTasks(w http.ResponseWriter, r *http.Request) { //esto sirve para mostar todos los datos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getdatos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers:", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var Data dato
	var datos = alldatos{}
	pol := newCn()
	pol.abrir()
	rows, err := pol.db.Query("select * from usuario")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&Data.ID, &Data.Estado, &Data.Nombre, &Data.Apellido, &Data.Correo, &Data.Contrasena, &Data.Fechanacimiento, &Data.Pais, &Data.Foto, &Data.Creditos)
		if err != nil {
			log.Fatalln(err)
		}

		datos = append(datos, Data)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos)
	fmt.Printf("entro")
	pol.cerrar()
}

func getDataPrueba(w http.ResponseWriter, r *http.Request) { //esto sirve para mostar todos los datos
	w.Header().Set("Content-Type", "application/json")
	var datos = alldatos{}
	json.NewEncoder(w).Encode(datos)
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
	router.HandleFunc("/data", getdatos).Methods("GET")
	router.HandleFunc("/datas", getDataPrueba).Methods("GET")
	router.HandleFunc("/categorias", getCategorias).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", router))
}
