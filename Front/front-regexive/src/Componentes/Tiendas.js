import React, {useEffect, useState} from 'react'
import MosaicoTiendas from './MosaicoTiendas'
import '../css/Tiendas.css'

const axios=require('axios').default
function Tiendas(props) {
    const [tiendas, settiendas] = useState([])
    //const [productos, setproductos] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(()=>{
        async function Obtener(){
            if(tiendas.length===0){
                console.log(props.mandarArchivo)
                const data=await axios.post('http://localhost:3000/cargartienda', props.mandarArchivo);
                //console.log(data.data.data)
                settiendas(data.data.Datos)
                setloading(true)
            }
        }
        Obtener()
    });
    if (loading === false) {
        return (
            <div className="ui segment carga">
                <div className="ui active dimmer">
                    <div className="ui text loader">Loading</div>
                </div>
                <p />
            </div>
        )
    } else {
        return (
            <div className="Tiendas">
                <br></br>
                <MosaicoTiendas tiendas={tiendas} />
                
            </div>
        )
    }
}

export default Tiendas
