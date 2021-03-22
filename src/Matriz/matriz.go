package Matriz

import (
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strconv"
)

//Matriz
type NCVertical struct{
	ESTE, NORTE, SUR, OESTE interface{}
	Departamento 	string
}

type NCHorizontal struct{
	ESTE, NORTE, SUR, OESTE interface{}
	dia 	int
}

type MatrizDis struct {
	mes 	int
	CabH 	*NCHorizontal
	CabV 	*NCVertical
}

type NodoMatrizDis struct{
	ESTE, NORTE, SUR, OESTE interface{}
	Departamento string
	fil			 int
	dia          int
	cola         *ColaPedidos
}

//Lista anios
type NodoListaanio struct {
	anioo      	int
	mesesisto   int
	dia 		int
	departamento string
	mess        *ListaEnlazadaMes
	Siganio     *NodoListaanio
	Antanio     *NodoListaanio
}

type ListaEnlazadaanio struct {
	CabezaAnio 	*NodoListaanio
	colaAnio   	*NodoListaanio
}

//Lista mes
type NodoListaMes struct {
	Mes    		int
	Dia 		int
	departamento string
	matriz 		*MatrizDis
	Sigmes    	*NodoListaMes
	Antmes    	*NodoListaMes
}

type ListaEnlazadaMes struct {
	Cabezames 	*NodoListaMes
	colames   	*NodoListaMes
}

//Cola
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
		b := CrearNodoListaMes(nanio.mesesisto, nanio.dia, nanio.departamento, nil)
		nanio.mess.InsertarNodoMes(b)
		inseranio.CabezaAnio = nanio
		inseranio.colaAnio = nanio
		return
	} else {
		auxanio := inseranio.CabezaAnio
		if nanio.anioo < inseranio.CabezaAnio.anioo{
			a = CrearListaMes()
			nanio.mess = a
			b := CrearNodoListaMes(nanio.mesesisto, nanio.dia, nanio.departamento, nil)
			nanio.mess.InsertarNodoMes(b)
			nanio.Siganio = inseranio.CabezaAnio
			inseranio.CabezaAnio.Antanio = nanio
			inseranio.CabezaAnio = nanio
			return
		} 
		for auxanio != nil {
			if nanio.anioo == auxanio.anioo{
				b := CrearNodoListaMes(nanio.mesesisto, nanio.dia, nanio.departamento, nil)
				auxanio.mess.InsertarNodoMes(b)
				return
			}
			if nanio.anioo < auxanio.anioo{
				a = CrearListaMes()
				nanio.mess = a
				b := CrearNodoListaMes(nanio.mesesisto, nanio.dia, nanio.departamento, nil)
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
			b := CrearNodoListaMes(nanio.mesesisto, nanio.dia, nanio.departamento, nil)
			nanio.mess.InsertarNodoMes(b)
			inseranio.colaAnio.Siganio = nanio
			nanio.Antanio = inseranio.colaAnio
			inseranio.colaAnio = nanio
			return
		}
	}
}

func CrearNodoListaAnio(resanio int, resmes int, resdia int, depa string) *NodoListaanio {
	return &NodoListaanio{resanio, resmes, resdia, depa , nil, nil, nil}
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
		a := &MatrizDis{n.Mes, nil, nil}
		b := &NodoMatrizDis{ESTE:nil ,OESTE:nil ,SUR:nil, NORTE:nil, Departamento: n.departamento,fil: 0, dia:n.Dia, cola:nil }
		a.insertar(b)
		n.matriz = a
		inser.Cabezames = n
		inser.colames = n
	} else {
		aux := inser.Cabezames
		if n.Mes < inser.Cabezames.Mes{
			a := &MatrizDis{n.Mes, nil, nil}
			b := &NodoMatrizDis{ESTE:nil ,OESTE:nil ,SUR:nil, NORTE:nil, Departamento: n.departamento,fil: 0, dia:n.Dia, cola:nil }
			a.insertar(b)
			n.matriz = a
			n.Sigmes = inser.Cabezames
			inser.Cabezames.Antmes = n
			inser.Cabezames = n
			return
		}
		for aux != nil {
			if n.Mes == aux.Mes{
				a := &MatrizDis{n.Mes, nil, nil}
				b := &NodoMatrizDis{ESTE:nil ,OESTE:nil ,SUR:nil, NORTE:nil, Departamento: n.departamento,fil: 0, dia:n.Dia, cola:nil }
				a.insertar(b)
				n.matriz = a
				return
			}
			if n.Mes < aux.Mes{
				a := &MatrizDis{n.Mes, nil, nil}
				b := &NodoMatrizDis{ESTE:nil ,OESTE:nil ,SUR:nil, NORTE:nil, Departamento: n.departamento,fil: 0, dia:n.Dia, cola:nil }
				a.insertar(b)
				n.matriz = a
				n.Antmes = aux.Antmes
				n.Sigmes = aux
				aux.Antmes = n
				aux.Antmes.Sigmes = n
				return
			}
			aux = aux.Sigmes
		}
		if aux == nil{
			a := &MatrizDis{n.Mes, nil, nil}
			b := &NodoMatrizDis{ESTE:nil ,OESTE:nil ,SUR:nil, NORTE:nil, Departamento: n.departamento,fil: 0, dia:n.Dia, cola:nil }
			a.insertar(b)
			n.matriz = a
			n.Antmes = inser.colames
			inser.colames.Sigmes = n
			inser.colames = n
		}
	}
}

func CrearNodoListaMes(mes int, dia int, depa string,  matriz *MatrizDis) *NodoListaMes {
	return &NodoListaMes{mes, dia, depa, matriz , nil, nil}
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
		GraficarMatriz(auxgraficar.matriz)
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
func (this *MatrizDis) getCVertical(Departamento string) interface{} {
	if this.CabV==nil{
		return nil
	}
	var aux interface{}=this.CabV
	for aux !=nil{
		if aux.(*NCVertical).Departamento[0]==Departamento[0]{
			return aux
		}
		aux = aux.(*NCVertical).SUR
	}
	return nil
}

func (this *MatrizDis) getHorizontal(dia int) interface{} {
	if this.CabH == nil{
		return nil
	}
	var aux interface{} = this.CabH
	for aux != nil{
		if aux.(*NCHorizontal).dia == dia{
			return aux
		}
		aux = aux.(*NCHorizontal).ESTE
	}
	return nil
}

func (this *MatrizDis) crearNodoVertical(departamento string) *NCVertical{
	if this.CabV == nil{
		nueva :=&NCVertical{
			ESTE:nil,
			OESTE:nil,
			SUR:nil,
			NORTE:nil,
			Departamento: departamento,
		}
		this.CabV = nueva
		return nueva
	}
	var aux interface{}=this.CabV
	if departamento[0]<= aux.(*NCVertical).Departamento[0]{
		nueva :=&NCVertical{
			ESTE:nil,
			OESTE:nil,
			SUR:nil,
			NORTE:nil,
			Departamento: departamento,
		}
		nueva.SUR = this.CabV
		this.CabV.NORTE = nueva
		this.CabV = nueva
		return nueva
	}
	for aux.(*NCVertical).SUR!=nil{
		if departamento[0]>aux.(*NCVertical).Departamento[0] && departamento[0]<aux.(*NCVertical).SUR.(*NCVertical).Departamento[0]{
			nueva := &NCVertical{
				ESTE:nil,
				OESTE:nil,
				SUR:nil,
				NORTE:nil,
				Departamento: departamento,
			}
			tmp := aux.(*NCVertical).SUR
			tmp.(*NCVertical).NORTE=nueva
			nueva.SUR=tmp
			aux.(*NCVertical).SUR=nueva
			nueva.NORTE = aux
			return nueva
		} 
		aux=aux.(*NCVertical).SUR
	}
	nueva := &NCVertical{
		ESTE:nil,
		OESTE:nil,
		SUR:nil,
		NORTE:nil,
		Departamento: departamento,
	}
	aux.(*NCVertical).SUR=nueva
	nueva.NORTE=aux
	return nueva
}

func (this *MatrizDis) crearNodoHorizontal(dia int) *NCHorizontal{
	if this.CabH == nil{
		nueva :=&NCHorizontal{
			ESTE:nil,
			OESTE:nil,
			SUR:nil,
			NORTE:nil,
			dia: dia,
		}
		this.CabH = nueva
		return nueva
	}
	var aux interface{}=this.CabH
	if dia<= aux.(*NCHorizontal).dia{
		nueva :=&NCHorizontal{
			ESTE:nil,
			OESTE:nil,
			SUR:nil,
			NORTE:nil,
			dia: dia,
		}
		nueva.ESTE = this.CabH
		this.CabH.OESTE = nueva
		this.CabH = nueva
		return nueva
	}
	for aux.(*NCHorizontal).ESTE!=nil{
		if dia>aux.(*NCHorizontal).dia && dia<aux.(*NCHorizontal).SUR.(*NCHorizontal).dia{
			nueva := &NCHorizontal{
				ESTE:nil,
				OESTE:nil,
				SUR:nil,
				NORTE:nil,
				dia: dia,
			}
			tmp := aux.(*NCHorizontal).ESTE
			tmp.(*NCHorizontal).OESTE=nueva
			nueva.ESTE=tmp
			aux.(*NCHorizontal).ESTE=nueva
			nueva.OESTE = aux
			return nueva
		} 
		aux=aux.(*NCHorizontal).ESTE
	}
	nueva := &NCHorizontal{
		ESTE:nil,
		OESTE:nil,
		SUR:nil,
		NORTE:nil,
		dia: dia,
	}
	aux.(*NCHorizontal).ESTE=nueva
	nueva.OESTE=aux
	return nueva
}

func (this *MatrizDis) ultimoV(cabecH *NCHorizontal, dia int) interface{} {
	if cabecH.SUR == nil{
		return cabecH
	}
	aux := cabecH.SUR
	if dia <= aux.(*NodoMatrizDis).dia{
		return cabecH
	}
	for aux.(*NodoMatrizDis).SUR!=nil{
		if dia > aux.(*NodoMatrizDis).dia && dia <= aux.(*NodoMatrizDis).SUR.(*NodoMatrizDis).dia{
			return aux
		}
		aux = aux.(*NodoMatrizDis).SUR
	}
	if dia <= aux.(*NodoMatrizDis).dia{
		return aux.(*NodoMatrizDis).NORTE
	}
	return aux
}

func (this *MatrizDis) ultimoH(cabecV *NCVertical, departamento string) interface{} {
	if cabecV.ESTE == nil{
		return cabecV
	}
	aux := cabecV.ESTE
	if departamento[0] <= aux.(*NodoMatrizDis).Departamento[0]{
		return cabecV
	}
	for aux.(*NodoMatrizDis).ESTE!=nil{
		if departamento[0] > aux.(*NodoMatrizDis).Departamento[0] && departamento[0] <= aux.(*NodoMatrizDis).ESTE.(*NodoMatrizDis).Departamento[0]{
			return aux
		}
		aux = aux.(*NodoMatrizDis).ESTE
	}
	if departamento[0] <= aux.(*NodoMatrizDis).Departamento[0]{
		return aux.(*NodoMatrizDis).OESTE
	}
	return aux
}

func (this *MatrizDis) insertar(np *NodoMatrizDis){
	vert := this.getCVertical(np.Departamento)
	horiz := this.getHorizontal(np.dia)
	if vert==nil{
		vert=this.crearNodoVertical(np.Departamento)
	}
	if horiz==nil{
		horiz=this.crearNodoHorizontal(np.dia)
	}
	izquierda := this.ultimoH(vert.(*NCVertical), np.Departamento)
	superior := this.ultimoV(horiz.(*NCHorizontal),np.dia)
	if reflect.TypeOf(izquierda).String()=="*Matriz.superior"{
		if izquierda.(*NodoMatrizDis).ESTE==nil{
			izquierda.(*NodoMatrizDis).ESTE=np
			np.OESTE=izquierda
		}else{
			tmp:= izquierda.(*NodoMatrizDis).ESTE
			izquierda.(*NodoMatrizDis).ESTE=np
			np.OESTE=izquierda
			tmp.(*NodoMatrizDis).OESTE=np
			np.ESTE=tmp
		}
	}else{
		if izquierda.(*NCVertical).ESTE==nil{
			izquierda.(*NCVertical).ESTE=np
			np.OESTE=izquierda
		}else{
			tmp:=izquierda.(*NCVertical).ESTE
			izquierda.(*NCVertical).ESTE=np
			np.OESTE=izquierda
			tmp.(*NodoMatrizDis).OESTE=np
			np.ESTE=tmp
		}
	}
	//Sup
	if reflect.TypeOf(superior).String()=="*Matriz.NodoMatrizDis"{
		if superior.(*NodoMatrizDis).SUR==nil{
			superior.(*NodoMatrizDis).SUR= np
			np.NORTE=superior
		}else{
			tmp:= superior.(*NodoMatrizDis).SUR
			superior.(*NodoMatrizDis).SUR=np
			np.NORTE=superior
			tmp.(*NodoMatrizDis).NORTE=np
			np.SUR=tmp
		}
	}else{
		if superior.(*NCHorizontal).SUR==nil{
			superior.(*NCHorizontal).SUR=np
			np.NORTE=superior
		}else{
			tmp:= superior.(*NCHorizontal).SUR
			superior.(*NCHorizontal).SUR=np
			np.NORTE=superior
			tmp.(*NodoMatrizDis).NORTE=np
			np.SUR=tmp
		}
	}
}

var contmat int = 0

func GraficarMatriz(matriz *MatrizDis){
	var auxgraficarV interface{} = matriz.CabV
	var fila int = 0
	var col int = 0
	archivo, _ := os.Create("Matriz" + strconv.Itoa(contmat) + ".dot")
	_, _ = archivo.WriteString("digraph grafico{" + "\n")
	_, _ = archivo.WriteString("rankdir=LR" + "\n")
	_, _ = archivo.WriteString("node[shape=box]" + "\n")
	_, _ = archivo.WriteString("concentrate=true" + "\n")
	_, _ = archivo.WriteString("compound=true;" + "\n")
	_, _ = archivo.WriteString("nodo00 [label =\"" + strconv.Itoa(matriz.mes)  + "\"];\n")
	for auxgraficarV != nil{
		if fila != 0{
			_, _ = archivo.WriteString("nodo" + strconv.Itoa(fila-1) + "0->nodo" + strconv.Itoa(fila) + "0\n")
		}
		if fila == 0 {
			_, _ = archivo.WriteString("nodo00->nodo10\n")
		}
		fila++
		_, _ = archivo.WriteString("nodo"+strconv.Itoa(fila)+ "0[label =\"" + auxgraficarV.(*NCVertical).Departamento  + "\",color=blue,style=filled];\n")
		tmp1:=auxgraficarV.(*NCVertical).ESTE
		for tmp1!=nil{
			var PEder bool = false
			_, _ = archivo.WriteString("nodo"+strconv.Itoa(fila) + strconv.Itoa(tmp1.(*NodoMatrizDis).dia)+"[shape=polygon,color=red,style=filled ];\n")
			tmp1.(*NodoMatrizDis).fil = fila
			if PEder == true{
				tmp3 := tmp1.(*NodoMatrizDis).OESTE
				_, _ = archivo.WriteString("nodo"+strconv.Itoa(fila)+strconv.Itoa(tmp3.(*NodoMatrizDis).dia)+"->nodo"+strconv.Itoa(fila)+ strconv.Itoa(tmp1.(*NodoMatrizDis).dia) + "\n")
			}
			if PEder == false{
				_, _ = archivo.WriteString("nodo"+strconv.Itoa(fila)+"0->nodo"+strconv.Itoa(fila)+ strconv.Itoa(tmp1.(*NodoMatrizDis).dia) + "\n")
				PEder = true
			}
			tmp2:=tmp1.(*NodoMatrizDis).NORTE
			for reflect.TypeOf(tmp2).String()!="*Matriz.NCHorizontal"{
				tmp4 := tmp2.(NodoMatrizDis).SUR
				_, _ = archivo.WriteString("nodo"+strconv.Itoa(tmp4.(*NodoMatrizDis).fil)+strconv.Itoa(tmp4.(*NodoMatrizDis).dia)+"->nodo"+ strconv.Itoa(tmp2.(*NodoMatrizDis).fil) + strconv.Itoa(tmp2.(*NodoMatrizDis).dia) + "\n")
				tmp2 = tmp2.(*NodoMatrizDis).NORTE
			}
			if reflect.TypeOf(tmp2).String()=="*Matriz.NCHorizontal"{
				tmp5 := tmp2.(*NCHorizontal).SUR
				if col == 0{
					_, _ = archivo.WriteString("nodo00->nodo0"+ strconv.Itoa(tmp2.(*NCHorizontal).dia)+"\n")
				}
				col++
				_, _ = archivo.WriteString("nodo0"+ strconv.Itoa(tmp2.(*NCHorizontal).dia)+"[label =\"" + strconv.Itoa(tmp2.(*NCHorizontal).dia)  + "\",color=blue,style=filled];\n")
				_, _ = archivo.WriteString("nodo"+ strconv.Itoa(tmp5.(*NodoMatrizDis).fil) + strconv.Itoa(tmp5.(*NodoMatrizDis).dia) + "->nodo0"+ strconv.Itoa(tmp2.(*NCHorizontal).dia) + "\n")
			}
			tmp1 = tmp1.(*NodoMatrizDis).ESTE
		}
		auxgraficarV = auxgraficarV.(*NCVertical).SUR
	}
	_, _ = archivo.WriteString("}" + "\n")
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./Matriz"+strconv.Itoa(contmat)+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("Matriz"+strconv.Itoa(contmat)+".pdf", cmd, os.FileMode(mode))
	contmat++
}

//NODOPEDIDOS

