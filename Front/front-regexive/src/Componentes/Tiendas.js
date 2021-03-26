import React, {useEffect, useState} from 'react'
import MosaicoTiendas from './MosaicoTiendas'
import '../css/Tiendas.css'
const axios=require('axios')
function Tiendas() {
    const [productos, setproductos] = useState([])
    const [loading, setloading] = useState(false)
    useEffect(()=>{
        async function Obtener(){
            if(productos.length===0){
                const data=await axios.get('https://gorest.co.in/public-api/products');
                console.log(data.data.data)
                setproductos(data.data.data)
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
                <MosaicoTiendas productos={productos} />
            </div>
        )
    }
}

export default Tiendas
