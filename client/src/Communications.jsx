export default class Coms {
  constructor(url){
    const ws = new WebSocket("ws://" + url + "/wsgame")
    ws.onopen = (event) => this.setup(ws)
    ws.onmessage = (event) => this.messageRecived(event.data)

    this.ws = ws
  }

  setup() {
  }

  messageRecived(msg) {
    console.log(msg)
  }

  sendShot(x, y, angle, weapon){
    this.sendMessage({
      action: "attack",
      data: { angle, x, y, weapon }
    })

  }
  sendMove(x, y){
    this.sendMessage({
      action: "move",
      data: { x, y }
    })
  }

  sendMessage(msg) {
    this.ws.send(JSON.stringify(msg))
  }
}

