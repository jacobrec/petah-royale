import React, { Component } from 'react';
import JGraphics from "jgraphics"
import World from "./World"
import Coms from "./Communications"
import Controller from "./Controller"
import UI from "./UI"
import './App.css';

class MyGame extends JGraphics {
  // Setup is called once
  setup(){
    this.draw.background("rgb(49,49,49)")

    this.coms = new Coms("localhost:8049")
    this.world = new World(80, 60, this.coms)
    this.ui = new UI()
    this.controller = new Controller(this.world, this)

    this.view.setDimensions(this.world.size.width, this.world.size.height)
    this.view.setCenter(this.world.size.width/2, this.world.size.height/2)
    this.view.isYAxisUpPositive(true)
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
        <MyGame fps={30} id="jgraphic-panel" width="800" height="600"/>
      </div>
    )
  }
}

