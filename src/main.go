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
var resultadoBusca Tiendas
var eliminando mandarEliminar
var resultadoID Departamentos
var longit int
var buscar mandarBuscar
var veclineado []*Listas.ListaEnlazada

//Struct Buscar
type mandarBuscar struct {
	Departamento string `json:Departamento`
	Nombre       string `json:Nombre`
	Calificacion int    `json:Calificacion`
}

//Struct Eliminar
type mandarEliminar struct {
	Nombre       string `json:Nombre`
	Categoria    string `json:Categoria`
	Calificacion int    `json:Calificacion`
}

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
	Linealizando()
}

//EXPORTAR JSON
func exportarJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	var mandarJson *Raiz = exportarVector()
	json.NewEncoder(w).Encode(mandarJson)
}

//BuscarID
func buscarID(w http.ResponseWriter, r *http.Request) {
	cargar := mux.Vars(r)
	b, _ := strconv.Atoi(cargar["id"])
	var mensaje string
	mensaje = "No se encontraron tiendas"
	numMandar := b
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	if veclineado[numMandar] == nil {
		json.NewEncoder(w).Encode(mensaje)
	} else {
		json.NewEncoder(w).Encode(buscandoID(numMandar))
	}
}

func EliminarTienda(w http.ResponseWriter, r *http.Request) {
	var mens string = ""
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No inserto we :c")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &eliminando)
	mens = EliminandoTienda()
	if mens == "" {
		mens = "No se encontro ningun nodo"
		json.NewEncoder(w).Encode(mens)
	} else {
		json.NewEncoder(w).Encode(mens)
	}
}

func BuscarTienda(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No inserto we :c")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &buscar)
	resultadoBusca = BuscandoTienda()
	json.NewEncoder(w).Encode(resultadoBusca)
}

//Funciones utiles

//Exportar
func exportarVector() *Raiz {
	var mandarTienda Tiendas
	var MandarDepartamento Departamentos
	var mandarDatos Datos
	var multiTiendas []Tiendas
	var multiDepartamnetos []Departamentos
	var multiDatos []Datos
	var mandarRaiz Raiz
	var n int = 4
	var buscaindice int = 1
	var buscadepa int = 0
	var mandarindice string
	for i := 0; i < len(veclineado); i++ {
		fmt.Println(len(veclineado))
		var recibirlong int = veclineado[i].PasarNodoID()
		for j := 0; j < recibirlong; j++ {
			var nodorecibir *Listas.Nodo = veclineado[j].RecorrerID(NombreID)
			NombreID = nodorecibir.Nombre
			fmt.Println(NombreID)
			mandarTienda = Tiendas{Nombre: nodorecibir.Nombre, Descripcion: nodorecibir.Descripcion, Contacto: nodorecibir.Contacto, Calificacion: nodorecibir.Calificacion}
			multiTiendas = append(multiTiendas, mandarTienda)
		}
		var indicetotal int = len(mandar.Datos)
		var paso int = 0
		paso = (indicetotal * 5 * buscaindice) - 1
		if i == n {
			n = n + 5
			var aceptando bool = false
			if i == paso {
				var mandardepar string = mandar.Datos[0].Departamentos[buscadepa].Nombre
				mandarindice = mandar.Datos[(buscaindice - 1)].Indice
				MandarDepartamento = Departamentos{Nombre: mandardepar, Tiendas: multiTiendas}
				multiDepartamnetos = append(multiDepartamnetos, MandarDepartamento)
				buscaindice = buscaindice + 1
				buscadepa = buscadepa + 1
				aceptando = true
			}
			if aceptando == true {
				mandarDatos = Datos{Indice: mandarindice, Departamentos: multiDepartamnetos}
				multiDatos = append(multiDatos, mandarDatos)
				aceptando = false
				buscaindice = 1
			}
		}

	}
	mandarRaiz = Raiz{Datos: multiDatos}
	return &mandarRaiz
}

//BuscarID
var NombreID string = ""

func buscandoID(numero int) *Departamentos {
	var recibirlong int = veclineado[numero].PasarNodoID()
	var mandarTienda Tiendas
	var MandarDepartamento Departamentos
	var multiTiendas []Tiendas
	for i := 0; i < recibirlong; i++ {
		var nodorecibir *Listas.Nodo = veclineado[numero].RecorrerID(NombreID)
		NombreID = nodorecibir.Nombre
		fmt.Println(NombreID)
		mandarTienda = Tiendas{Nombre: nodorecibir.Nombre, Descripcion: nodorecibir.Descripcion, Contacto: nodorecibir.Contacto, Calificacion: nodorecibir.Calificacion}
		multiTiendas = append(multiTiendas, mandarTienda)
	}
	MandarDepartamento = Departamentos{Tiendas: multiTiendas}
	NombreID = ""
	return &MandarDepartamento
}

//Eliminar Tienda
func EliminandoTienda() string {
	var mensajemostrar string = ""
	var primero, segundo, tercero int
	var secundaria string
	secundaria = eliminando.Nombre
	for j := 0; j < len(mandar.Datos[0].Departamentos); j++ {
		if mandar.Datos[0].Departamentos[j].Nombre == eliminando.Categoria {
			var h int
			//Encotrando posicion vector
			if j == 0 {
				h = 0
			} else {
				h = j - 1
			}
			primero = j + h
			break
		}
	}
	var salir bool = false
	for i := 0; i < len(eliminando.Nombre); i++ {
		for j := 0; j < len(mandar.Datos); j++ {
			if string(secundaria[0]) == mandar.Datos[j].Indice {
				segundo = (primero * len(mandar.Datos)) + j
				salir = true
				break
			}
		}
		if salir == true {
			break
		}
	}
	var calif int
	var eliminarNombre string = eliminando.Nombre
	calif = eliminando.Calificacion
	tercero = segundo*5 + calif
	var nodorecibir string = veclineado[tercero].Eliminartienda(eliminarNombre)
	mensajemostrar = nodorecibir
	fmt.Println("El mensaje que lleva es" + mensajemostrar)
	return mensajemostrar
}

//buscar tienda
func BuscandoTienda() Tiendas {
	var primero, segundo, tercero int
	var secundaria string
	secundaria = buscar.Nombre
	for j := 0; j < len(mandar.Datos[0].Departamentos); j++ {
		if mandar.Datos[0].Departamentos[j].Nombre == buscar.Departamento {
			var h int
			//Encotrando posicion vector
			if j == 0 {
				h = 0
			} else {
				h = j - 1
			}
			primero = j + h
			break
		}
	}
	for i := 0; i < len(buscar.Nombre); i++ {
		for j := 0; j < len(mandar.Datos); j++ {
			if string(secundaria[0]) == mandar.Datos[j].Indice {
				segundo = (primero * len(mandar.Datos)) + j
				break
			}
		}
	}
	var calif int
	var buscarNombre string
	calif = buscar.Calificacion
	tercero = segundo*5 + calif
	buscarNombre = buscar.Nombre
	var nodorecibir *Listas.Nodo = veclineado[tercero].PasarNodo(buscarNombre)
	var mandarTienda Tiendas = Tiendas{Nombre: nodorecibir.Nombre, Descripcion: nodorecibir.Descripcion, Contacto: nodorecibir.Contacto, Calificacion: nodorecibir.Calificacion}
	return mandarTienda
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
	for i := 0; i < len(mandar.Datos); i++ {
		var h int
		//Encotrando posicion vector
		if i == 0 {
			h = 0
		} else {
			h = i - 1
		}
		primero = i + h
		fmt.Println("Sumo primero" + strconv.Itoa(primero))
		for j := 0; j < len(mandar.Datos[i].Departamentos); j++ {
			//Encotrando posicion vector
			segundo = (primero * len(mandar.Datos)) + j
			fmt.Println("Sumo segundo" + strconv.Itoa(segundo))
			for k := 0; k < len(mandar.Datos[i].Departamentos[j].Tiendas); k++ {
				var calif int
				calif = mandar.Datos[i].Departamentos[j].Tiendas[k].Calificacion
				tercero = segundo*5 + calif
				fmt.Println("calculo tercero" + strconv.Itoa(tercero))
				//Crear Nodo
				nodomandar := Listas.Nodo{mandar.Datos[i].Departamentos[j].Tiendas[k].Nombre, mandar.Datos[i].Departamentos[j].Tiendas[k].Descripcion, mandar.Datos[i].Departamentos[j].Tiendas[k].Contacto, mandar.Datos[i].Departamentos[j].Tiendas[k].Calificacion, nil, nil}
				//Agregar Nodo a vector
				if veclin[tercero] == nil {
					nuevalista := Listas.CrearLista()
					nuevalista.InsertarNodo(&nodomandar)
					veclin[tercero] = nuevalista
				} else {
					veclin[tercero].InsertarNodo(&nodomandar)
				}
			}
		}
	}
	veclineado = veclin
}

//main
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/cargartienda", agregar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", BuscarTienda).Methods("POST")
	router.HandleFunc("/id/{id}", buscarID).Methods("GET")
	router.HandleFunc("/guardar", exportarJson).Methods("GET")
	router.HandleFunc("/Eliminar", EliminarTienda).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
}
