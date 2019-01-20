import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route } from "react-router-dom";
import Game from './Game';
import Setup, { Join } from './Setup';
import './index.css';


class App extends React.Component {
  render() {
    return (
      <Router>
        <React.Fragment>
          <Route path="/" exact component={Setup} />
          <Route path="/join" exact component={Join} />
          <Route path="/game/:id" component={Game} />
        </React.Fragment>
      </Router>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('root'));

