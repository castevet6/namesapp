import React, { useState, useEffect } from 'react'
import Container from 'react-bootstrap/Container'
import Row from 'react-bootstrap/Row'

import "bootstrap/dist/css/bootstrap.css";

import NamesView from './components/NamesView'
import Footer from './components/Footer'
import SortSelect from './components/SortSelect'

// services imports
import nameService from './services/names'

// util
import sortBy from './util/util'

const App = () => {
    // APP STATES;
    // list of names fetched from data source
    const [ names, setNames ] = useState([])
    // names list sort term: name or amount ?
    const [ sortTerm, setSortTerm ] = useState("name")
    // names list ording: ascending (true) or descending (false)
    const [ ascending, setAscending ] = useState(true)

    // fetch names list from data source
    useEffect(() => {
        nameService.getAll()
            .then(namearray => {
                // use custom sort function from util
                const sortedArray = sortBy (sortTerm) (ascending) (namearray)
                setNames(sortedArray)
            })
    }, [sortTerm, ascending]) // useEffect dependencies

    return (
        <>        
            <Container>   
                <Row>
                    <h2 className="app-title">Solita Dev Academy 2021: Names application</h2>
                </Row>
                <SortSelect sortTerm={sortTerm}
                        setSortTerm={setSortTerm}
                        ascending={ascending}
                        setAscending={setAscending}
                />
                <NamesView names={names} sortTerm={sortTerm} ascending={ascending} />
                <Footer />
            </Container>
        </>
    )
}

export default App