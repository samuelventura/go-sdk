import './App.css';

function clickHandler() {
  console.log("click")
  fetch('/ping')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.log("error:", error))
}

function App() {
  return (
    <div className="App">
      <button onClick={clickHandler}>Ping</button>
    </div>
  );
}

export default App;
