import React, { Component } from 'react';
import JGraphics from "jgraphics"
import World from "./World"
import Coms from "./Communications"
import Controller from "./Controller"
import UI from "./UI"
import './App.css';

const server = "cybera.pelliott.ca:8049"
class MyGame extends JGraphics {
  // Setup is called once
  setup(){

    this.world = new World(80, 60)
    this.coms = new Coms(server+"/game/"+this.props.slug, this.world)
    this.ui = new UI()
    this.controller = new Controller(this.world, this)

    this.world.initView = () => ((game, world) => {
      game.draw.background("rgb(49,49,49)")
      game.view.setDimensions(world.size.width, world.size.height)
      game.view.setCenter(world.size.width/2, world.size.height/2)
      game.view.isYAxisUpPositive(true)
    })(this, this.world)
  }

  // Loop is called many times based on the fps
  loop(delta){
    this.world.update(delta)
    this.ui.draw(this.draw, this.world)
  }
}

export default class Game extends Component {
  render() {
    return (
      <div className="App">
        <MyGame slug={this.props.match.params.id} fps={40} id="jgraphic-panel" width="800" height="600" host={window.location.hostname}/>
      </div>
    )
  }
}


