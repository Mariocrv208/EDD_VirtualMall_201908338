import React, { useEffect, useState } from 'react'
import Tabla from './Table'
import '../css/Carrito.css'


function Carrito() {
    const encabezados = ["id", "Nombre", "Apellido", "De acuerdo","Acciones"]
    const [listado, setlistado] = useState([
        ["Mario", "Rodriguez", true],
        ["Tiwu", "Gonzales", false]
    ])
    useEffect(() => {
        let data = localStorage.getItem('usuarios')
        if (data != null) {
            setlistado(JSON.parse(data))
        }
    }, [])
    return (
        <div className="Carrito">
            <Tabla data={listado}
                encabezados={encabezados}
            />
        </div>

    )
}

export default Carrito
