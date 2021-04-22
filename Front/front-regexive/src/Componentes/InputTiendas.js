import React, {useState} from 'react'
import { Input } from 'semantic-ui-react'

function CrearArchivo() {
    const [Datos, setDatos] = useState([])
  
    return (
        <Input className="ui input" id="TArc" label="Cargar Tiendas" type='file' onChange={
            (e)=>{
                if (e.target.files[0]!=null){
                    let reader = new FileReader()
                    reader.readAsText(e.target.files[0], "UTF-8")
                    reader.onload=(a)=>{
                    setDatos(a.target.result)
                    
                }
                }
                console.log(Datos)
            }
        }/>
    )
}

export default CrearArchivo

