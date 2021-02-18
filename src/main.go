package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"./Listas"

	"github.com/gorilla/mux"
)

//variables
var mandar Raiz
var longit int

//Struct Json
type Raiz struct {
	Datos []Datos `json:"Datos"`
}

type Datos struct {
	Indice        string          `json:Indice`
	Departamentos []Departamentos `json:"Departamentos"`
}

type Departamentos struct {
	Nombre  string    `json:Nombre`
	Tiendas []Tiendas `json:"Tiendas"`
}

type Tiendas struct {
	Nombre       string `json:Nombre`
	Descripcion  string `json:Descripcion`
	Contacto     string `json:Contacto`
	Calificacion int    `json:Calificacion`
}

//funciones
//Server
func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Si jala we uwu")
}

func agregar(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No inserto we :c")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &mandar)
	json.NewEncoder(w).Encode(mandar)
}

func Mostrar(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < len(mandar.Datos); i++ {
		fmt.Println("Indice: " + mandar.Datos[i].Indice)
		for j := 0; j < len(mandar.Datos[i].Departamentos); j++ {
			fmt.Println("Nombre: " + mandar.Datos[i].Departamentos[j].Nombre)
			for k := 0; k < len(mandar.Datos[i].Departamentos[j].Tiendas); k++ {
				fmt.Println("Nombre: " + mandar.Datos[i].Departamentos[j].Tiendas[k].Nombre)
				fmt.Println("Calificacion: " + strconv.Itoa(mandar.Datos[i].Departamentos[j].Tiendas[k].Calificacion))
				fmt.Println("Contacto: " + mandar.Datos[i].Departamentos[j].Tiendas[k].Contacto)
				fmt.Println("Descripcion: " + mandar.Datos[i].Departamentos[j].Tiendas[k].Descripcion)
			}
		}
	}
	json.NewEncoder(w).Encode(mandar)

}

//CrearVector
func longector(x int, y int) int {
	x = len(mandar.Datos)
	y = len(mandar.Datos[0].Departamentos)
	longit = 5 * x * y
	return longit
}

//vector
var veclin = make([]*Listas.ListaEnlazada, longit)

func InsertarNodo() {
	var primero, segundo, tercero int
	var nodomandar *Listas.Nodo
	for i := 0; i < len(mandar.Datos); i++ {
		var h int
		//Encotrando posicion vector
		if i == 0 {
			h = 0
		} else {
			h = i - 1
		}
		primero = i + h

		for j := 0; j < len(mandar.Datos[i].Departamentos); j++ {
			//Encotrando posicion vector
			segundo = (primero * len(mandar.Datos)) + j
			for k := 0; k < len(mandar.Datos[i].Departamentos[j].Tiendas); k++ {
				//Crear Nodo
				nodomandar = Listas.NuevoNodo(mandar.Datos[i].Departamentos[j].Tiendas[k].Nombre, mandar.Datos[i].Departamentos[j].Tiendas[k].Descripcion, mandar.Datos[i].Departamentos[j].Tiendas[k].Contacto, mandar.Datos[i].Departamentos[j].Tiendas[k].Calificacion)
				//Posicionar Nodo
				var calif int
				calif = mandar.Datos[i].Departamentos[j].Tiendas[k].Calificacion
				tercero = segundo*5 + calif
				println(tercero)
				//Agregar Nodo a vector
				veclin[tercero] = Listas.InsertarNodo(nodomandar)
				println(veclin[tercero])
			}
		}
	}
}

//main
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/agregar", agregar).Methods("POST")
	router.HandleFunc("/Mostrar", Mostrar).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))

}
