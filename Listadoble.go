package main

import "fmt"

type Node struct {
	Nombre       string
	Descripcion  string
	contacto     string
	calificacion int
	Sig          *Node
	Ant          *Node
}

type ListaEnlazada struct {
	cabeza   *Node
	longitud int
}

func (i *ListaEnlazada) InsertarNodo(n *Node) {
	nuevo := i.cabeza
	i.cabeza = n
	i.cabeza.Sig = nuevo
	/*i.cabeza.Sig.Ant = nuevo*/
	i.longitud++
}

func (i *ListaEnlazada) Eliminar(dato int, nombre string) {
	if i.longitud == 0 {
		return
	}
	if i.cabeza.calificacion == dato {
		i.cabeza = i.cabeza.Sig
		i.longitud--
		return
	}
	borrar := i.cabeza
	for borrar.Sig.Nombre != nombre {
		if borrar.Sig.Sig == nil {
			return
		}
		borrar = borrar.Sig
	}
	borrar.Sig = borrar.Sig.Sig
	i.longitud--
}

func (i *ListaEnlazada) Buscar(dato int, nombre string) {
	if i.longitud == 0 {
		return
	}
	if i.cabeza.calificacion == dato {
		fmt.Printf("Estudiante encontrado con exito")
		fmt.Printf("\n")
		return
	}
	buscando := i.cabeza
	for buscando.Sig.Nombre != nombre {
		buscando = buscando.Sig
	}
	fmt.Printf("Estudiante encontrado con exito")
	fmt.Printf("\n")
}

func (i ListaEnlazada) ImprimirDatos() {
	imprimiendo := i.cabeza
	for i.longitud != 0 {
		fmt.Printf("%d ", imprimiendo.calificacion)
		fmt.Printf(imprimiendo.Nombre)
		fmt.Printf("\n")
		imprimiendo = imprimiendo.Sig
		i.longitud--
	}
}

func main() {
	miLista := ListaEnlazada{}
	nodo1 := &Node{calificacion: 5, Nombre: "zapates"}
	nodo2 := &Node{calificacion: 5, Nombre: "comida"}
	nodo3 := &Node{calificacion: 5, Nombre: "tacos"}
	miLista.InsertarNodo(nodo1)
	miLista.InsertarNodo(nodo2)
	miLista.InsertarNodo(nodo3)
	miLista.ImprimirDatos()
	miLista.Buscar(5, "zapates")
	miLista.Eliminar(5, "comida")
	miLista.ImprimirDatos()

}
