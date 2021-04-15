package ArbolB

type Nodo struct {
	Max int
	NodoPadre *Nodo
	Keys []*Key
}

func NewNodo(max int) *Nodo{
	keys:=make([]*Key, max)
	n := Nodo{max, nil, keys}
	return &n
}

func (this *Nodo) Colocar(i int, llave *Key) {
	this.Keys[i] = llave
}