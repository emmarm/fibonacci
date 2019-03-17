import React, { Component } from "react";

class App extends Component {
  state = {
    number: 1,
    response: "",
    error: ""
  };
  setNumber = e => {
    const number = e.target.value;

    this.setState(() => ({ number }));
  };
  onSubmit = e => {
    e.preventDefault();
    fetch(`/api/fibonacci/${e.target[0].value}`)
      .then(res => res.text())
      .then(response => {
        this.setState(() => ({ response }));
      })
      .catch(error => {
        this.setState(() => ({ error }));
      });
  };
  render() {
    return (
      <div>
        <form onSubmit={this.onSubmit} method="get" action="">
          <input
            type="number"
            value={this.state.number}
            onChange={this.setNumber}
          />
          <button type="submit">Fibonaccisize!</button>
        </form>
        <p>{this.state.response}</p>
        <p>{this.state.error}</p>
      </div>
    );
  }
}

export default App;
