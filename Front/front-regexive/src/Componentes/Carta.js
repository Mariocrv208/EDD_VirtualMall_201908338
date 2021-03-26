import React from 'react'
import '../css/Carta.css'

function Carta(props) {
    return (
        <div className="column carta">
            <div className="ui card">
                <div className="image">
                    <img src={props.imagen} />
                </div>
                <div className="content">
                    <div className="header">{props.nombre}</div>
                    <div className="meta">
                        <a>{props.categoria}</a>
                    </div>
                    <div className="description">{props.descripcion}</div>
                    <div className="ui basic green button center fluid" onClick={()=>{console.log(props.id)}}>Comprar</div>
                </div>
                <div className="extra content">
                    <span className="right floated">Joined in 2013</span>
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default Carta