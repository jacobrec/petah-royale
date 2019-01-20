import React, { Component } from 'react';

const url = "http://localhost:8049"
const jid = "join-loc"

function makeGame(){
  fetch(`${url}/new`, { method: 'post' })
  .then((r) => r.json())
  .then((r) => window.location.href = `game/${r}`)
  .catch((e) => console.error(e))
}

function joinGame(){
  const loc = document.getElementById(jid).value
  if(loc)
    window.location.href = `game/${loc}`
}

export class Join extends Component {
  constructor(props) {
    super(props)
    window.addEventListener("keypress", (ev) => {
      if(ev.key === "Enter") {
        joinGame()
      }
    })
  }

  render() {
    return (
      <div className="App">
        <input id={jid}/>
        <button onClick={(ev) => joinGame()}>Join Game</button>
      </div>
    )
  }
}

export default class Setup extends Component {
  render() {
    return (
      <div className="App">
        <button onClick={(ev) => makeGame()}>Create Game</button>
        <button onClick={(ev) => window.location.href='join' }>Join Game</button>
      </div>
    )
  }
}
