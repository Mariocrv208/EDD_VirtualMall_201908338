import React from 'react' 
import { Header, Image } from 'semantic-ui-react'
import '../css/Header.css'

const HeaderMandar = () => (
  <Header as='h1'>
    <Image circular src='./HeaderIcon.png'/>
    <Header.Content>EDD REGEXIVE 201908338</Header.Content>
  </Header>
)

export default HeaderMandar