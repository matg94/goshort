import React from 'react';
import { Button, Card, Grid} from "@material-ui/core";
import TextBoxAction from './TextBoxAction';


function Toolbox() {
  return (
    <Grid
        container
        spacing={0}
        direction="column"
        alignItems="center"
        justify="center">
            <TextBoxAction description="Find where a go-short url points to" placeholder="Search" buttonText="Search"></TextBoxAction>
            <TextBoxAction description="Create a go-short url with a random link. These persist when you leave the page." placeholder="URL to redirect to" buttonText="Create"></TextBoxAction>
    </Grid>
  );
}

export default Toolbox;