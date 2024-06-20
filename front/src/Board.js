
class Board {
    constructor(width, height) {
        this.width = width;
        this.height = height;
        this.cells = [];
    }

    static fromJSON(json) {
        const { width, height, cells } = json;
        return new Board(width, height, cells);
    }
}

const CellState = Object.freeze({
    Unknown: 0,
    Opened_no_mines_nearby: 1,
    Mine: 2,
    Opened_1_mine_nearby: 3,
    Opened_2_mine_nearby: 4,
    Opened_3_mine_nearby: 5,
    Opened_4_mine_nearby: 6,
    Opened_5_mine_nearby: 7,
    Opened_6_mine_nearby: 8,
    Opened_7_mine_nearby: 9,
    Opened_8_mine_nearby: 10,
    Opened_9_mine_nearby: 11,
    Flagged: 12,
    Selected_cell: 13,
});


export default Board;