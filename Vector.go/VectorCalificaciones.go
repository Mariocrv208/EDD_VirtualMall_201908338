package Vector

import (
	"./Lista"
	"./Server"
)

var veclin = make([]Lista.ListaEnlazada, Server.longit)

//Encontrar Posicion en vector
func Posicionamiento() {

}

func InsertarNodo() {
	var primero, segundo, tercero int
	for i := 0; i < len(Server.Raiz.Datos); i++ {
		var h int
		//Encotrando posicion vector
		if i == 0 {
			h = 0
		} else {
			h = i - 1
		}
		primero = i + h

		for j := 0; j < len(Server.Raiz.Datos[i].Departamentos); j++ {
			//Encotrando posicion vector
			segundo = (primero * len(Server.mandar.Datos)) + j
			for k := 0; k < len(Server.Raiz.Datos[i].Departamentos[j].Tiendas); k++ {
				//Crear Nodo
				Lista.NuevoNodo{Nombre: Server.Raiz.Datos[i].Departamentos[j].Tiendas[k].Nombre, Descripcion: Server.Raiz.Datos[i].Departamentos[j].Tiendas[k].Descripcion, Contacto: Server.Raiz.Datos[i].Departamentos[j].Tiendas[k].Contacto, Calificacion: Server.Raiz.Datos[i].Departamentos[j].Tiendas[k].Calificacion}

				//Posicionar Nodo
				var calif int
				calif = Server.Raiz.Datos[i].Departamentos[j].Tiendas[k].Calificacion
				tercero = segundo*5 + calif
				veclin = append(veclin[tercero], Lista.InsertarNodo)
				for m := 0; m < len(veclin); m++ {
					println(veclin[m])
				}
			}
		}
	}
}
