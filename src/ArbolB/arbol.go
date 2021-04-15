package ArbolB

type Arbol struct{
	k int
	Raiz *Nodo
}

func NewArbol(nivel int)*Arbol  {
	a:=Arbol{nivel, nil}
	nodoraiz:= NewNodo(nivel)
	a.Raiz=nodoraiz
	return &a
}

func (this *Arbol) Insertar(newkey *Key)  {
	if this.Raiz.Keys[0]==nil{
		this.Raiz.Colocar(0,newkey)
	}else if this.Raiz.Keys[0].Izquierdo == nil{
		lugarinsertado:= -1
		node :=this.Raiz
		lugarinsertado=this.colocarNodo(node, newkey)
		if lugarinsertado !=-1{
			if lugarinsertado == node.Max-1{
				middle := node.Max/2
				llavecentral:=node.Keys[middle]
				derecho:=NewNodo(this.k)
				izquierdo:=NewNodo(this.k)
				indiceizquierdo := 0
				indicederecho :=0
				for j:=0;j<node.Max;j++{
					if node.Keys[j].Value < llavecentral.Value{
						izquierdo.Colocar(indiceizquierdo,node.Keys[j])
						indiceizquierdo++
						node.Colocar(j, nil)
					}else if node.Keys[j].Value>llavecentral.Value{
						derecho.Colocar(indicederecho,node.Keys[j])
						indicederecho++
						node.Colocar(j,nil)
					}
				}
				node.Colocar(middle, nil)
				this.Raiz=node
				this.Raiz.Colocar(0,llavecentral)
				izquierdo.NodoPadre=this.Raiz
				derecho.NodoPadre=this.Raiz
				llavecentral.Izquierdo=izquierdo
				llavecentral.Derecho=derecho
			}
		}
	}else if this.Raiz.Keys[0].Izquierdo!=nil{
		node:=this.Raiz
		for node.Keys[0].Izquierdo!=nil{
			loop:=0
			for i:=0;i<node.Max;i,loop=i+1, loop+1{
				if node.Keys[i]!=nil{
					if node.Keys[i].Value>newkey.Value{
						node=node.Keys[i].Izquierdo
						break
					}
				}else{
					node=node.Keys[i-1].Derecho
					break
				}
			}
			if loop == node.Max{
				node=node.Keys[loop-1].Derecho
			}
		}
		indiceColocado :=this.colocarNodo(node, newkey)
		if indiceColocado==node.Max-1{
			for node.NodoPadre!=nil{
				indicemedio:=node.Max/2
				llavecentral:=node.Keys[indicemedio]
				izquierdo:=NewNodo(this.k)
				derecho:=NewNodo(this.k)
				indiceizquierdo, indicederecho :=0,0
				for i:=0;i<node.Max;i++{
					if node.Keys[i].Value<llavecentral.Value{
						izquierdo.Colocar(indiceizquierdo,node.Keys[i])
						indiceizquierdo++
						node.Colocar(i,nil)
					}else if node.Keys[i].Value>llavecentral.Value{
						derecho.Colocar(indicederecho,node.Keys[i])
						indicederecho++
						node.Colocar(i,nil)
					}
				}
				node.Colocar(indicemedio,nil)
				llavecentral.Izquierdo=izquierdo
				llavecentral.Derecho=derecho
				node=node.NodoPadre
				izquierdo.NodoPadre=node
				derecho.NodoPadre=node
				for i:=0; i<izquierdo.Max;i++{
					if izquierdo.Keys[i]!=nil{
						if izquierdo.Keys[i].Izquierdo!=nil{
							izquierdo.Keys[i].Izquierdo.NodoPadre=izquierdo
						}
						if izquierdo.Keys[i].Derecho!=nil{
							izquierdo.Keys[i].Derecho.NodoPadre=izquierdo
						}
					}
				}
				for i:=0; i<derecho.Max;i++{
					if derecho.Keys[i]!=nil{
						if derecho.Keys[i].Izquierdo!=nil{
							derecho.Keys[i].Izquierdo.NodoPadre=derecho
						}
						if derecho.Keys[i].Derecho!=nil{
							derecho.Keys[i].Derecho.NodoPadre=derecho
						}
					}
				}
				lugarcolocado:=this.colocarNodo(node,llavecentral)
				if lugarcolocado==node.Max-1 {
					if node.NodoPadre == nil {
						indicecentralraiz := node.Max / 2
						llavecentralraiz := node.Keys[indicecentralraiz]
						izquierdoraiz := NewNodo(this.k)
						derechoraiz := NewNodo(this.k)
						indiceizquierdoraiz, indicederechoraiz := 0, 0
						for i := 0; i < node.Max; i++ {
							if node.Keys[i].Value < llavecentralraiz.Value {
								izquierdoraiz.Colocar(indiceizquierdoraiz, node.Keys[i])
								indiceizquierdoraiz++
								node.Colocar(i, nil)
							} else if node.Keys[i].Value > llavecentralraiz.Value {
								derechoraiz.Colocar(indicederechoraiz, node.Keys[i])
								indicederechoraiz++
								node.Colocar(i, nil)
							}
						}
						node.Colocar(indicecentralraiz, nil)
						node.Colocar(0, llavecentralraiz)
						for i := 0; i < this.k; i++ {
							if izquierdoraiz.Keys[i] != nil {
								izquierdoraiz.Keys[i].Izquierdo.NodoPadre = izquierdoraiz
								izquierdoraiz.Keys[i].Derecho.NodoPadre = izquierdoraiz
							}
						}
						for i := 0; i < this.k; i++ {
							if derechoraiz.Keys[i] != nil {
								derechoraiz.Keys[i].Izquierdo.NodoPadre = derechoraiz
								derechoraiz.Keys[i].Derecho.NodoPadre = derechoraiz
							}
						}
						llavecentralraiz.Izquierdo = izquierdoraiz
						llavecentralraiz.Derecho = derechoraiz
						izquierdoraiz.NodoPadre = node
						derechoraiz.NodoPadre = node
						this.Raiz = node
					}
					continue
				}else{
					break
				}
			}
		}
	}
}


func (this *Arbol) colocarNodo(node *Nodo,newkey *Key)int  {
	index:= -1
	for i:=0; i<node.Max;i++ {
		if node.Keys[i]==nil{
			placed:=false
			for j:=i-1;j>=0 ;j++{
				if node.Keys[j].Value > newkey.Value{
					node.Colocar(j+1,node.Keys[j])
				}else{
					node.Colocar(j+1,newkey)
					node.Keys[j].Derecho=newkey.Izquierdo
					if j+2<this.k&&node.Keys[j+2]!=nil{
						node.Keys[j+2].Izquierdo=newkey.Derecho
					}
					placed=true
					break
				}
			}
			if placed == false{
				node.Colocar(0,newkey)
				node.Keys[1].Izquierdo=newkey.Derecho
			}
			index = i
			break
		}
	}
	return index
}