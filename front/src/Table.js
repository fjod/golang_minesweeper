import React, {useEffect, useState} from 'react';
import Board from "./Board";

const Table = (board) => {

    function getUnicodeRune(value) {
        // Check if the input value is a valid integer
        if (!Number.isInteger(value)) {
            throw new Error('Input must be an integer');
        }

        if (value == 1){
            return "᫅";
        }
        if (value == 2){
            return "🟥";
        }
        if (value == 3){
            return "1";
        }
        if (value == 4){
            return "2";
        }
        if (value == 5){
            return "3";
        }
        if (value == 6){
            return "4";
        }
        if (value == 7){
            return "5";
        }
        if (value == 8){
            return "6";
        }
        if (value == 9){
            return "7";
        }
        if (value == 10){
            return "8";
        }
        if (value == 11){
            return "9";
        }
        if (value == 12){
            return "🚩";
        }
    }

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
            <button onMouseDown={(event) => handleCellClick(row, col ,event)}>
                {getUnicodeRune(b.cells[row][col])}
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