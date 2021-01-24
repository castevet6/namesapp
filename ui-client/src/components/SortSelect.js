import React from 'react'

import Row from 'react-bootstrap/Row'
import Button from 'react-bootstrap/Button'
import Col from 'react-bootstrap/Col'

// render content conditionally based on state variables "sortTerm" and "ascending"
const SortSelect = ({sortTerm, setSortTerm, ascending, setAscending}) => {
    return (
        <>
            <Row>
            {sortTerm === "name"
                ?
                <Col>
                    <span>Sort list by: </span>
                    <Button variant="secondary"
                            disabled
                    >
                        Name
                    </Button>
                    <Button variant="primary"
                            onClick={() => setSortTerm("amount")}
                    >
                        Amount
                    </Button>
                </Col>
                :
                <Col>
                    <span>Sort list by: </span>
                    <Button variant="primary"
                            onClick={() => setSortTerm("name")}
                    >
                        Name
                    </Button>
                    <Button variant="secondary"
                            disabled
                    >
                        Amount
                    </Button>
                </Col>
            }
            {ascending
                ?
                <Col>
                    <span>Ordering: </span>
                    <Button variant="secondary"
                            disabled
                    >
                        Ascending
                    </Button>
                    <Button variant="primary"
                            onClick={() => setAscending(false)}
                    >
                        Descending
                    </Button>
                </Col>
                :
                <Col>
                    <span>Ordering: </span>
                    <Button variant="primary"
                            onClick={() => setAscending(true)}
                    >
                        Ascending
                    </Button>
                    <Button variant="secondary"
                            disabled
                    >
                        Descending
                    </Button>
                </Col>
                
            }
            </Row>
        </>
    )
}

export default SortSelect