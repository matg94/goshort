import './App.css';
import React from 'react';
import URLResults from './components/URLResults';
import HeaderBar from './pages/HeaderBar';
import Toolbox from './components/Toolbox';
import {Paper} from '@material-ui/core';

function App() {
  return (
    <div className="App">
      <HeaderBar/>
      <Toolbox></Toolbox>
    </div>
  );
}

export default App;
