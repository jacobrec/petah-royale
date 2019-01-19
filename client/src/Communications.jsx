export default class Coms {
  constructor(url){
    const ws = new WebSocket("ws://" + url)
    ws.onopen = (event) => this.setup(ws)
    ws.onmessage = (event) => this.messageRecived(event.data)

    this.ws = ws
  }

  setup() {
  }

  messageRecived(msg) {

  }

  sendMessage() {
  }
}

