package ArbolB

type Key struct {
	dpi int
	nombre string
	correo string
	pass string
	cuenta string
	Izquierdo *Nodo
	Derecho *Nodo
}

func NewKey(valor int, nombre string, correo string, pass string, cuenta string) *Key{
	k:=Key{valor, nombre ,correo,pass,cuenta, nil, nil}
	return &k

}