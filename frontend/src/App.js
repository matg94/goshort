import './App.css';
import React from 'react';
import URLResults from './components/URLResults';
import Toolbox from './components/Toolbox';
import {Paper} from '@material-ui/core';

function App() {
  return (
    <div className="App">
      <Paper>
        <h1>Go Short</h1>
      </Paper>
      <Toolbox></Toolbox>
    </div>
  );
}

export default App;
