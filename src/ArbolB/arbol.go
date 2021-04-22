package ArbolB

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

)

type Arbol struct{
	k int
	Raiz *Nodo
}

type Llavesitas struct{
	Llave string `json:"LLave"`
}

var Llave Llavesitas

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
					if node.Keys[j].dpi < llavecentral.dpi{
						izquierdo.Colocar(indiceizquierdo,node.Keys[j])
						indiceizquierdo++
						node.Colocar(j, nil)
					}else if node.Keys[j].dpi>llavecentral.dpi{
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
					if node.Keys[i].dpi>newkey.dpi{
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
					if node.Keys[i].dpi<llavecentral.dpi{
						izquierdo.Colocar(indiceizquierdo,node.Keys[i])
						indiceizquierdo++
						node.Colocar(i,nil)
					}else if node.Keys[i].dpi>llavecentral.dpi{
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
							if node.Keys[i].dpi < llavecentralraiz.dpi {
								izquierdoraiz.Colocar(indiceizquierdoraiz, node.Keys[i])
								indiceizquierdoraiz++
								node.Colocar(i, nil)
							} else if node.Keys[i].dpi > llavecentralraiz.dpi {
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
			for j:=i-1;j>=0 ;j--{
				if node.Keys[j].dpi > newkey.dpi{
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

func (this *Arbol) Graficar() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"green\"]\n")
	m := make(map[string]*Nodo)
	graficando(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	f, err := os.Create("ArbolB.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./ArbolB.dot").Output()
	mode := int(0772)
	ioutil.WriteFile("ArbolB.pdf", cmd, os.FileMode(mode))
}						

func graficando(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene2 := arr[fmt.Sprint(&(*actual))]
	if contiene2 != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\".7 .3 1.0\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				fmt.Fprintf(cad, "{<f%d>DPI: %d|", j, actual.Keys[i].dpi)
				fmt.Fprintf(cad, "Nombre: %s|", actual.Keys[i].nombre)
				fmt.Fprintf(cad, "Correo: %s|", actual.Keys[i].correo)
				fmt.Fprintf(cad, "Password: %s|", actual.Keys[i].pass)
				fmt.Fprintf(cad, "Cuenta: %s}|", actual.Keys[i].cuenta)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficando(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficando(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}

func (this *Arbol) GraficarArbolCifrado() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"green\"]\n")
	m := make(map[string]*Nodo)
	graficandocifrado(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	f, err := os.Create("ArbolBCi.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./ArbolBCi.dot").Output()
	mode := int(0772)
	ioutil.WriteFile("ArbolBCi.pdf", cmd, os.FileMode(mode))
}

func (this *Arbol) GraficarArbolCifradoSencible() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"green\"]\n")
	m := make(map[string]*Nodo)
	graficandocifradoS(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	f, err := os.Create("ArbolBCiS.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./ArbolBCiS.dot").Output()
	mode := int(0772)
	ioutil.WriteFile("ArbolBCiS.pdf", cmd, os.FileMode(mode))
}

func encrypt(dato string, llave string) string{
	key := []byte("gopostmediumkeyleisock5894621652")
	plaintext := []byte(dato)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("gopostmedium")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	retorno := string(key)
	return retorno
}

func graficandocifrado(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\".7 .3 1.0\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				fmt.Fprintf(cad, "{<f%d>DPI: %s|", j, encrypt(strconv.Itoa(actual.Keys[i].dpi), Llave.Llave))
				fmt.Fprintf(cad, "Nombre: %s|", encrypt(actual.Keys[i].nombre,Llave.Llave))
				fmt.Fprintf(cad, "Correo: %s|", encrypt(actual.Keys[i].correo,Llave.Llave))
				fmt.Fprintf(cad, "Password: %s|", actual.Keys[i].pass)
				fmt.Fprintf(cad, "Cuenta: %s}|", encrypt(actual.Keys[i].cuenta, Llave.Llave))
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficandocifrado(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficandocifrado(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}
func graficandocifradoS(actual *Nodo, cad *strings.Builder, arr map[string]*Nodo, padre *Nodo, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene := arr[fmt.Sprint(&(*actual))]
	if contiene != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\".7 .3 1.0\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				fmt.Fprintf(cad, "{<f%d>DPI: %s|", j, encrypt(strconv.Itoa(actual.Keys[i].dpi), Llave.Llave))
				fmt.Fprintf(cad, "Nombre: %s|", actual.Keys[i].nombre)
				fmt.Fprintf(cad, "Correo: %s|", encrypt(actual.Keys[i].correo, Llave.Llave))
				fmt.Fprintf(cad, "Password: %s|", actual.Keys[i].pass)
				fmt.Fprintf(cad, "Cuenta: %s}|", actual.Keys[i].cuenta)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		graficandocifradoS(actual.Keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficandocifradoS(actual.Keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}


var Usuario Key

func Buscar(actual *Nodo, dato Key, padre *Nodo, pos int) Key {

	if actual == nil {
		return Usuario
	}
	j := 0
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			return Usuario
		} else {
			if enlace {

				enlace = false
				i--
				j++

			} else {

				if dato.dpi == actual.Keys[i].dpi || dato.correo == actual.Keys[i].cuenta {
					if dato.pass == actual.Keys[i].pass {
						Usuario.dpi = actual.Keys[i].dpi
						Usuario.nombre = actual.Keys[i].nombre
						Usuario.correo = actual.Keys[i].correo
						Usuario.pass = actual.Keys[i].pass
						Usuario.cuenta = actual.Keys[i].cuenta
						fmt.Println(Usuario.nombre)
						fmt.Println(Usuario.correo)
						fmt.Println(Usuario.pass)
						fmt.Println(Usuario.cuenta)
					}
				}
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.Keys[i+1] == nil {

						j++
						break
					}
				}
			}
		}
	}
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.Keys[i] == nil {
			break
		}
		Buscar(actual.Keys[i].Izquierdo, dato, actual, ji)
		ji++
		ji++
		Buscar(actual.Keys[i].Derecho, dato, actual, ji)
		ji++
		ji--
	}
	if padre != nil {

	}

	return Usuario

}




