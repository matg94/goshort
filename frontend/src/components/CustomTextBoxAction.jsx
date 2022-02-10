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
      maxHeight: "30vh",
      minHeight: "20vh",
      margin: 5,
      justifyContent: "space-between",
      alignItems: "center"
  },
  container: {
    maxWidth: "60vw",
    minWidth: "40vw",
    maxHeight: "60vh",
    minHeight: "40vh",
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
    maxWidth: "30%",
    minWidth: "30%",
    margin: "3%",
  }
})

function CustomTextBoxAction(props) {
  const [ originalURLInput, setoriginalURLInput ] = useState("")
  const [ desiredURLInput, setdesiredURLInput ] = useState("")
  const [ urls, setUrls ] = useState([])

  const handleRemoveUrl = (short, original) => {
    const updatedURLS = urls.filter(
      url => url.original !== original && url.short !== short
    )
    setUrls(updatedURLS)
  }

  const handleSubmit = () => {
    axios
        .post(`${config.SERVER_URL}/${props.submitURL}`, {
            "url": originalURLInput,
            "short": desiredURLInput
        })
        .then(res => {
          if (urls.filter(url => url.original === originalURLInput && url.short === res.data.url).length > 0) {
            return
          }
          setUrls([...urls, {
            "original": originalURLInput,
            "short": res.data.url,
          }])
          NotificationManager.success("URL Shortened", "Success!", 3000)
        })
        .catch(err => { NotificationManager.error(`${err}`, "Failed!", 5000) })
  }

  const classes = useStyles(useStyles)
  return (
    <React.Fragment>
      <Paper className={classes.container}>
        <h4>{props.description}</h4>
        <Card className={classes.root} elevation={0}>
          <input 
              className={classes.searchBox} 
              type="text" 
              placeholder={props.placeholderOne} 
              onChange={(event) => setoriginalURLInput(event.target.value)}
              aria-label="CustomTextBoxAction"
          />
          <input 
              className={classes.searchBox}
              type="text"
              placeholder={props.placeholderTwo} 
              onChange={(event) => setdesiredURLInput(event.target.value)}
              aria-label="CustomTextBoxAction"
          />
          <Button className={classes.button} size="sm" variant="outlined" color="primary" onClick={handleSubmit}>{props.buttonText}</Button>
        </Card>
      </Paper>
      <URLResults handleRemoveUrl={handleRemoveUrl} URLs={urls}></URLResults>
    </React.Fragment>
  );
}

export default CustomTextBoxAction;