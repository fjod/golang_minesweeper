
class Game {
    constructor(board, minesLeft, steps, gameOver) {
        this.board = board;
        this.minesLeft = minesLeft;
        this.steps = steps;
        this.gameOver = gameOver;
    }

    static fromJSON(json) {
        console.log(json);
        const { board, minesLeft, steps, gameOver} = json;
        return new Game(board, minesLeft, steps, gameOver);
    }
}

export default Game;