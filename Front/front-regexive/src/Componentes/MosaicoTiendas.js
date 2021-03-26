import React from 'react'
import Carta from './Carta'
import '../css/MosaicoTiendas.css'

function MosaicoTiendas(props) {
    return (
        <div className="ui segment mosaico container">
            <div className="ui four column link cards row">
                {props.productos.map((c, index) => (
                    <Carta nombre={c.name}
                        categoria={c.categories[0].name}
                        descripcion={c.description.substring(0,50)}
                        imagen={c.image}
                        precio={c.price}
                        id={c.id}
                        key={c.id}
                    />
                ))}
            </div>
        </div>
    )
}

export default MosaicoTiendas
