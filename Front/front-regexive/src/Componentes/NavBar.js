import {React, useState} from 'react'
import { Link } from 'react-router-dom'
import{Menu} from "semantic-ui-react"
import '../css/Nav.css'

const urls = ['Tiendas','Carrito','Reportes','Archivo']
const colores = ['blue','red','green','teal']
const opciones = ['Principal','Carrito','Reportes','Archivos']
function NavBar() {
    const [activo, setactivo]= useState(colores[0])
    return (
        <Menu inverted className = "Nav">
            {colores.map((c,index)=>(
                <Menu.Item as={Link} to={urls[index]}
                    key={c}
                    name={opciones[index]}
                    active={activo=== c}
                    color = {c}
                    onClick={()=>setactivo(c)}
                />
            ))}
        </Menu>
    )
}

export default NavBar
