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
	idUsuario        int     `json:"idusuario"`
	estado           string  `json:"estado"`
	nombre           string  `json:"Content"`
	apellido         string  `json:"apellido"`
	correo           string  `json:"correo"`
	contrasena       string  `json:"contrasena"`
	fecha_nacimiento string  `json:"fecha_nacimiento"`
	pais             string  `json:"pais"`
	foto             string  `json:"foto"`
	creditos         float32 `json:"creditos"`
}

//Persistence
var datos = alldatos{
	{
		idUsuario:        0,
		estado:           "esto",
		nombre:           "es",
		apellido:         "una",
		correo:           "prueba",
		contrasena:       "de",
		fecha_nacimiento: "si",
		pais:             "funciona",
		foto:             "esto",
		creditos:         1.1,
	},
}

type alldatos []dato

func getTasks(w http.ResponseWriter, r *http.Request) { //esto sirve para mostar todos los datos
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getdatos(w http.ResponseWriter, r *http.Request) {

	var Data dato

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

		var ID int
		var ESTADO string
		var NOMBRE string
		var APELLIDO string
		var CORREO string
		var CONTRASENA string
		var FECHA string
		var PAIS string
		var FOTO string
		var CREDITOS string

		err := rows.Scan(&ID, &ESTADO, &NOMBRE, &APELLIDO, &CORREO, &CONTRASENA, &FECHA, &PAIS, &FOTO, &CREDITOS)
		if err != nil {
			log.Fatalln(err)
		}

		Data.idUsuario = ID

		datos = append(datos, Data)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datos)
	fmt.Printf("%#v\n", datos)
	pol.cerrar()
}

func getDataPrueba(w http.ResponseWriter, r *http.Request) { //esto sirve para mostar todos los datos
	w.Header().Set("Content-Type", "application/json")
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
	router.HandleFunc("/tasks", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
