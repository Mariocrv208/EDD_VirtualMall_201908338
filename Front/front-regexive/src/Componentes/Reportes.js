import React from 'react'
import '../css/archivos.css'
import 'bootstrap/dist/css/bootstrap.css'
import CuadroC from './CuadroC'
import Buscar from './Buscar'

function Reportes() {
    return (
        <div className="archivos">
            <div>
                <>
                <CuadroC/>
                <Buscar/>
                </>
            </div>
        </div>
    )
}

export default Reportes
