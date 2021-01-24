import axios from 'axios'

export const getAll = () => {
    const request = axios.get('/api/v1/names')
    return request.then(response => response.data)
}

const nameService = {
    getAll
}

export default nameService
