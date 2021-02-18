package Lista

//variables

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

func NuevoNodo(nombre, descripcion, contacto string, calificacion int) *Nodo {
	return &Node{Nombre: nombre, Descripcion: descripcion, Contacto: contacto, Calificacion: calificacion}
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
