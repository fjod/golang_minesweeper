import React, {useEffect, useState} from 'react';
import Board from "./Board";

const Table = (board) => {
    useEffect(() => {
        // This code will run only once, on the initial render
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8080/init');
                const jsonData = await response.json();
                setRenderCount(Board.fromJSON(jsonData));
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []); // Empty dependency array ensures this effect runs only once

    const [b, setRenderCount] = useState(board);

    const renderCell = (row, col) => (
        <td key={`${row}-${col}`}>
            <button onClick={(event) => handleCellClick(row, col ,event)}>
                {b.cells[row][col]}
            </button>
        </td>
    );

    const handleCellClick = (row, col, event) => {
        console.log(event.button)
        switch (event.button) {

            case 0:
                fetch(`http://localhost:8080/step/${row}/${col}/1`)
                    .then(response => response.json())
                    .then(data => setRenderCount(Board.fromJSON(data)))
                    .catch(error => console.error(error));
                break;
            case 1:
                fetch(`http://localhost:8080/step/${row}/${col}/2`)
                    .then(response => response.json())
                    .then(data => setRenderCount(Board.fromJSON(data)))
                    .catch(error => console.error(error));
                break;
            case 2:
                fetch(`http://localhost:8080/step/${row}/${col}/2`)
                    .then(response => response.json())
                    .then(data => setRenderCount(Board.fromJSON(data)))
                    .catch(error => console.error(error));
                break;
            default:
                console.log('Other mouse button clicked');
        }
        console.log(`Button (${row}, ${col}) clicked`);
    };

    const renderRow = (row) => (
        <tr key={row}>
            {Array.from({ length: b.width }, (_, col) => renderCell(row, col))}
        </tr>
    );

    return (
        <table>
            <tbody>{Array.from({ length: b.height }, (_, row) => renderRow(row))}</tbody>
        </table>
    );
};

export default Table;