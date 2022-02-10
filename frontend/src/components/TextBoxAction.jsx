import React from 'react';
import { useState } from 'react';
import axios from 'axios';
import { Button, makeStyles, Card, Paper } from "@material-ui/core";
import URLResults from './URLResults';
import config from '../config.json';
import {NotificationManager} from 'react-notifications';

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
  container: {
    maxWidth: "60vw",
    minWidth: "40vw",
    maxHeight: "30vh",
    minHeight: "20vh",
    margin: 10,
  },
  searchBox: {
    maxWidth: "68%",
    minWidth: "68%",
    minHeight: "4vh",
    margin: "3%",
  },
  button: {
    height: "4vh",
    maxWidth: "18%",
    minWidth: "18%",
    margin: "3%",
  }
})

function TextBoxAction(props) {
  const [ inputText, setInputText ] = useState("")
  const [ urls, setUrls ] = useState([])
  const handleInput = event => {
    setInputText(event.target.value)
  }
  const handleRemoveUrl = (short, original) => {
    const updatedURLS = urls.filter(
      url => url.original !== original && url.short !== short
    )
    setUrls(updatedURLS)
  }
  const handleSearch = () => {
    axios
        .post(`${config.SERVER_URL}/${props.submitURL}`, {"url": inputText})
        .then(res => {
          if (urls.filter(url => url.original === inputText && url.short === res.data.url).length > 0) {
            return
          }
          setUrls([...urls, {
            "original": inputText,
            "short": res.data.url,
          }])
          NotificationManager.success("", "Success!", 3000)
        })
        .catch(err => { NotificationManager.error(`${err}`, "Failed!", 5000) })
  }

  const classes = useStyles(useStyles)
  return (
    <React.Fragment>
      <Paper className={classes.container}>
        <h4>{props.description}</h4>
        <Card className={classes.root} elevation={0}>
          <input className={classes.searchBox} type="text" placeholder={props.placeholder} onChange={handleInput} aria-label="TextBoxAction" />
          <Button className={classes.button} size="sm" variant="outlined" color="primary" onClick={handleSearch}>{props.buttonText}</Button>
        </Card>
      </Paper>
      <URLResults handleRemoveUrl={handleRemoveUrl} URLs={urls}></URLResults>
    </React.Fragment>
  );
}

export default TextBoxAction;