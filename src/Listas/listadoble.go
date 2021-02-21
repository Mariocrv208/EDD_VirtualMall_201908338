package Listas

import (
	"fmt"
	"strconv"
)

type Nodo struct {
	Nombre       string
	Descripcion  string
	Contacto     string
	Calificacion int
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

func (inser *ListaEnlazada) InsertarNodo(n *Nodo) {
	if inser.cabeza == nil {
		inser.cabeza = n
		inser.cola = n
	} else {
		inser.cola.Sig = n
		n.Ant = inser.cola
		inser.cola = n
	}
}

func (i *Nodo) To_string() string {
	return "Nombre: " + i.Nombre + "Descripcion: " + i.Descripcion + "Contacto: " + i.Contacto + "Calificacion: " + strconv.Itoa(i.Calificacion)
}

//para imprimir
func (i *ListaEnlazada) To_string() string {
	var resultante string
	aux := i.cabeza
	for aux != nil {
		resultante += aux.To_string()
		aux = aux.Sig
	}
	return resultante
}

func (i *ListaEnlazada) Imprimir() {
	fmt.Println("Lista es: ")
	fmt.Println(i.To_string())
}

//para buscar e imprimir un nodo
func (i *ListaEnlazada) PasarNodo(BuscarNombre string) *Nodo {
	fmt.Println("entro: ")
	fmt.Print(BuscarNombre)
	aux := i.cabeza
	for aux != nil {
		if aux.Nombre == BuscarNombre {
			fmt.Print(aux.Nombre)
			break
		}
		aux = aux.Sig
	}
	return aux
}

//para buscarID

func (i *ListaEnlazada) PasarNodoID() int {
	var contador int = 0
	aux := i.cabeza
	for aux != nil {
		aux = aux.Sig
		contador++
	}
	return contador
}

func (i *ListaEnlazada) RecorrerID(ultNombre string) *Nodo {
	aux := i.cabeza
	if ultNombre == "" {
		ultNombre = i.cabeza.Nombre
		return aux
	}
	for aux != nil {
		if aux.Nombre == ultNombre {
			aux = aux.Sig
			ultNombre = aux.Nombre
			return aux
		}
	}
	return aux
}

var mandarmensajito string

//Eliminar
func (i *ListaEnlazada) Eliminartienda(nombrenodo string) *ListaEnlazada {
	mandarmensajito = ""
	var mandarMensaje string = ""
	fmt.Println("Entro a eliminar tienda")
	aux := i.cabeza
	if nombrenodo == i.cabeza.Nombre {
		i.cabeza = aux.Sig
		aux.Sig.Ant = nil
		mandarMensaje = "Se elimino en primer nodo"
		mandarmensajito = mandarMensaje
		fmt.Println("Se elimino en primer nodo")
		return i
	}
	for aux != nil {
		if nombrenodo == aux.Nombre {
			aux.Sig.Ant = aux.Ant
			aux.Ant.Sig = aux.Sig
			mandarMensaje = "Se elimino en nodo intermedio"
			mandarmensajito = mandarMensaje
			fmt.Println("Se elimino en nodo intermedio")
			return i
		}
		if nombrenodo == i.cola.Nombre {
			i.cola = aux.Ant
			aux.Ant.Sig = nil
			mandarMensaje = "Se elimino en nodo final"
			mandarmensajito = mandarMensaje
			fmt.Println("Se elimino en nodo final")
			return i
		}
		aux = aux.Sig
	}
	mandarMensaje = "No se encontro tienda"
	mandarmensajito = mandarMensaje
	fmt.Println("No encontro nodo")
	return i
}
