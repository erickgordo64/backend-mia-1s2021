package main

import "C"
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-yaml/yaml"
	_ "github.com/go-yaml/yaml"
	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
	"github.com/rs/cors"
)

// La primera variable es un mapa donde la clave es en realidad un puntero a un WebSocket, el valor es un booleano.
// La segunda variable es un canal que actuará como una cola de mensajes enviados por los clientes.
var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan Message)           // Broadcast channel

// Este es solo un objeto con métodos para tomar una conexión HTTP normal y actualizarla a un WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Definiremos un objeto para guardar nuestros mensajes, para interactuar con el servicio ***Gravatar*** que nos proporcionará un avatar único.
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type cn struct {
	db *sql.DB
}

func newCn() *cn {
	return &cn{db: nil}
}

func (db *cn) abrir() {
	db.db, _ = sql.Open("godror", "pruebas/1234@localhost:1521/xe")

}

func (db *cn) cerrar() {
	defer db.db.Close()
}

//estructuras para el mapeo del archivo

type resultado struct {
	Visitante int `mapstructure: visitante yaml: visitante`
	Local     int `mapstructure: local yaml: local`
}

type prediccion struct {
	Visitante int `mapstructure: visitante yaml: visitante`
	Local     int `mapstructure: local yaml: local`
}

type predicciones struct {
	Deporte    string     `mapstructure: deporte yaml: deporte`
	Fecha      string     `mapstructure: fecha yaml: fecha`
	Visitante  string     `mapstructure: visitante yaml: visitante`
	Local      string     `mapstructure: local yaml: local`
	Prediccion prediccion `mapstructure: prediccion yaml: prediccion`
	Resultado  resultado  `mapstructure: resultado yaml: resultado`
}

type jornadas struct {
	Jornada      string         `mapstructure: jornada yaml: jornada`
	Predicciones []predicciones `mapstructure: predicciones yaml: predicciones`
}

type resultados struct {
	Temporada string     `mapstructure: temporada yaml: temporada`
	Tier      string     `mapstructure: tier yaml: tier`
	Jornadas  []jornadas `mapstructure: jornadas yaml: jornadas`
}

type Archivo struct {
	Nombre     string       `mapstructure: nombre yaml: nombre`
	Apellido   string       `mapstructure: apellido yaml: apellido`
	Password   string       `mapstructure: password yaml: password`
	Username   string       `mapstructure: username yaml: username`
	Resultados []resultados `mapstructure:  resultados yaml: resultados`
}

//----------------------aqui terminan las estruturas del mapeo

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
	{
		ID:      2,
		Name:    "Task two",
		Content: "Some Content",
	},
	{
		ID:      3,
		Name:    "Task three",
		Content: "Some Content",
	},
	{
		ID:      4,
		Name:    "Task four",
		Content: "Some Content",
	}, {
		ID:      5,
		Name:    "Task five",
		Content: "Some Content",
	},
}

type allTasks []task

type usser struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// Persistence
var ussers = allussers{}

type allussers []usser

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

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var User usser
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}
	json.Unmarshal(reqBody, &User)

	pol := newCn()
	pol.abrir()
	rows, err := pol.db.Query("select idusuario, username, password from usuario where username=:1 and password=:2", User.Username, User.Password)

	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&User.ID, &User.Username, &User.Password)
		if err != nil {
			log.Fatalln(err)
		}
	}

	json.NewEncoder(w).Encode(User)
	pol.cerrar()

}

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
	//w.Header().Set("Access-Control-Allow-Headers:", "*")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "*")
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

func uploader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(2000)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, fileinfo, err := r.FormFile("archivo") //recibo el archivo de un form con su clave o parametro key es archivo

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	f, err := os.OpenFile("./file/"+fileinfo.Filename, os.O_WRONLY|os.O_CREATE, 0666) //obtngo el archivo

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	defer f.Close()

	io.Copy(f, file) //copio el archivo en file a f

	//fmt.Fprintf(w, fileinfo.Filename)

	raw, err := ioutil.ReadFile("./file/" + fileinfo.Filename) //leo el archivo

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var dat map[string]interface{}

	if err := yaml.Unmarshal(raw, &dat); err != nil { //reconozco el archivo con el yaml.unmarshal
		panic(err)
	}

	var arch *Archivo
	//var result *resultados

	sqlStatement := `INSERT INTO masiva(CODIGO_CLIENTE, NOMBRE_CLIENTE, APELLIDO_CLIENTE,PASSWORD,USERNAME, TEMPORADA, TIER, JORNADA, DEPORTE, FECHA, VISITANTE_NOMBRE, LOCAL_NOMBRE, VISITANTE_PREDICCION, LOCAL_PREDICCION, VISITANTE_RESULTADO, LOCAL_RESULTADO) values (:1, :2,:3,:4,:5,:6,:7, :8, :9, TO_DATE(:10,'DD/MM/YYYY HH24:MI'), :11, :12, :13, :14, :15, :16)`
	pol := newCn()
	pol.abrir()

	for key := range dat { //mapeo el archivo, con el primer for se puede llenar la tabla usuario
		fmt.Println(key)
		mapstructure.Decode(dat[key], &arch)
		// con estos for lleno las tablas
		for i := 0; i < len(arch.Resultados); i++ {
			for j := 0; j < len(arch.Resultados[i].Jornadas); j++ {
				for k := 0; k < len(arch.Resultados[i].Jornadas[j].Predicciones); k++ {
					//fmt.Println(key+"-"+arch.Nombre+"-"+arch.Apellido+"-"+arch.Username+"-"+arch.Password+"-"+arch.Resultados[i].Temporada+"-"+arch.Resultados[i].Tier+"-"+arch.Resultados[i].Jornadas[j].Jornada+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Deporte+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Fecha+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Visitante+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Local+"-", arch.Resultados[0].Jornadas[0].Predicciones[0].Prediccion.Visitante, "-", arch.Resultados[0].Jornadas[0].Predicciones[0].Prediccion.Local, "-", arch.Resultados[0].Jornadas[0].Predicciones[0].Resultado.Visitante, "-", arch.Resultados[0].Jornadas[0].Predicciones[0].Resultado.Visitante)
					_, err = pol.db.Exec(sqlStatement, key, arch.Nombre, arch.Apellido, arch.Password, arch.Username, arch.Resultados[i].Temporada, arch.Resultados[i].Tier, arch.Resultados[i].Jornadas[j].Jornada, arch.Resultados[i].Jornadas[j].Predicciones[k].Deporte, arch.Resultados[i].Jornadas[j].Predicciones[k].Fecha, arch.Resultados[i].Jornadas[j].Predicciones[k].Visitante, arch.Resultados[i].Jornadas[j].Predicciones[k].Local, arch.Resultados[0].Jornadas[0].Predicciones[0].Prediccion.Visitante, arch.Resultados[0].Jornadas[0].Predicciones[0].Prediccion.Local, arch.Resultados[0].Jornadas[0].Predicciones[0].Resultado.Visitante, arch.Resultados[0].Jornadas[0].Predicciones[0].Resultado.Visitante)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
	pol.cerrar()
	json.NewEncoder(w).Encode(arch)

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "bienvenido a mi api")
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Para cerrar la conexión una vez termina la función
	defer ws.Close()

	// Registramos nuestro nuevo cliente al agregarlo al mapa global de "clients" que fue creado anteriormente.
	clients[ws] = true

	// Bucle infinito que espera continuamente que se escriba  un nuevo mensaje en el WebSocket, lo desserializa de JSON a un objeto Message y luego lo arroja al canal de difusión.
	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		// Si hay un error, registramos ese error y eliminamos ese cliente de nuestro mapa global de clients
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		// Send the newly received message to the broadcast channel
		broadcast <- msg

		reader(ws)
	}

	/*log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	reader(ws)*/
}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		//fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func Archiv(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	http.ServeFile(w, r, "./public/p.png")
}

func main() {
	fmt.Println("helloworld")

	//go handleMessages()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/data", getdatos).Methods("GET")
	router.HandleFunc("/datas", getDataPrueba).Methods("GET")
	router.HandleFunc("/categorias", getCategorias).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/archivo", uploader).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/ws", handleConnection).Methods("POST")
	router.Handle("/public/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	router.HandleFunc("/down", Archiv)

	log.Fatal(http.ListenAndServe(":4000", c.Handler(router)))
}
