package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	ID       int
	Username string
	Password string
}

type Mensaje struct {
	Mensaje string
}

var lista_usuarios []User

func mensaje_inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la API de golang")
}

func agregarUsuario(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var nuevoUsuario User

	json.NewDecoder(request.Body).Decode(&nuevoUsuario)

	lista_usuarios = append(lista_usuarios, nuevoUsuario)

	respuesta := Mensaje{Mensaje: "Usuario agregado correctamente"}

	output, _ := json.Marshal(respuesta)

	fmt.Fprint(response, string(output))
}

func mostrarUsuarios(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	output, _ := json.Marshal(lista_usuarios)

	fmt.Fprint(response, string(output))
}

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/", mensaje_inicial).Methods("GET")
	mux.HandleFunc("/agregarUsuario", agregarUsuario).Methods("POST")
	mux.HandleFunc("/mostrarUsuarios", mostrarUsuarios).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("El servidor esta corriendo en el puerto 3000")
	fmt.Println("http://localhost:3000/api")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(mux)))
}
