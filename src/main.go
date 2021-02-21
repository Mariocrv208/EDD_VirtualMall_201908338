package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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
	file, _ := json.MarshalIndent(mandarJson, "", " ")
	_ = ioutil.WriteFile("201908338_guardar.json", file, 0644)
}

//Graficando
func graficando(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	var mensaje string = Grafico()
	json.NewEncoder(w).Encode(mensaje)
}

//Funciones utiles
//BuscarID
func buscarID(w http.ResponseWriter, r *http.Request) {
	cargar := mux.Vars(r)
	b, _ := strconv.Atoi(cargar["id"])
	var mensaje string
	mensaje = "No se encontraron tiendas"
	numMandar := b
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	if veclin[numMandar] == nil {
		json.NewEncoder(w).Encode(mensaje)
	} else {
		json.NewEncoder(w).Encode(buscandoID(numMandar))
	}
}

//eliminar tienda
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

//buscar especifico
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

//Reportes
func Grafico() string {
	fmt.Println("Entro a graficar")
	//Encabezados Graf
	archivo, _ := os.Create("graficoLinealizado.dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	var califindice = 1
	//Nodos base
	_, _ = archivo.WriteString("subgraph cluster0{" + "\n")
	_, _ = archivo.WriteString("edge[minlen=0.1, dir=fordware]" + "\n")
	_, _ = archivo.WriteString("cabecera [label=\"cabecera\", shape=none]" + "\n")
	for i := 0; i < len(veclin); i++ {
		if califindice == 6 {
			califindice = 1
		}
		_, _ = archivo.WriteString("struct" + strconv.Itoa(i) + "[shape=record,label=\"...|...|{ " + strconv.Itoa(califindice) + " | pos:" + strconv.Itoa(i) + "}\"];" + "\n")
		califindice++
	}
	_, _ = archivo.WriteString("cabecera -> struct0" + "\n")
	for i := 0; i < len(veclin); i++ {
		if i+1 == len(veclin) {
			_, _ = archivo.WriteString("struct" + strconv.Itoa(i) + " -> null;" + "\n")
		} else {
			_, _ = archivo.WriteString("struct" + strconv.Itoa(i) + " -> struct" + strconv.Itoa(i+1) + ";" + "\n")
		}
	}
	_, _ = archivo.WriteString("}" + "\n")
	fmt.Println("Termino nodo base")
	//Nodos Tiendas
	var cont int = 0
	var mult int = 0
	var multmulti2 []int
	var multimult []int
	for i := 0; i < len(veclin); i++ {
		_, _ = archivo.WriteString("subgraph cluster" + strconv.Itoa(i+1) + "{" + "\n")
		_, _ = archivo.WriteString("edge[dir=both]" + "\n")
		var recibirlong int
		if veclin[i] != nil {
			recibirlong = veclin[i].PasarNodoID()
			fmt.Println("longitud lista en nodo")
			fmt.Println(recibirlong)
		}
		var pas int
		if recibirlong != 0 {
			mult = i
			var buscanombre string = ""
			pas = cont
			multimult = append(multimult, mult)
			multmulti2 = append(multmulti2, pas)
			var permiso int = 1
			for j := 0; j < recibirlong; j++ {
				var nodorecibir *Listas.Nodo = veclin[i].RecorrerID(buscanombre)
				buscanombre = nodorecibir.Nombre
				var nom string = nodorecibir.Nombre
				var con string = nodorecibir.Contacto
				_, _ = archivo.WriteString("nodo" + strconv.Itoa(cont) + " [shape=record,label=\"{ " + nom + " | " + con + "}\"];" + "\n")
				if permiso == 2 {
					_, _ = archivo.WriteString("nodo" + strconv.Itoa(cont-1) + " -> nodo" + strconv.Itoa(cont) + ";" + "\n")
					permiso = 0
				}
				permiso++
				cont++
				pas++
				fmt.Println(nodorecibir.Nombre)
				fmt.Println("CargoTiendas")
			}
		}
		_, _ = archivo.WriteString("}" + "\n")
	}
	fmt.Println("Termino")
	for i := 0; i < len(veclin); i++ {
		for j := 0; j < len(multimult); j++ {
			if multimult[j] == i {
				_, _ = archivo.WriteString("struct" + strconv.Itoa(i) + " -> nodo" + strconv.Itoa(multmulti2[j]) + " [lhead=cluster" + strconv.Itoa(i+1) + "];" + "\n")
			}
		}

	}
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./graficoLinealizado.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("grafica.pdf", cmd, os.FileMode(mode))
	return "Grafico listo"
}

//Exportar
func exportarVector() *Raiz {
	fmt.Println("EntroGuardar")
	var mandarTienda Tiendas
	var MandarDepartamento Departamentos
	var mandarDatos Datos
	var multiTiendas []Tiendas
	var multiDepartamnetos []Departamentos
	var multiDatos []Datos
	var mandarRaiz Raiz
	var indicebuscando int = 1
	var buscadeparta int = 1
	var paso int = len(veclin)
	var contador int = 0
	for i := 0; i < len(mandar.Datos); i++ {
		for j := 0; j < len(mandar.Datos[i].Departamentos); j++ {
			for k := contador; k < paso; k += 5 {
				//para tiendas
				var recibirlong int = 0
				if veclin[k] != nil {
					recibirlong = veclin[k].PasarNodoID()
					fmt.Println("longitud lista en nodo")
					fmt.Println(recibirlong)
				}
				if recibirlong != 0 {
					var buscanombre string = ""
					for j := 0; j < recibirlong; j++ {
						var nodorecibir *Listas.Nodo = veclin[k].RecorrerID(buscanombre)
						buscanombre = nodorecibir.Nombre
						fmt.Println(nodorecibir.Nombre)
						mandarTienda = Tiendas{Nombre: nodorecibir.Nombre, Descripcion: nodorecibir.Descripcion, Contacto: nodorecibir.Contacto, Calificacion: nodorecibir.Calificacion}
						multiTiendas = append(multiTiendas, mandarTienda)
						fmt.Println("CargoTiendas")
					}
				}
				if k == paso-1 {
					paso = paso - 1
					contador = contador + 1
				}
				if contador == 4 {
					break
				}

			}
			MandarDepartamento = Departamentos{Nombre: mandar.Datos[0].Departamentos[buscadeparta-1].Nombre, Tiendas: mandar.Datos[0].Departamentos[buscadeparta-1].Tiendas}
			multiDepartamnetos = append(multiDepartamnetos, MandarDepartamento)
		}
		mandarDatos = Datos{Indice: mandar.Datos[indicebuscando-1].Indice, Departamentos: mandar.Datos[indicebuscando-1].Departamentos}
		multiDatos = append(multiDatos, mandarDatos)
		indicebuscando = indicebuscando + 1
	}

	mandarRaiz = Raiz{Datos: multiDatos}
	return &mandarRaiz
}

//BuscarID
var NombreID string = ""

func buscandoID(numero int) *Departamentos {
	var recibirlong int = veclin[numero].PasarNodoID()
	var mandarTienda Tiendas
	var MandarDepartamento Departamentos
	var multiTiendas []Tiendas
	fmt.Println(recibirlong)
	for i := 0; i < recibirlong; i++ {
		var nodorecibir *Listas.Nodo = veclin[numero].RecorrerID(NombreID)
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
	var primero, segundo, tercero int
	for j := 0; j < len(mandar.Datos[0].Departamentos); j++ {
		if mandar.Datos[0].Departamentos[j].Nombre == eliminando.Categoria {
			var h int
			var ps int
			//Encotrando posicion vector
			if j == 0 {
				h = 0
				ps = 0
			} else if j == 1 {
				h = 0
				ps = 1
			} else {
				h = 1
				ps = j + 1
			}
			primero = ps - h
			break
		}
	}
	for i := 0; i < len(eliminando.Nombre); i++ {
		for j := 0; j < len(mandar.Datos); j++ {
			if string(eliminando.Nombre[0]) == mandar.Datos[j].Indice {
				segundo = (primero * len(mandar.Datos)) + j
				break
			}
		}
	}
	tercero = (segundo*5 + eliminando.Calificacion) - 1
	fmt.Println(tercero)
	var nodorecibir *Listas.ListaEnlazada = veclin[tercero].Eliminartienda(eliminando.Nombre)
	var mensaje string
	if nodorecibir == veclin[tercero] {
		mensaje = "Se elimino una tienda"
		veclin[tercero] = nil
		veclin[tercero] = nodorecibir
	} else {
		mensaje = "No se elimino nada"
	}
	fmt.Println("El mensaje que lleva es " + mensaje)
	return mensaje
}

//buscar tienda
func BuscandoTienda() Tiendas {
	var primero, segundo, tercero int
	for j := 0; j < len(mandar.Datos[0].Departamentos); j++ {
		if mandar.Datos[0].Departamentos[j].Nombre == buscar.Departamento {
			var h int = 1
			var ps int
			//Encotrando posicion vector
			if j == 0 {
				h = 0
				ps = 0
			} else if j == 1 {
				h = 0
				ps = 1
			} else {
				h = 1
				ps = j + 1
			}
			primero = ps - h
			break
		}
	}
	for i := 0; i < len(buscar.Nombre); i++ {
		for j := 0; j < len(mandar.Datos); j++ {
			if string(buscar.Nombre[0]) == mandar.Datos[j].Indice {
				segundo = (primero * len(mandar.Datos)) + j
				break
			}
		}
	}
	tercero = (segundo*5 + buscar.Calificacion) - 1
	var nodorecibir *Listas.Nodo = veclin[tercero].PasarNodo(buscar.Nombre)
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
func listan() {
	for i := 0; i < len(veclin); i++ {
		nuevalista := Listas.CrearLista()
		veclin[i] = nuevalista
	}
}

var veclin []*Listas.ListaEnlazada

//Linealizando vector
func Linealizando() {
	//vector
	var primero, segundo, tercero int
	veclin = make([]*Listas.ListaEnlazada, longector())
	listan()
	for i := 0; i < len(mandar.Datos[0].Departamentos); i++ { //Entra Columna
		var h int = 1
		var ps int
		//Encotrando posicion vector
		if i == 0 {
			h = 0
			ps = 0
		} else if i == 1 {
			h = 0
			ps = 1
		} else {
			h = 1
			ps = i + 1
		}
		primero = ps - h
		for j := 0; j < len(mandar.Datos); j++ { //Entra fila
			segundo = (primero * len(mandar.Datos)) + j
			for k := 0; k < len(mandar.Datos[j].Departamentos[i].Tiendas); k++ {
				//Crear Nodo
				nodomandar := Listas.Nodo{mandar.Datos[j].Departamentos[i].Tiendas[k].Nombre, mandar.Datos[j].Departamentos[i].Tiendas[k].Descripcion, mandar.Datos[j].Departamentos[i].Tiendas[k].Contacto, mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion, nil, nil}
				fmt.Println("formo nodo")
				if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 1 {
					tercero = segundo*5 + 0
					fmt.Println(tercero)
					veclin[tercero].InsertarNodo(&nodomandar)
					fmt.Println("Insertonodo")
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 2 {
					tercero = segundo*5 + 1
					veclin[tercero].InsertarNodo(&nodomandar)
					fmt.Println("Insertonodo")
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 3 {
					tercero = segundo*5 + 2
					veclin[tercero].InsertarNodo(&nodomandar)
					fmt.Println("Insertonodo")
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 4 {
					tercero = segundo*5 + 3
					veclin[tercero].InsertarNodo(&nodomandar)
					fmt.Println("Insertonodo")
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 5 {
					tercero = segundo*5 + 4
					veclin[tercero].InsertarNodo(&nodomandar)
					fmt.Println("Insertonodo")
				}
			}
		}
	}
	/*for i := 0; i < len(veclin); i++ {
		fmt.Println(veclin[i])
	}*/
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
	router.HandleFunc("/getArreglo", graficando).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
