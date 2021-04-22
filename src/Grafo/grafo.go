package Grafo

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type Vertice struct{
	Nombre 	string
	id		int
	Inicio 	string
	Final 	string
	enlaces []*Enlace
	Adyacentes *list.List
}

type Enlace struct{
	inicio 		*Vertice
	destino 	*Vertice
	distancia 	int
	enlaces 	*list.List
}

type ListaAdyacencia struct {
	Lista *list.List
}

func NewEnlace(inicio *Vertice,distancia int,destino *Vertice) *Enlace{
	a:=list.New()
	return &Enlace{inicio, destino, distancia, a}
}

func NewVertice(nom string, id int, inicio string, final string)* Vertice{
	a:=list.New()
	return &Vertice{nom, id ,inicio, final, nil,a}
}

func NewListaAdyacencia()*ListaAdyacencia{
	a:=list.New()
	return &ListaAdyacencia{a}
}

func (this *ListaAdyacencia)getVertice(n string)*Vertice{
	for e:=this.Lista.Front();e!=nil;e=e.Next(){
		a:=e.Value
		b:=a.(*Vertice)
		if b.Nombre==n{
			return b
		}
	}
	return nil
}

func (this *ListaAdyacencia) Insertar(n string, id int, inicio string, final string){
	if this.getVertice(n)==nil{
		i:=NewVertice(n, id, inicio, final)
		this.Lista.PushBack(i)
	}else{
		fmt.Println("Ya esta creado we")
	}
}

func (this *ListaAdyacencia)Enlazar(n1 string, n2 string, distancia int){
	var origen *Vertice
	var destino *Vertice
	origen=this.getVertice(n1)
	destino=this.getVertice(n2)
	if origen==nil ||destino==nil{
		fmt.Println("No se encontro vertice")
	}
	origen.Adyacentes.PushFront(destino)
	destino.Adyacentes.PushFront(origen)
	c:= NewEnlace(origen, distancia, destino)
	d:= NewEnlace(destino,distancia,origen)
	origen.enlaces = append(origen.enlaces, c)
	destino.enlaces = append(destino.enlaces,d)

}

func contiene(buscando *list.List, elemento *Vertice) bool{
	for e:=buscando.Front();e!=nil;e=e.Next(){
		if (e.Value).(*Vertice)==elemento{
			return true
		}
	}
	return false
}



func (this *ListaAdyacencia)GraficaGrafo(){
	aux:=list.New()
	archivo, _ := os.Create("Grafo.dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("layout=circo" + "\n")
	_, _ = archivo.WriteString("node[shape=circle]" + "\n")
	for e:=this.Lista.Front();e!=nil;e=e.Next(){
		tmp:=(e.Value).(*Vertice)
		if contiene(aux, tmp)==false{
			aux.PushBack(tmp)
			if tmp.Final != ""{
				_, _ = archivo.WriteString("nodo" + strconv.Itoa(tmp.id) + "[label =\"" + tmp.Nombre + "\" style =\"filled\" color =\"blue\"];\n")
			}
			if tmp.Inicio != ""{
				_, _ = archivo.WriteString("nodo" + strconv.Itoa(tmp.id) + "[label =\"" + tmp.Nombre + "\" style =\"filled\" color =\"red\"];\n")
			}
			if tmp.Inicio == "" && tmp.Final == ""{
				_, _ = archivo.WriteString("nodo" + strconv.Itoa(tmp.id) + "[label =\"" + tmp.Nombre + "\"];\n")
			}
		}
		for temporal := tmp.Adyacentes.Front();temporal!=nil;temporal=temporal.Next(){
			verticetemporal:=(temporal.Value).(*Vertice)
			if contiene(aux, verticetemporal)==false{
				aux.PushBack(verticetemporal)
				if verticetemporal.Final != ""{
					_, _ = archivo.WriteString("nodo" + strconv.Itoa(verticetemporal.id) + "[label =\"" + verticetemporal.Nombre + "\" style =\"filled\" color =\"blue\"];\n")
				}
				if verticetemporal.Inicio != ""{
					_, _ = archivo.WriteString("nodo" + strconv.Itoa(verticetemporal.id) + "[label =\"" + verticetemporal.Nombre + "\" style =\"filled\" color =\"red\"];\n")
				}
				if verticetemporal.Inicio == "" && verticetemporal.Final == ""{
					_, _ = archivo.WriteString("nodo" + strconv.Itoa(verticetemporal.id) + "[label =\"" + verticetemporal.Nombre + "\"];\n")
				}
			}
		}
		for h:=0; h<len(tmp.enlaces);h++{
			_, _ = archivo.WriteString("nodo" + strconv.Itoa(tmp.enlaces[h].inicio.id) + "->nodo" + strconv.Itoa(tmp.enlaces[h].destino.id)+ "[label =\"" + strconv.Itoa(tmp.enlaces[h].distancia) + "\"];\n")
		}
	}
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./Grafo.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("Grafo.pdf", cmd, os.FileMode(mode))
}

func (this *ListaAdyacencia) BFS(){
	aux := list.New()
	for e:=this.Lista.Front();e!=nil;e=e.Next(){
		tmp:=(e.Value).(*Vertice)
		if contiene(aux, tmp) == false{
			aux.PushBack(tmp)
		}
		for j:=tmp.Adyacentes.Front();j!=nil;j=j.Next(){
			temporalj:=(j.Value).(*Vertice)
			if contiene(aux, temporalj)==false{
				aux.PushBack(temporalj)
			}
		}
	}
	for e:=aux.Front();e!=nil;e=e.Next(){
		fmt.Println(e.Value)
	}
}

