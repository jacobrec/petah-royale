const pallet = {
  wall: "#F8B500",
  player: "#2E94B5",
  enemy: "#D6231E",
  background: "#1F2226",
  shot: "#FFF4E0"
}

export default class UI {
  draw(gfx, world) {
    if(!world.player.alive)
      return gfx.background(pallet.enemy)
    gfx.background(pallet.background)

    // Draw Player
    gfx.rectangle(world.player.x, world.player.y, world.player.size, world.player.size, pallet.player)

    // Draw Walls
    world.walls.forEach((wall) => {
      gfx.rectangle(wall.x, wall.y, wall.width, wall.height, pallet.wall)
    })

    // Draw Other Players
    world.enemies.forEach((person) => {
      gfx.rectangle(person.x, person.y, person.size, person.size, pallet.enemy)
    })

    // Draw Shots
    world.shots.forEach((shot) => {
      gfx.line(shot.x1, shot.y1, shot.x2, shot.y2, pallet.shot)
      if(shot.stamp + world.bulletLife < Date.now()){
        shot.dead = true
      }
    })
    world.shots = world.shots.filter((s) => !s.dead)
  }
}
