import './App.css';
import React from 'react';
import MainPage from './pages/MainPage'
import {createTheme, ThemeProvider} from '@material-ui/core/styles';

const theme = createTheme({
  status: {
      danger: "#ff5555",
      success: "#50fa7b"
  },
  palette: {
      type: 'dark',
      primary: {
          main: "#bd93f9"
      },
      secondary: {
          main: "#ff79c6"
      },
      background: {
          default: "#282a36",
          paper: "#44475a"
      }
  },
  text: {
    primary: "#f8f8f2"
  }
})

function App() {
  return (
    <ThemeProvider theme={theme}>
      <div className="App">
        <MainPage/>
      </div>
    </ThemeProvider>
  );
}

export default App;
