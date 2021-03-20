package Matriz

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type MatrizDis struct {
	Departamento string
	dia          int
	cola         *ColaPedidos
}

type NodoListaanio struct {
	anioo      int
	mesesisto      int
	mess       *ListaEnlazadaMes
	Siganio    *NodoListaanio
	Antanio   *NodoListaanio
}

type ListaEnlazadaanio struct {
	CabezaAnio *NodoListaanio
	colaAnio   *NodoListaanio
}

type NodoListaMes struct {
	Mes    int
	matriz *MatrizDis
	Sigmes    *NodoListaMes
	Antmes    *NodoListaMes
}

type ListaEnlazadaMes struct {
	Cabezames *NodoListaMes
	colames   *NodoListaMes
}

type NodoPedido struct {
	Fecha        string
	Tienda       string
	Departamento string
	califiacion  int
	productors   []int
	sig          *NodoPedido
}

type ColaPedidos struct {
	Cabeza *NodoPedido
}

//LISTAANIO
func CrearListaAnio() *ListaEnlazadaanio {
	return &ListaEnlazadaanio{nil, nil}
}

func (inseranio *ListaEnlazadaanio) InsertarNodoAnio(nanio *NodoListaanio) {
	var a *ListaEnlazadaMes
	if inseranio.CabezaAnio == nil {
		a = CrearListaMes()
		nanio.mess = a
		b := CrearNodoListaMes(nanio.mesesisto, nil)
		nanio.mess.InsertarNodoMes(b)
		inseranio.CabezaAnio = nanio
		inseranio.colaAnio = nanio
		return
	} else {
		auxanio := inseranio.CabezaAnio
		if nanio.anioo < inseranio.CabezaAnio.anioo{
			a = CrearListaMes()
			nanio.mess = a
			b := CrearNodoListaMes(nanio.mesesisto, nil)
			nanio.mess.InsertarNodoMes(b)
			nanio.Siganio = inseranio.CabezaAnio
			inseranio.CabezaAnio.Antanio = nanio
			inseranio.CabezaAnio = nanio
			return
		} 
		for auxanio != nil {
			if nanio.anioo == auxanio.anioo{
				b := CrearNodoListaMes(nanio.mesesisto, nil)
				auxanio.mess.InsertarNodoMes(b)
				return
			}
			if nanio.anioo < auxanio.anioo{
				a = CrearListaMes()
				nanio.mess = a
				b := CrearNodoListaMes(nanio.mesesisto, nil)
				nanio.mess.InsertarNodoMes(b)
				nanio.Antanio = auxanio.Antanio
				nanio.Siganio = auxanio
				auxanio.Antanio = nanio
				auxanio.Antanio.Siganio = nanio
				return
			}
			auxanio = auxanio.Siganio
		}
		if auxanio == nil{
			a = CrearListaMes()
			nanio.mess = a
			b := CrearNodoListaMes(nanio.mesesisto, nil)
			nanio.mess.InsertarNodoMes(b)
			inseranio.colaAnio.Siganio = nanio
			nanio.Antanio = inseranio.colaAnio
			inseranio.colaAnio = nanio
			return
		}
	}
}

func CrearNodoListaAnio(resanio int, resmes int) *NodoListaanio {
	return &NodoListaanio{resanio, resmes, nil , nil, nil}
}

func GraficarListaAnio(lisanio *ListaEnlazadaanio){
	auxgraficar := lisanio.CabezaAnio
	var contlisanio int = 0
	archivo, _ := os.Create("ListaAnio" + strconv.Itoa(contlisanio) + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("rankdir=LR" + "\n")
	_, _ = archivo.WriteString("node[shape=box]" + "\n")
	_, _ = archivo.WriteString("concentrate=true" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	var conta int = 0
	for auxgraficar != nil{
		GraficarListaMes(auxgraficar.mess)
		fmt.Println(auxgraficar.anioo)
		_, _ = archivo.WriteString("nodo" + strconv.Itoa(auxgraficar.anioo) + "[label =\"" + strconv.Itoa(auxgraficar.anioo) + "\"];\n")
		if conta != 0 {
			_, _ = archivo.WriteString("nodo" + strconv.Itoa(auxgraficar.Antanio.anioo) + "->nodo" + strconv.Itoa(auxgraficar.anioo) + "\n")
		}
		auxgraficar = auxgraficar.Siganio
		conta++
	}
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./ListaAnio"+strconv.Itoa(contlisanio)+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("ListaAnio"+strconv.Itoa(contlisanio)+".pdf", cmd, os.FileMode(mode))
	contlisanio++
}

//LISTAMESES

func CrearListaMes() *ListaEnlazadaMes {
	return &ListaEnlazadaMes{nil, nil}
}

func (inser *ListaEnlazadaMes) InsertarNodoMes(n *NodoListaMes) {
	if inser.Cabezames == nil {
		fmt.Println("primero")
		fmt.Println(n.Mes)
		inser.Cabezames = n
		inser.colames = n
	} else {
		aux := inser.Cabezames
		if n.Mes < inser.Cabezames.Mes{
			n.Sigmes = inser.Cabezames
			inser.Cabezames.Antmes = n
			inser.Cabezames = n
			return
		}
		for aux != nil {
			if n.Mes == aux.Mes{

				return
			}
			if n.Mes < aux.Mes{
				n.Antmes = aux.Antmes
				n.Sigmes = aux
				aux.Antmes = n
				aux.Antmes.Sigmes = n
				return
			}
			aux = aux.Sigmes
		}
		if aux == nil{
			n.Antmes = inser.colames
			inser.colames.Sigmes = n
			inser.colames = n
		}
	}
}

func CrearNodoListaMes(mes int, matriz *MatrizDis) *NodoListaMes {
	return &NodoListaMes{mes, matriz , nil, nil}
}

var contlismes int = 0

func GraficarListaMes(lisanio *ListaEnlazadaMes){
	auxgraficar := lisanio.Cabezames
	archivo, _ := os.Create("ListaMes" + strconv.Itoa(contlismes) + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("rankdir=LR" + "\n")
	_, _ = archivo.WriteString("node[shape=box]" + "\n")
	_, _ = archivo.WriteString("concentrate=true" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	var conta int = 0
	for auxgraficar != nil{
		fmt.Println(auxgraficar.Mes)
		_, _ = archivo.WriteString("nodo" + strconv.Itoa(auxgraficar.Mes) + "[label =\"" + getMes(auxgraficar.Mes) + "\"];\n")
		if conta != 0 {
			_, _ = archivo.WriteString("nodo" + strconv.Itoa(auxgraficar.Antmes.Mes) + "->nodo" + strconv.Itoa(auxgraficar.Mes) + "\n")
		}
		auxgraficar = auxgraficar.Sigmes
		conta++
	}
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./ListaMes"+strconv.Itoa(contlismes)+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("ListaMes"+strconv.Itoa(contlismes)+".pdf", cmd, os.FileMode(mode))
	contlismes++
}

func getMes(dato int) string{
	switch dato {
	case 1:
		return "Enero"
	case 2:
		return "Febrero"
	case 3:
		return "Marzo"
	case 4:
		return "Abril"
	case 5:
		return "Mayo"
	case 6:
		return "Junio"
	case 7:
		return "Julio"
	case 8:
		return "Agosto"
	case 9:
		return "Septiembre"
	case 10:
		return "Octubre"
	case 11:
		return "Noviembre"
	case 12:
		return "Diciembre"
	}
	return ""
}


//MATRIZMES

//NODOPEDIDOS
