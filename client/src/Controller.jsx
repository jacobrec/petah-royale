export default class Controller {
  constructor(world, jgraphics){
    this.world = world
    this.jgraphics = jgraphics

    this.setupKeys()
  }

  setupKeys() {
    this.jgraphics.input.onMouseClick((x, y) => { this.world.player.shootTowards(x, y) },
      {x:0, y:0, width:this.world.size.width, height: this.world.size.height})
    this.jgraphics.input.onKeyDown('w', () => { this.world.player.vy = 5 })
    this.jgraphics.input.onKeyDown('s', () => { this.world.player.vy = -5 })
    this.jgraphics.input.onKeyDown('d', () => { this.world.player.vx = 5 })
    this.jgraphics.input.onKeyDown('a', () => { this.world.player.vx = -5 })

    this.jgraphics.input.onKeyUp('w', () => { this.world.player.vy = 0 })
    this.jgraphics.input.onKeyUp('s', () => { this.world.player.vy = 0 })
    this.jgraphics.input.onKeyUp('d', () => { this.world.player.vx = 0 })
    this.jgraphics.input.onKeyUp('a', () => { this.world.player.vx = 0 })
  }
}
