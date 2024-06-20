import './App.css';
import Table from "./Table";
import Board from "./Board";

let b;

function App() {
  fetch('http://localhost:8080/init')
      .then(response => response.json())
      .then(data => b = Board.fromJSON(data))
      .catch(error => console.error(error));

  return (
    <div className="App">
      <header className="App-header">
        <Table board={b}/>
      </header>
    </div>
  );
}

export default App;
