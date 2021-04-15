package ArbolB

type Key struct {
	Value int
	Izquierdo *Nodo
	Derecho *Nodo
}

func NewKey(valor int) *Key{
	k:=Key{valor, nil, nil}
	return &k

}