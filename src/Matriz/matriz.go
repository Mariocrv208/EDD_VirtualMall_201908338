package Matriz

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type MatrizDis struct {
	Departamento string
	dia          int
	cola         *ColaPedidos
}

type ListaEnlazadaMes struct {
	Cabeza *NodoMes
	cola   *NodoMes
}

type NodoMes struct {
	Mes    int
	matriz *MatrizDis
	Sig    *NodoMes
	Ant    *NodoMes
}

type ArbolAn struct {
	raiz *NodoAnio
}

type NodoAnio struct {
	anio      int
	mes       *ListaEnlazadaMes
	Factor    int
	Izquierdo *NodoAnio
	Derecho   *NodoAnio
}

type NodoPedido struct {
	Fecha        string
	Tienda       string
	Departamento string
	califiacion  int
	productors   []int
	sig          *NodoPedido
}

type ColaPedidos struct {
	Cabeza *NodoPedido
}

//ARBOLANIOS
func NuevoArbolAnio() *ArbolAn {
	return &ArbolAn{nil}
}

func NuevoNodoAnio(anio int, mes *ListaEnlazadaMes) *NodoAnio {
	return &NodoAnio{anio, mes, 0, nil, nil}
}

func rotacioniziAnio(n *NodoAnio, n1 *NodoAnio) *NodoAnio {
	n.Izquierdo = n1.Derecho
	n1.Derecho = n
	if n1.Factor == -1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}

func rotaciondedeAnio(n *NodoAnio, n1 *NodoAnio) *NodoAnio {
	n.Derecho = n1.Izquierdo
	if n1.Factor == 1 {
		n.Factor = 0
		n1.Factor = 0
	} else {
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}

func rotaciondeizAnio(n *NodoAnio, n1 *NodoAnio) *NodoAnio {
	n2 := n1.Izquierdo
	n.Derecho = n2.Izquierdo
	n2.Izquierdo = n
	n1.Izquierdo = n2.Derecho
	n2.Derecho = n1
	if n2.Factor == 1 {
		n.Factor = -1
	} else {
		n.Factor = 0
	}
	if n2.Factor == -1 {
		n1.Factor = 1
	} else {
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func rotacionizdeAnio(n *NodoAnio, n1 *NodoAnio) *NodoAnio {
	n2 := n1.Derecho
	n.Izquierdo = n2.Derecho
	n2.Derecho = n
	n1.Derecho = n2.Izquierdo
	n2.Izquierdo = n1
	if n2.Factor == 1 {
		n1.Factor = -1
	} else {
		n1.Factor = 0
	}
	if n2.Factor == -1 {
		n.Factor = 1
	} else {
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func InsertarAnio(raiz *NodoAnio, anio int, mes *ListaEnlazadaMes, hc *bool) *NodoAnio {
	var n1 *NodoAnio
	if raiz == nil {
		raiz = NuevoNodoAnio(anio, mes)
		*hc = true
	} else if anio < raiz.anio {
		izq := InsertarAnio(raiz.Izquierdo, anio, mes, hc)
		raiz.Izquierdo = izq
		if *hc {
			switch raiz.Factor {
			case 1:
				raiz.Factor = 0
				*hc = false
				break
			case 0:
				raiz.Factor = -1
				break
			case -1:
				n1 = raiz.Izquierdo
				if n1.Factor == -1 {
					raiz = rotacioniziAnio(raiz, n1)
				} else {
					raiz = rotacionizdeAnio(raiz, n1)
				}
				*hc = false
			}
		}
	} else if anio > raiz.anio {
		der := InsertarAnio(raiz.Derecho, anio, mes, hc)
		raiz.Derecho = der
		if *hc {
			switch raiz.Factor {
			case 1:
				n1 = raiz.Derecho
				if n1.Factor == 1 {
					raiz = rotaciondedeAnio(raiz, n1)
				} else {
					raiz = rotaciondeizAnio(raiz, n1)
				}
				*hc = false
				break
			case 0:
				raiz.Factor = 1
				break
			case -1:
				raiz.Factor = 0
				*hc = false
			}
		}
	}
	return raiz
}

func (this *ArbolAn) InsertarArbolAnio(anio int, mes *ListaEnlazadaMes) {
	b := false
	a := &b
	this.raiz = InsertarAnio(this.raiz, anio, mes, a)
}

func imprimirnodosAnio(raiz *NodoAnio) string {
	var linea string = ""
	if raiz.Izquierdo == nil && raiz.Derecho == nil {
		linea = "nodo" + strconv.Itoa(raiz.anio) + "[label =\"" + strconv.Itoa(raiz.anio) + "\"];\n"
	} else {
		linea = "nodo" + strconv.Itoa(raiz.anio) + "[label =\"" + strconv.Itoa(raiz.anio) + "\"];\n"
	}
	if raiz.Izquierdo != nil {
		linea = linea + imprimirnodosAnio(raiz.Izquierdo) + "nodo" + strconv.Itoa(raiz.anio) + "->nodo" + strconv.Itoa(raiz.Izquierdo.anio) + "\n"
	}
	if raiz.Derecho != nil {
		linea = linea + imprimirnodosAnio(raiz.Derecho) + "nodo" + strconv.Itoa(raiz.anio) + "->nodo" + strconv.Itoa(raiz.Derecho.anio) + "\n"
	}
	return linea
}

var contprod int = 1

func GraficarAnio(raiz *ArbolAn) {
	archivo, _ := os.Create("ArbolMamalon" + strconv.Itoa(contprod) + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("rankdir=UD" + "\n")
	_, _ = archivo.WriteString("node[shape=box]" + "\n")
	_, _ = archivo.WriteString("concentrate=true" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	_, _ = archivo.WriteString(imprimirnodosAnio(raiz.raiz))
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./ArbolMamalon"+strconv.Itoa(contprod)+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("ArbolMamalon"+strconv.Itoa(contprod)+".pdf", cmd, os.FileMode(mode))
	contprod++
}

//LISTAMESES

//MATRIZMES

//NODOPEDIDOS
