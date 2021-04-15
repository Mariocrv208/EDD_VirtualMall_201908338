import React from 'react'
import InputTiendas from './InputTiendas'
import InputProductos from './InputProductos'
import InputPedidos from './InputPedidos'
import btnTienda from './btnTienda'
import '../css/archivos.css'
import 'bootstrap/dist/css/bootstrap.css'

function archivos() {
    return (
        <>
        <div className="archivos">
            <label><h1>Tiendas</h1></label>
            <InputTiendas/>
            <btnTienda/>
        </div>
        <div className="archivos">
            <label><h1>Productos</h1></label>
            <InputProductos/>
        </div>
        <div className="archivos">
            <label><h1>Pedidos</h1></label>
            <InputPedidos/>
        </div>
        </>
    )

    
}

export default archivos
