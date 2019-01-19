export default class UI {
  draw(gfx, world) {
    gfx.background("rgb(49,49,49)")

    // Draw Player
    gfx.ellipse(world.player.x, world.player.y, world.player.size, world.player.size, "#0099CC")

    // Draw Walls
    world.walls.forEach((wall) => {
      gfx.rectangle(wall.x, wall.y, wall.width, wall.height, "#FF00FF")
    })
  }
}
