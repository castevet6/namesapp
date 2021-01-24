import React from 'react'
import Table from 'react-bootstrap/Table'

import NameRow from './NameRow'

const NamesView = ({names, sortTerm, ascending}) => {

    const showNames = () => {
        if (names) {
            return names.map((name, idx) => <NameRow key={idx} name={name} />)
        }
        return <p>No names in the database</p>
    }

    return (
        <>
            <Table size="sm" striped hover>
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Amount</th>
                    </tr>
                </thead>
                <tbody>
                    {showNames()}
                </tbody>
            </Table>
        </>
    )
}

export default NamesView