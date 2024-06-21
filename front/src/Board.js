
class Board {
    constructor(width, height, cells) {
        this.width = width;
        this.height = height;
        this.cells = cells;
    }

    static fromJSON(json) {
        console.log(json);
        const { width, height, cells } = json;
        return new Board(width, height, cells);
    }
}

export default Board;