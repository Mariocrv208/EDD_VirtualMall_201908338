import React from 'react'
import {BrowserRouter as Router,Route} from 'react-router-dom'
import archivos from './Componentes/archivos'
import CreateUser from './Componentes/CreateUser'
import Reportes from './Componentes/Reportes'
import Carrito from './Componentes/Carrito'
import NavBar from './Componentes/NavBar'
import Header from './Componentes/Header'
import Tiendas from './Componentes/Tiendas'

function App() {
  return (
    <>
    <Header/>
    <Router>
      <NavBar/>
      <Route path="/Tiendas" component={Tiendas}/> 
      <Route path="/Productos" component={CreateUser}/>
      <Route path="/Carrito" component={Carrito}/>
      <Route path="/Reportes" component={Reportes}/>
      <Route path="/Archivo" component={archivos}/>
    </Router>
    </>
  )
}

export default App

