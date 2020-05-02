package persona

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	NIF       string     `gorm:"PRIMARY_KEY"`
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Pass      string     `json:"pass"`
	Direccion *Direccion `json:"direccion"`
}
type Direccion struct {
	Calle  string `json:"calle"`
	Ciudad string `json:"ciudad"`
}

var db *gorm.DB

//IniciarBase funcion que sirve para iniciar la base de datos
func IniciarBase() {
	fmt.Printf("\nCreamos la base de datos")
	var err error
	db, err = gorm.Open("sqlite3", "personas.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Person{})
	fmt.Printf("\nSe migra el estilo")
	//defer db1.Close()
}
func TraerPersonas() []Person {
	// Read
	var resultado []Person
	db.Find(&resultado) // find all persons with id 1
	return resultado
}

func TraePersona(Nif string) (Person, bool) {
	// Read
	var resultado Person
	db.Find(&resultado, "nif = ?", Nif) // find all persons with id 1
	if resultado.NIF != "" {
		return resultado, false
	}
	return resultado, true
}
func CrearPersona(nif string, nombre string, apellido string, calle string, ciudad string) Person {
	direccion := Direccion{Calle: calle, Ciudad: ciudad}
	persona := Person{NIF: nif, Nombre: nombre, Apellido: apellido, Direccion: &direccion}
	// Read
	db.Create(&persona)
	return persona
}
func EliminarPersona(nif string) int {
	// Read
	fmt.Printf("\nElimino nif %s", nif)
	persona, _ := TraePersona(nif)
	fmt.Printf("\nElimino nif %v", persona)
	db.Delete(&persona)
	codigo := 200
	return codigo
}
func Login(user string, password string) bool {

	persona, errPersona := TraePersona(user)
	expectedPassword := persona.Pass

	// Si no esta en la base de datos, o la contraseña no corresponde devolveremos un status de "Unauthorized" status
	if expectedPassword != password || errPersona {
		return false
	}
	return true
}
func UpdatePersona(persona Person, newNombre string) {
	// Update - update product's price to 2000
	db.Model(&persona).Update("Nombre", newNombre)
}
