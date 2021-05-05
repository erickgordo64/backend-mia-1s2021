package main

import "C"
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"

	"github.com/go-yaml/yaml"
	_ "github.com/go-yaml/yaml"
	_ "github.com/godror/godror"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
	"github.com/rs/cors"
)

// La primera variable es un mapa donde la clave es en realidad un puntero a un WebSocket, el valor es un booleano.
// La segunda variable es un canal que actuará como una cola de mensajes enviados por los clientes.         // Broadcast channel

// Este es solo un objeto con métodos para tomar una conexión HTTP normal y actualizarla a un WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan Message
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Message),
	}
}

func (h *Hub) run() {
	for {
		select {
		case message := <-h.broadcast:
			for client := range h.clients {
				if err := client.WriteJSON(message); err != nil {
					log.Printf("error occurred: %v", err)
				}
			}
		}
	}
}

//Definiremos un objeto para guardar nuestros mensajes, para interactuar con el servicio ***Gravatar*** que nos proporcionará un avatar único.
type Message struct {
	Message string `json:"message"`
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

//estructura correo

type correo struct {
	Correo string `json:"correo"`
}

var RCorreo = CorreoR{}

type CorreoR []correo

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

	sqlStatement := `INSERT INTO masiva(ID, NOMBRE_CLIENTE, APELLIDO_CLIENTE,PASSWORD,USERNAME, TEMPORADA, TIER, JORNADA, DEPORTE, FECHA, VISITANTE_NOMBRE, LOCAL_NOMBRE, VISITANTE_PREDICCION, LOCAL_PREDICCION, VISITANTE_RESULTADO, LOCAL_RESULTADO) values (:1, :2,:3,:4,:5,:6,:7, :8, :9, TO_DATE(:10,'DD/MM/YYYY HH24:MI'), :11, :12, :13, :14, :15, :16)`
	pol := newCn()
	pol.abrir()

	for key := range dat { //mapeo el archivo, con el primer for se puede llenar la tabla usuario
		//fmt.Println(key)
		mapstructure.Decode(dat[key], &arch)
		// con estos for lleno las tablas
	}

	for key := range dat {
		fmt.Println(key)

		for i := 0; i < len(arch.Resultados); i++ {
			fmt.Println("	" + arch.Resultados[i].Temporada)
			for j := 0; j < len(arch.Resultados[i].Jornadas); j++ {
				fmt.Println("		" + arch.Resultados[i].Jornadas[j].Jornada)
				for k := 0; k < len(arch.Resultados[i].Jornadas[j].Predicciones); k++ {
					fmt.Println("			" + arch.Resultados[i].Jornadas[j].Predicciones[k].Deporte + "-" + arch.Resultados[i].Jornadas[j].Predicciones[k].Fecha)
					//fmt.Println("codigo cliente: "+key+"-nombre cliente: "+arch.Nombre+"-apellido cliente: "+arch.Apellido+"- username: "+arch.Username+"-"+arch.Password+"- temporada: "+arch.Resultados[i].Temporada+"- tier: "+arch.Resultados[i].Tier+"- Jornada: "+arch.Resultados[i].Jornadas[j].Jornada+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Deporte+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Fecha+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Visitante+"-"+arch.Resultados[i].Jornadas[j].Predicciones[k].Local+"-", arch.Resultados[i].Jornadas[j].Predicciones[k].Prediccion.Visitante, "-", arch.Resultados[i].Jornadas[j].Predicciones[k].Prediccion.Local, "-", arch.Resultados[i].Jornadas[j].Predicciones[k].Resultado.Visitante, "-", arch.Resultados[i].Jornadas[j].Predicciones[k].Resultado.Visitante)
					_, err = pol.db.Exec(sqlStatement, key, arch.Nombre, arch.Apellido, arch.Password, arch.Username, arch.Resultados[i].Temporada, arch.Resultados[i].Tier, arch.Resultados[i].Jornadas[j].Jornada, arch.Resultados[i].Jornadas[j].Predicciones[k].Deporte, arch.Resultados[i].Jornadas[j].Predicciones[k].Fecha, arch.Resultados[i].Jornadas[j].Predicciones[k].Visitante, arch.Resultados[i].Jornadas[j].Predicciones[k].Local, arch.Resultados[i].Jornadas[j].Predicciones[k].Prediccion.Visitante, arch.Resultados[i].Jornadas[j].Predicciones[k].Prediccion.Local, arch.Resultados[i].Jornadas[j].Predicciones[k].Resultado.Visitante, arch.Resultados[i].Jornadas[j].Predicciones[k].Resultado.Local)
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

var hub = NewHub()

func Socket(w http.ResponseWriter, r *http.Request) {
	// Start a go routine
	go hub.run()
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		delete(hub.clients, ws)
		ws.Close()
		log.Printf("Closed!")
	}()

	hub.clients[ws] = true

	log.Println("Connected!")

	read(hub, ws)
	return

}

func read(hub *Hub, client *websocket.Conn) {
	for {
		var message Message
		err := client.ReadJSON(&message)
		if err != nil {
			log.Printf("error occurred: %v", err)
			delete(hub.clients, client)
			break
		}
		log.Println(message)

		hub.broadcast <- message
	}
}

func Archiv(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	http.ServeFile(w, r, "./public/p.png")
}

func RecCorreo(w http.ResponseWriter, r *http.Request) {
	var RCorreo = CorreoR{}
	var Remail correo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Datos no validos")
	}
	json.Unmarshal(reqBody, &Remail)
	RCorreo = append(RCorreo, Remail)
	fmt.Println(Remail.Correo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Remail)

	const letterBytes = "abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 9)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	var a string = string(b) + "1"
	fmt.Println("el resultad es", a)
	/*
		pol := newCn()
		pol.abrir()
		sqlStatement := `UPDATE cliente set pass=:1 where correo_electronico=:2`
		_, err = pol.db.Exec(sqlStatement, a, Remail.Correo)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("Se actualizo la contrasena del usuario")
		}
		pol.cerrar()*/

	auth := smtp.PlainAuth("", "miaproyecto4319@gmail.com", "2006001077", "smtp.gmail.com")
	to := []string{Remail.Correo} // Array de correos de destino

	msg := []byte("From: QuinielaApp ✉ <miaproyecto4319@gmail.com> \r\n" +
		"Subject: Cambio de Contrasena\r\n" +
		"\r\n" +
		"Esta es su nueva contrasena " + a + " por favor cambiarla en su perfil.\r\n")

	error := smtp.SendMail("smtp.gmail.com:587", auth, "miaproyecto4319@gmail.com", to, msg)
	if error != nil {
		fmt.Println("Informamos el error", error)
	}

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
	router.HandleFunc("/ws", Socket).Methods("GET")
	router.HandleFunc("/reccontra", RecCorreo).Methods("PUT")
	router.Handle("/public/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	router.HandleFunc("/down", Archiv)

	log.Fatal(http.ListenAndServe(":4000", c.Handler(router)))
}
