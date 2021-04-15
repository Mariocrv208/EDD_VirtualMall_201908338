import React from 'react'
import { Input } from 'semantic-ui-react'

const InputExampleAction = () => (
  <Input className="ui input" id="PedArc" label="Cargar Pedidos" type='file' onChange={
    (e)=>{
        if (e.target.files[0]!=null){
            let reader = new FileReader()
            reader.readAsText(e.target.files[0], "UTF-8")
            reader.onload=(a)=>{
            console.log(a.target.result)
            //setArchivo(a.target.result)
        }
        }
    }
}/>
)

export default InputExampleAction