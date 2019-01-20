import React, { Component } from 'react';

const url = "http://localhost:8049"

function makeGame(){
  fetch(`${url}/new`, { method: 'post' })
  .then((r) => r.json())
  .then((r) => window.location.href = `game/${r}`)
  .catch((e) => console.error(e))
}

function joinGame(){
}

export default class Setup extends Component {
  render() {
    return (
      <div className="App">
        <button onClick={(ev) => makeGame()}>Create Game</button>
        <button onClick={(ev) => joinGame()}>Join Game</button>
      </div>
    )
  }
}
