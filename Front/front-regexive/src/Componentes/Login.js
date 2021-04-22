import { React, useState } from 'react'
import '../css/CreateUser.css'

function CreateUser() {
    const [dpi, setdpi] = useState("")
    const [password, setpassword] = useState("")

    const enviar = () => {
        var json = [
            dpi,
            password
        ]
        var datos = localStorage.getItem("usuarios")
        if (datos == null || datos === undefined) {
            localStorage.setItem("usuarios", JSON.stringify([json]))
        } else {
            datos = JSON.parse(datos)
            datos.push(json)
            console.log(datos)
            localStorage.setItem("usuarios", JSON.stringify(datos))
        }
        alert(JSON.stringify(json))
    }
    return (
        <div className="UserList">
            <br></br>
            <div className="ui segment container formulario form">

                <div className="field">
                    <label>DPI</label>
                    <input type="text" name="DPI" placeholder="dpi..." onChange={e => setdpi(e.target.value)} />
                </div>
                <div className="field">
                    <label>Contraseña</label>
                    <input type="text" name="Contra" placeholder="asdf1234" onChange={e => setpassword(e.target.value)} />
                </div>
                <button className="ui button" onClick={enviar} >Submit</button>
            </div>
        </div>
    )
}

export default CreateUser