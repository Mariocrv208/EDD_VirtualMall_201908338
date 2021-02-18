package Listas

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

func NuevoNodo(nombre string, descripcion string, contacto string, calificacion int) *Nodo {
	return &Nodo{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: calificacion}
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
