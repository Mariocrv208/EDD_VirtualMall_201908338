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
	Linealizando()
}

//CrearVector
func longector() int {
	var x int
	var y int
	x = len(mandar.Datos)
	y = len(mandar.Datos[0].Departamentos)
	longit = 5 * x * y
	return longit
}

//Linealizando vector
func Linealizando() {
	//vector
	var veclin = make([]*Listas.ListaEnlazada, longector())
	var primero, segundo, tercero int
	fmt.Println("Entro a linealizar")
	for i := 0; i < len(mandar.Datos); i++ {
		var h int
		//Encotrando posicion vector
		if i == 0 {
			h = 0
		} else {
			h = i - 1
		}
		primero = i + h
		fmt.Println("Sumo primero")
		for j := 0; j < len(mandar.Datos[i].Departamentos); j++ {
			//Encotrando posicion vector
			segundo = (primero * len(mandar.Datos)) + j

			for k := 0; k < len(mandar.Datos[i].Departamentos[j].Tiendas); k++ {
				var calif int
				calif = mandar.Datos[i].Departamentos[j].Tiendas[k].Calificacion
				tercero = segundo*5 + calif
				//Crear Nodo
				nodomandar := Listas.Nodo{mandar.Datos[i].Departamentos[j].Tiendas[k].Nombre, mandar.Datos[i].Departamentos[j].Tiendas[k].Descripcion, mandar.Datos[i].Departamentos[j].Tiendas[k].Contacto, mandar.Datos[i].Departamentos[j].Tiendas[k].Calificacion, nil, nil}
				fmt.Println("Creo el nodo")
				//Posicionar Nodo
				fmt.Println(nodomandar)
				fmt.Println("calculo tercero")
				fmt.Println(tercero)
				for m := 0; m < len(veclin); m++ {
					fmt.Println(veclin[m])
				}
				fmt.Println(veclin[tercero])
				//Agregar Nodo a vector
				if veclin[tercero] == nil {
					fmt.Println("creo lista e inserto")
					nuevalista := Listas.CrearLista()
					nuevalista.InsertarNodo(&nodomandar)
					veclin[tercero] = nuevalista
				} else {
					fmt.Println("inserto en lista existente")
					veclin[tercero].InsertarNodo(&nodomandar)
				}

				fmt.Println(veclin[tercero])
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
