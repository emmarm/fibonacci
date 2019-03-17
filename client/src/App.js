import React, { Component } from "react";

class App extends Component {
  state = {
    number: 1,
    message: "",
    error: ""
  };
  setNumber = e => {
    const number = e.target.value;

    this.setState(() => ({ number }));
  };
  onSubmit = e => {
    e.preventDefault();
    fetch(`/api/fibonacci/${e.target[0].value}`, {
      method: "get",
      dataType: "json",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json"
      }
    })
      .then(res => res.json())
      .then(response => {
        console.log(response);
        const { message, error } = response;
        this.setState(() => ({ message, error }));
      })
      .catch(err => {
        this.setState(() => ({ err }));
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
        <p>{this.state.message}</p>
        <p>{this.state.error}</p>
      </div>
    );
  }
}

export default App;
