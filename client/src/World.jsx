export default class World {
  constructor() {
    this.user = {}
    this.enemies = []
    this.walls = []

    this.size = { width: 80, height: 60 }
  }

  update(delta) {
    this.user = new Movable(0, 0)
  }
}

class Movable {
  constructor(x, y) {
    this.x = x
    this.y = y

    this.vx = 0
    this.vy = 0
  }

  update(delta) {
    this.x += this.vx * delta
    this.y += this.vy * delta
  }
}
