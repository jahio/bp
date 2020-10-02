import React, { useState } from 'react';
import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import NewEntry from './newEntry'
import Graph from './graph'

function App() {
  const initialState = { systolic: "", diastolic: "", heartrate: "" };
  const [state, setState] = useState(initialState);

  return (
    <Router>
      <Link to="/">New Entry</Link>
      <Link to="/graph">Graphs</Link>

      <Switch>
        <Route path="/">
          <NewEntry state={state} setState={setState} />
        </Route>
        <Route path="/graph">
          <Graph />
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
