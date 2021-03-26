import React from 'react'
import { useHistory } from "react-router-dom";
function Fila(props) {
    console.log(props.data)
    const history = useHistory();
    const mandar = () => {
        history.push(`/view/${props.index}`)
    }
    return (
        <tr key={props.index}>
            <td>{props.index}</td>
            {props.data.map((dato) => (
                <td>{dato.toString()}</td>
            ))}
            <td className="centered">
                <button className="ui inverted primary button" onClick={mandar}>Visualizar</button>
            </td>
        </tr>
    )
}

export default Fila