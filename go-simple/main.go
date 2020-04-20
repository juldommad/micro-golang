package main

import (
	"encoding/json"
	"go-simple/persona"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPeople(w http.ResponseWriter, req *http.Request) {
	people := persona.TraerPersonas()
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	resultado := persona.TraePersona(params["nif"])
	json.NewEncoder(w).Encode(&resultado)
}

func CreatePerson(w http.ResponseWriter, req *http.Request) {
	var post persona.Person
	_ = json.NewDecoder(req.Body).Decode(&post)
	respuesta := persona.CrearPersona(post.NIF, post.Nombre, post.Apellido, post.Direccion.Calle, post.Direccion.Ciudad)
	json.NewEncoder(w).Encode(&respuesta)
}
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	respuesta := persona.EliminarPersona(params["nif"])
	json.NewEncoder(w).Encode(&respuesta)
}

func main() {
	persona.IniciarBase()
	router := mux.NewRouter() // enroutador que sirve para definir nuestras rutas
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{nif}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{nif}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
