	package main

	import (
		"encoding/json"
		"fmt"
		"io/ioutil"
		"log"
		"net/http"
		"strconv"
		"strings"

		"./ArbolAVL"
		"./Listas"
		"./Matriz"

		"github.com/gorilla/mux"
	)

//variables
var mandar Raiz
var Mandarprod RaizProd
var Mandarpedidos RaizPedido
var resultadoBusca Tiendas
var eliminando mandarEliminar
var resultadoID Departamentos
var longit int
var buscar mandarBuscar

//Struct Pedidos
type RaizPedido struct {
	Pedidos []Pedidos `json:"Pedidos"`
}

type Pedidos struct {
	Fecha        string    `json:"Fecha"`
	Tienda       string    `json:"Tienda"`
	Departamento string    `json:"Departamento"`
	Calificacion int       `json:"Calificacion"`
	Codigos      []Codigos `json:"Productos"`
}

type Codigos struct {
	Codigo int `json:"Codigo"`
}

//Struct Productos
type RaizProd struct {
	Invetarios []Invetarios `json:"Invetarios"`
}

type Invetarios struct {
	Tienda       string      `json:"Tienda"`
	Departamento string      `json:"Departamento"`
	Calificacion int       `json:"Calificacion"`
	Productos    []Productos `json:"Productos"`
}

type Productos struct {
	NombreProd      string  `json:"Nombre"`
	CodigoProd      int     `json:"Codigo"`
	DescripcionProd string  `json:"Descripcion"`
	PrecioProd      float32 `json:"Precio"`
	CantidadProd    int     `json:"Cantidad"`
	ImagenProd      string  `json:"Imagen"`
}

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
	Logo         string `json:Logo`
}

//funciones
//Server
func inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Si jala we uwu")
}

//AgregarPedidos
func agregarpedi(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No inserto we :c")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &Mandarpedidos)
	json.NewEncoder(w).Encode(Mandarpedidos)
	imprimirListaAnio()
}

//AgregarProductos
func agregarprod(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "No inserto we :c")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.Unmarshal(reqBody, &Mandarprod)
	json.NewEncoder(w).Encode(Mandarprod)
	imprimirProductos()
}

//AgregarTiendas
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
	var mandarJson = guardar(0, 0)
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
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusFound)
	aux := veclin[b].Cabeza
	var paso []Tiendas
	for aux != nil {
		var nt Tiendas = Tiendas{Nombre: aux.Nombre, Descripcion: aux.Descripcion, Contacto: aux.Contacto, Calificacion: aux.Calificacion}
		paso = append(paso, nt)
		aux = aux.Sig
	}
	if veclin[b] == nil {
		json.NewEncoder(w).Encode(mensaje)
	} else {
		json.NewEncoder(w).Encode(paso)
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
	json.NewEncoder(w).Encode(mens)
}

//Guardar
func guardar(in int, fi int) Raiz {
	var reg Raiz
	Indic := make([]string, len(mandar.Datos))
	for i := 0; i < len(mandar.Datos); i++ {
		Indic[i] = mandar.Datos[i].Indice
	}
	Depart := make([]string, len(mandar.Datos[0].Departamentos))
	for i := 0; i < len(mandar.Datos[0].Departamentos); i++ {
		Depart[i] = mandar.Datos[0].Departamentos[i].Nombre
	}
	reg.Datos = make([]Datos, len(mandar.Datos))
	for i := 0; i < len(mandar.Datos); i++ {
		reg.Datos[i].Indice = Indic[i]
		reg.Datos[i].Departamentos = make([]Departamentos, len(mandar.Datos[0].Departamentos))
		for j := 0; j < len(mandar.Datos[0].Departamentos); j++ {
			reg.Datos[i].Departamentos[j].Nombre = Depart[j]
			if in+5 <= len(veclin) {
				fi = fi + 5
			} else if in+4 <= len(veclin) {
				fi = fi + 4
			} else if in+3 <= len(veclin) {
				fi = fi + 3
			} else if in+2 <= len(veclin) {
				fi = fi + 2
			} else if in+1 <= len(veclin) {
				fi = fi + 1
			}
			reg.Datos[i].Departamentos[j].Tiendas = obtenerT(in, fi)
			in = fi
		}
	}
	return reg
}

func obtenerT(i int, f int) []Tiendas {
	var t []Tiendas
	var aux Tiendas
	for i < f {
		a := veclin[i].Getcabeza()
		for a != nil {
			aux.Nombre = a.Nombre
			aux.Descripcion = a.Descripcion
			aux.Contacto = a.Contacto
			aux.Calificacion = a.Calificacion
			t = append(t, aux)
			a = a.Sig
		}
		i++
	}
	return t
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
	for len(veclin) > Listas.Pasomayor {
		Listas.GraficarVeclin(veclin)
		Grafico()
	}
		return "Grafico listo"
}

//Exportar
func exportarVector() *Raiz {
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
				}
				if recibirlong != 0 {
					var buscanombre string = ""
					for j := 0; j < recibirlong; j++ {
						var nodorecibir *Listas.Nodo = veclin[k].RecorrerID(buscanombre)
						buscanombre = nodorecibir.Nombre
						mandarTienda = Tiendas{Nombre: nodorecibir.Nombre, Descripcion: nodorecibir.Descripcion, Contacto: nodorecibir.Contacto, Calificacion: nodorecibir.Calificacion}
						multiTiendas = append(multiTiendas, mandarTienda)
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
	for j := 0; j < len(mandar.Datos); j++ {
		if string(eliminando.Nombre[0]) == mandar.Datos[j].Indice {
			segundo = (primero * len(mandar.Datos)) + j
			break
		}
	}
	tercero = (segundo*5 + eliminando.Calificacion) - 1
	var nodorecibir *Listas.ListaEnlazada
	var mensaje string
	if veclin[tercero] != nil {
		nodorecibir = veclin[tercero].Eliminartienda(eliminando.Nombre)
		veclin[tercero] = nil
		veclin[tercero] = nodorecibir
		mensaje = "Se elimino tienda"
	} else {
		mensaje = "No se elimino nada tienda"
	}
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
				nodomandar := Listas.Nodo{mandar.Datos[j].Departamentos[i].Tiendas[k].Nombre, mandar.Datos[j].Departamentos[i].Tiendas[k].Descripcion, mandar.Datos[j].Departamentos[i].Tiendas[k].Contacto, mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion, mandar.Datos[j].Departamentos[i].Tiendas[k].Logo, nil, nil}
				if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 1 {
					tercero = segundo*5 + 0
					veclin[tercero].InsertarNodo(&nodomandar)
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 2 {
					tercero = segundo*5 + 1
					veclin[tercero].InsertarNodo(&nodomandar)
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 3 {
					tercero = segundo*5 + 2
					veclin[tercero].InsertarNodo(&nodomandar)
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 4 {
					tercero = segundo*5 + 3
					veclin[tercero].InsertarNodo(&nodomandar)
				} else if mandar.Datos[j].Departamentos[i].Tiendas[k].Calificacion == 5 {
					tercero = segundo*5 + 4
					veclin[tercero].InsertarNodo(&nodomandar)
				}
			}
		}
	}
}

//FuncionCrearArbol
func imprimirProductos() {
	for j := 0; j < len(Mandarprod.Invetarios); j++ {
		a := ArbolAVL.NuevoArbol()
		for i := 0; i < len(Mandarprod.Invetarios[j].Productos); i++ {
			a.InsertarArbol(Mandarprod.Invetarios[j].Productos[i].NombreProd, Mandarprod.Invetarios[j].Productos[i].CodigoProd, Mandarprod.Invetarios[j].Productos[i].DescripcionProd, Mandarprod.Invetarios[j].Productos[i].PrecioProd, Mandarprod.Invetarios[j].Productos[i].CantidadProd, Mandarprod.Invetarios[j].Productos[i].ImagenProd)
		}
		ArbolAVL.Graficar(a)
	}
}

//FuncionCrearArbolAnio
func imprimirListaAnio() {
	c:= Matriz.CrearListaAnio()
	for i := 0; i < len(Mandarpedidos.Pedidos); i++ {
		var aniopaso []string
		var paso int
		var paso2 int
		var paso3 int
		var mandarcodigos []int
		aniopaso = strings.Split(Mandarpedidos.Pedidos[i].Fecha, "-")
		paso, _ = strconv.Atoi(aniopaso[2])
		paso2, _ = strconv.Atoi(aniopaso[1])
		paso3, _ = strconv.Atoi(aniopaso[0])
		for j := 0; j < len(Mandarpedidos.Pedidos[i].Codigos); j++ {
			mandarcodigos = append(mandarcodigos, Mandarpedidos.Pedidos[i].Codigos[j].Codigo)
		}
		nodanio := Matriz.CrearNodoListaAnio(paso, paso2, paso3, Mandarpedidos.Pedidos[i].Departamento, Mandarpedidos.Pedidos[i].Tienda,Mandarpedidos.Pedidos[i].Calificacion, mandarcodigos)
		c.InsertarNodoAnio(nodanio)
	}
	Matriz.GraficarListaAnio(c)
}

//main
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", inicial).Methods("GET")
	router.HandleFunc("/cargarproductos", agregarprod).Methods("POST")
	router.HandleFunc("/cargarpedido", agregarpedi).Methods("POST")
	router.HandleFunc("/cargartienda", agregar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", BuscarTienda).Methods("POST")
	router.HandleFunc("/id/{id}", buscarID).Methods("GET")
	router.HandleFunc("/guardar", exportarJson).Methods("GET")
	router.HandleFunc("/Eliminar", EliminarTienda).Methods("DELETE")
	router.HandleFunc("/getArreglo", graficando).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))

}
