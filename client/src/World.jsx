export default class World {
  constructor(width, height) {
    this.player = new Player(40, 30)
    this.player.world = this
    this.enemies = []
    this.walls = []
    this.shots = []

    this.size = { width, height }

    /*
    this.enemies.push(new Moveable(10,10))
    this.enemies.push(new Moveable(20,10))
    this.enemies.push(new Moveable(10,50))
    this.enemies.push(new Immoveable(0,0,80,1))
    this.enemies.push(new Immoveable(0,0,1,60))
    this.enemies.push(new Immoveable(0,59,80,1))
    this.enemies.push(new Immoveable(79,0,1,60))
    */
  }

  update(delta) {
    this.player.update(delta, this.walls)
  }
}

class Immoveable {
  constructor(x, y, width, height) {
    this.x = x
    this.y = y
    this.width = width
    this.height = height
  }
}

class Moveable {
  constructor(x, y) {
    this.x = x
    this.y = y

    this.vx = 0
    this.vy = 0

    this.size = 1
  }

  update(delta, immovable) {
    const x = this.x + this.vx * delta
    const y = this.y + this.vy * delta

    const cr = this.size/2
    var didHit = false
    immovable.forEach((wall) => {
      if(circleSquareCollides(x+cr, y+cr, cr, wall.x, wall.y, wall.width, wall.height)){
        didHit = true
      }
    })
    if(!didHit){
      this.x = x
      this.y = y
    }
  }
}

class Player extends Moveable {
  constructor(x, y) {
    super(x, y)
  }

  shootTowards(x, y){
    const angle = Math.atan2( y - this.y, x - this.x )
    //this.coms.sendShot(this.x, this.y, angle, 0)
    this.world.shots.push({X1: x, Y1: y, X2: this.x, Y2:this.y})
  }

  update(delta, immovable) {
    super.update(delta, immovable)
    this.coms && this.coms.sendMove(this.x, this.y)
  }

}





// COLLISOIN DETECTION STUFF
function circleSquareCollides(cx, cy, cr, x, y, w, h){
  return (
    pointRectangleCollides(cx, cy, x, y, w, h) ||
    circleLineCollides(cx, cy, cr, x, y, x+w, y) ||
    circleLineCollides(cx, cy, cr, x+w,y,x+w,y+h) ||
    circleLineCollides(cx, cy, cr, x+w,y+h,x,y+h) ||
    circleLineCollides(cx, cy, cr, x,y+h,x,y)
  )
}
function pointRectangleCollides(cx, cy, x, y, w, h){
  return cx > x && cx < x + w && cy > y && cy < y + h
}
function circleLineCollides(cx, cy, cr, px1, py1, px2, py2){
  return inteceptCircleLineSeg({center: {x:cx,y:cy}, radius: cr}, {p1: {x: px1, y: py1}, p2: {x: px2, y: py2}})
}
function inteceptCircleLineSeg(circle, line){
  var b, c, d, v1, v2, u1, u2;
  v1 = {};
  v2 = {};
  v1.x = line.p2.x - line.p1.x;
  v1.y = line.p2.y - line.p1.y;
  v2.x = line.p1.x - circle.center.x;
  v2.y = line.p1.y - circle.center.y;
  b = (v1.x * v2.x + v1.y * v2.y);
  c = 2 * (v1.x * v1.x + v1.y * v1.y);
  b *= -2;
  d = Math.sqrt(b * b - 2 * c * (v2.x * v2.x + v2.y * v2.y - circle.radius * circle.radius));
  if (isNaN(d)) {
    return false
  }
  u1 = (b - d) / c;  // these represent the unit distance of point one and two on the line
  u2 = (b + d) / c;
  if(u1 <= 1 && u1 >= 0){  // add point if on the line segment
    return true;
  }
  if(u2 <= 1 && u2 >= 0){  // second add point if on the line segment
    return true;
  }
  return false;
}
