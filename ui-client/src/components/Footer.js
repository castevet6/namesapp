import React from 'react'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'

const Footer = () => {
    return (
        <Row>
            <Container>
                <Row>Copyright (C) 2021 Antti Manninen</Row>
                <Row><a href="mailto:antti.e.manninen@tuni.fi">antti.e.manninen@tuni.fi</a></Row>
            </Container>
        </Row>
    )
}

export default Footer