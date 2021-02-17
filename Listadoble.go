package Lista

import (
	"fmt"
)

//variables
var colum int
var fila int
var longit int

type Nodo struct {
	Nombre       string
	Descripcion  string
	contacto     string
	calificacion int
	Sig          *Nodo
	Ant          *Nodo
}

type ListaEnlazada struct {
	cabeza *Nodo
	cola   *Nodo
}

func CrearLista() *ListaEnlazada {
	return &ListaEnlazada{nil, nil}
}

func CrearVector() {
	var recivlong int
	recivlong = 5
	fmt.Print(recivlong)
}

func LongitudEncontrar() {

	longit = 5 * len(Raiz.Datos)
	s := make([]string, longit)
}

func (i *ListaEnlazada) InsertarNodo(n *Nodo) {
	if i.cabeza == nil {
		i.cabeza = n
		i.cola = n
	} else {
		i.cola.Sig = n
	}

	nuevo := i.cabeza
	i.cabeza = n
	i.cabeza.Sig = nuevo
	/*i.cabeza.Sig.Ant = nuevo*/
}
