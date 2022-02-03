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
            <TextBoxAction 
              description="Find where a go.short URL points to" 
              submitURL="golong" 
              placeholder="Enter Full Go.Short URL" 
              buttonText="GO LONG"></TextBoxAction>
            <TextBoxAction 
              description="Shorten a URL to a go.short one" 
              submitURL="goshort" 
              placeholder="Enter Full URL" 
              buttonText="GO SHORT"></TextBoxAction>
    </Grid>
  );
}

export default Toolbox;