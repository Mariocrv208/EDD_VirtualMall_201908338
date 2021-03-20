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
	Logo         string
	Sig          *Nodo
	Ant          *Nodo
}

type ListaEnlazada struct {
	Cabeza *Nodo
	cola   *Nodo
}

func CrearLista() *ListaEnlazada {
	return &ListaEnlazada{nil, nil}
}

func (inser *ListaEnlazada) InsertarNodo(n *Nodo) {
	if inser.Cabeza == nil {
		inser.Cabeza = n
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
	aux := i.Cabeza
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
	aux := i.Cabeza
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
	aux := i.Cabeza
	for aux != nil {
		aux = aux.Sig
		contador++
	}
	return contador
}

func (i *ListaEnlazada) Getcabeza() *Nodo {
	return i.Cabeza
}

func (i *ListaEnlazada) RecorrerID(ultNombre string) *Nodo {
	aux := i.Cabeza
	if ultNombre == "" {
		ultNombre = i.Cabeza.Nombre
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

//Eliminar
func (i *ListaEnlazada) Eliminartienda(nombrenodo string) *ListaEnlazada {
	fmt.Println("Entro a eliminar tienda")
	aux := i.Cabeza
	if nombrenodo == i.Cabeza.Nombre {
		i.Cabeza = aux.Sig
		aux.Sig.Ant = nil
		fmt.Println("Se elimino en primer nodo")
		return i
	}
	for aux != nil {
		if nombrenodo == i.cola.Nombre {
			i.cola = aux.Ant
			aux.Ant.Sig = nil
			fmt.Println("Se elimino en nodo final")
			return i
		}
		if nombrenodo == aux.Nombre {
			aux.Sig.Ant = aux.Ant
			aux.Ant.Sig = aux.Sig
			fmt.Println("Se elimino en nodo intermedio")
			return i
		}
		aux = aux.Sig
	}
	fmt.Println("No encontro nodo")
	return i
}
