package main

import (
	"encoding/json"
	"fmt"
	"go-simple/persona"
	"go-simple/token"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type myLoginStruct struct {
	UserNIF  string `json:"nif"`
	UserPass string `json:"pass"`
}

func getPeople(w http.ResponseWriter, req *http.Request) {
	people := persona.TraerPersonas()
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	resultado, errorPersona := persona.TraePersona(params["nif"])
	if errorPersona {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(&resultado)
}

func createPerson(w http.ResponseWriter, req *http.Request) {
	var post persona.Person
	_ = json.NewDecoder(req.Body).Decode(&post)
	respuesta := persona.CrearPersona(post.NIF, post.Nombre, post.Apellido, post.Direccion.Calle, post.Direccion.Ciudad)
	json.NewEncoder(w).Encode(&respuesta)
}

func deletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	respuesta := persona.EliminarPersona(params["nif"])
	json.NewEncoder(w).Encode(&respuesta)
}

func login(w http.ResponseWriter, req *http.Request) {
	var post myLoginStruct
	_ = json.NewDecoder(req.Body).Decode(&post)
	fmt.Println("\n********************")
	fmt.Printf("My nif is: %s", post.UserNIF)
	fmt.Printf("\nMy pass is: %s\n", post.UserPass)
	fmt.Println("********************")
	respuesta := ""
	isLoginOK := persona.Login(post.UserNIF, post.UserPass)
	if isLoginOK {
		err := false
		respuesta, err = token.GetToken(post.UserNIF)
		if err {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
	json.NewEncoder(w).Encode(&respuesta)
}

func main() {
	persona.IniciarBase()
	router := mux.NewRouter() // enroutador que sirve para definir nuestras rutas
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/people/{nif}", getPerson).Methods("GET")
	router.HandleFunc("/people", createPerson).Methods("POST")
	router.HandleFunc("/people/{nif}", deletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
