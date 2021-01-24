const NameRow = ({name}) => { 
    return (
        <tr>
            <td>{name.name}</td>
            <td>{name.amount}</td>
        </tr>
    )
}

export default NameRow