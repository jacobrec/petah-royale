const pallet = {
  wall: "#F8B500",
  player: "#2E94B5",
  enemy: "#D6231E",
  background: "#1F2226",
  shot: "#FFF4E0"
}

export default class UI {
  draw(gfx, world) {
    gfx.background(pallet.background)

    // Draw Player
    gfx.ellipse(world.player.x, world.player.y, world.player.size*1.2, world.player.size * 1.2, pallet.player)

    // Draw Walls
    world.walls.forEach((wall) => {
      gfx.rectangle(wall.x, wall.y, wall.width, wall.height, pallet.wall)
    })

    // Draw Other Players
    world.enemies.forEach((person) => {
      gfx.ellipse(person.x, person.y, person.size, person.size, "#FFFFFF")//pallet.enemy)
    })
  }
}
