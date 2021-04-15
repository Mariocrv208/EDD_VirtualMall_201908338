import React from 'react'
import { Input } from 'semantic-ui-react'


const InputExampleAction = () => (
  <Input className="ui input" id="ProdArc" label="Cargar Productos" type='file' onChange={
    (e)=>{
        if (e.target.files[0]!=null){
            let reader = new FileReader()
            reader.readAsText(e.target.files[0], "UTF-8")
            reader.onload=(a)=>{
        }
        }
    }
}/>
)



export default InputExampleAction