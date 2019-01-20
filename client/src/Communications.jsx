export default class Coms {
  constructor(url, world){
    const ws = new WebSocket("ws://" + url + "/wsgame")
    ws.onopen = (event) => this.setup(ws)
    ws.onmessage = (event) => this.messageRecived(JSON.parse(event.data), world)

    this.ws = ws

  }

  setup() {
  }

  messageRecived(msg, world) {
    switch(msg.action) {
      case "initial":
        return initWorld(world, msg.data, this)
      case "draw":
        return doDraw(world, msg.data)
      default:
        console.error("Unknown Message Type")
        console.error(msg)
    }

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


// message handlers
function initWorld(world, data, coms){
  console.log(data)
  console.log(world)
  world.player.id = data.id
  world.player.x = data.x
  world.player.y = data.y
  world.player.coms = coms

  world.size.height = data.world.height
  world.size.width = data.world.width

  data.world.walls.forEach((wall) => {
    world.walls.push(wall)
  })

  data.world.players.forEach((p) => {
    p.size = p.radius * 2
    if(p.id !== data.id)
      world.enemies.push(p)
  })
  console.log(world)
}
function doDraw(world, data){

  if(data.id === world.player.id)
    return

  for(let p in world.enemies) {
    if(world.enemies[p].id === data.id){
      world.enemies[p].x = data.x
      world.enemies[p].y = data.y
    }
  }



}
