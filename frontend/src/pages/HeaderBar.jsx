import { AppBar, Button, Paper } from '@material-ui/core';
import React from 'react';
import { useState } from 'react';
import { makeStyles } from '@material-ui/core';


const useStyles = makeStyles({
  root: {
    maxWidth: "60vw",
    minWidth: "40vw",
    maxHeight: "15vh",
    minHeight: "10vh",
    margin: 5,
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center"
  },
  button: {
    maxWidth: "18%",
    minWidth: "18%",
  }
})


function HeaderBar(props) {
  const classes = useStyles();
  const [buttons, setButtons] = useState([
    {
      "title": "Go Short",
      "disabled": true,    },
    {
      "title": "Go Long",
      "disabled": false,
    },
    {
      "title": "Go Custom",
      "disabled": false,
    },
  ])
  const handleOnClick = (index) => {
    let newButtons = [...buttons];
    newButtons.forEach(b => b.disabled = false)
    newButtons[index].disabled = true
    setButtons(newButtons)
    props.handlePageChange(newButtons[index].title)
  }
  return (
    <Paper>
      {
        buttons.map((button, index) => {
          return <Button key={index} variant="text" disabled={button.disabled} onClick={() => handleOnClick(index)} color="primary" className={classes.button}>{button.title}</Button>
        })
      }
    </Paper>
  );
}

export default HeaderBar;