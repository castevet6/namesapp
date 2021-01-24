import React from 'react'

import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'

import NamesView from './NamesView'

const MainView = ({names}) => {
    return (
        <>        
            <Container>   
                <Row>
                    <h2 className="app-title">Solita Dev Academy 2021: Names application</h2>
                </Row>
                <NamesView names={names} />
            </Container>
        </>
    )
}

export default MainView