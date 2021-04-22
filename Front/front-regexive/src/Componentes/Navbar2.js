import {React, useState} from 'react'
import { Link } from 'react-router-dom'
import{Menu} from "semantic-ui-react"
import '../css/Nav.css'

const urls = ['Login','Registro']
const colores = ['blue','red']
const opciones = ['Login','Registro']
function NavBar2() {
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

export default NavBar2
