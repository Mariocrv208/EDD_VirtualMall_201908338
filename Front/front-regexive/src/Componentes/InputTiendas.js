import React, {useState} from 'react'
import { Input } from 'semantic-ui-react'
import Tiendas from './Tiendas'

function CrearArchivo() {
    const [Datos, setDatos] = useState([])

    const TiendasCarga = () => (
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

      
        return (
            <div className="InputTiendas">
                <br></br>
                <Tiendas mandarArchivo={Datos} />
            </div>
        )
}

export default CrearArchivo

