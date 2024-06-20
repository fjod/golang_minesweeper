import React, {useState} from 'react';
import Board from "./Board";

const Table = (board) => {
    const [_, setRenderCount] = useState(board);

    const renderCell = (row, col) => (
        <td key={`${row}-${col}`}>
            <button onClick={() => handleCellClick(row, col)}>
                Cell (board.cells[row][col])
            </button>
        </td>
    );

    const handleCellClick = (row, col) => {
        fetch(`http://localhost:8080/step/${row}/${col}`)
            .then(response => response.json())
            .then(data => setRenderCount(Board.fromJSON(data)))
            .catch(error => console.error(error));
        console.log(`Button (${row}, ${col}) clicked`);
    };

    const renderRow = (row) => (
        <tr key={row}>
            {Array.from({ length: board.width }, (_, col) => renderCell(row, col))}
        </tr>
    );

    return (
        <table>
            <tbody>{Array.from({ length: board.height }, (_, row) => renderRow(row))}</tbody>
        </table>
    );
};

export default Table;