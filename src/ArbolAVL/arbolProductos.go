package ArbolAVL

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type ArbolProductos struct {
	raiz *NodoProduc
}

type NodoProduc struct {
	Nombre      string
	Codigo      int
	Descripcion string
	Precio      float32
	cantidad    int
	imagen      string
	Factor      int
	Izquierdo   *NodoProduc
	Derecho     *NodoProduc
}

func NuevoArbol() *ArbolProductos {
	return &ArbolProductos{nil}
}

func NuevoNodo(nombre string, codigo int, descripcion string, precio float32, cantidad int, imagen string) *NodoProduc {
	return &NodoProduc{nombre, codigo, descripcion, precio, cantidad, imagen, 0, nil, nil}
}

func rotacionizi(n *NodoProduc, n1 *NodoProduc) *NodoProduc {
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

func rotaciondede(n *NodoProduc, n1 *NodoProduc) *NodoProduc {
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

func rotaciondeiz(n *NodoProduc, n1 *NodoProduc) *NodoProduc {
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

func rotacionizde(n *NodoProduc, n1 *NodoProduc) *NodoProduc {
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

func InsertarProd(raiz *NodoProduc, nombre string, codigo int, descripcion string, precio float32, cantidad int, imagen string, hc *bool) *NodoProduc {
	var n1 *NodoProduc
	if raiz == nil {
		raiz = NuevoNodo(nombre, codigo, descripcion, precio, cantidad, imagen)
		*hc = true
	} else if codigo < raiz.Codigo {
		izq := InsertarProd(raiz.Izquierdo, nombre, codigo, descripcion, precio, cantidad, imagen, hc)
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
					raiz = rotacionizi(raiz, n1)
				} else {
					raiz = rotacionizde(raiz, n1)
				}
				*hc = false
			}
		}
	} else if codigo > raiz.Codigo {
		der := InsertarProd(raiz.Derecho, nombre, codigo, descripcion, precio, cantidad, imagen, hc)
		raiz.Derecho = der
		if *hc {
			switch raiz.Factor {
			case 1:
				n1 = raiz.Derecho
				if n1.Factor == 1 {
					raiz = rotaciondede(raiz, n1)
				} else {
					raiz = rotaciondeiz(raiz, n1)
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

func (this *ArbolProductos) InsertarArbol(nombre string, codigo int, descripcion string, precio float32, cantidad int, imagen string) {
	b := false
	a := &b
	this.raiz = InsertarProd(this.raiz, nombre, codigo, descripcion, precio, cantidad, imagen, a)
}

func imprimirnodos(raiz *NodoProduc) string {
	var linea string = ""
	if raiz.Izquierdo == nil && raiz.Derecho == nil {
		linea = "nodo" + strconv.Itoa(raiz.Codigo) + "[label =\"" + strconv.Itoa(raiz.Codigo) + " | " + raiz.Nombre + " | " + strconv.Itoa(int(raiz.Precio)) + " | " + strconv.Itoa(raiz.cantidad) + "\"];\n"
	} else {
		linea = "nodo" + strconv.Itoa(raiz.Codigo) + "[label =\"" + strconv.Itoa(raiz.Codigo) + " | " + raiz.Nombre + " | " + strconv.Itoa(int(raiz.Precio)) + " | " + strconv.Itoa(raiz.cantidad) + "\"];\n"
	}
	if raiz.Izquierdo != nil {
		linea = linea + imprimirnodos(raiz.Izquierdo) + "nodo" + strconv.Itoa(raiz.Codigo) + "->nodo" + strconv.Itoa(raiz.Izquierdo.Codigo) + "\n"
	}
	if raiz.Derecho != nil {
		linea = linea + imprimirnodos(raiz.Derecho) + "nodo" + strconv.Itoa(raiz.Codigo) + "->nodo" + strconv.Itoa(raiz.Derecho.Codigo) + "\n"
	}
	return linea
}

var contpedido int = 1

func Graficar(raiz *ArbolProductos) {
	archivo, _ := os.Create("Arbol" + strconv.Itoa(contpedido) + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("rankdir=UD" + "\n")
	_, _ = archivo.WriteString("node[shape=box]" + "\n")
	_, _ = archivo.WriteString("concentrate=true" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	_, _ = archivo.WriteString(imprimirnodos(raiz.raiz))
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./Arbol"+strconv.Itoa(contpedido)+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("Arbol"+strconv.Itoa(contpedido)+".pdf", cmd, os.FileMode(mode))
	contpedido++
}
