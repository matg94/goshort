import React from 'react';
import { Grid} from "@material-ui/core";
import TextBoxAction from '../components/TextBoxAction';


function GoShortPage() {
  return (
    <Grid
    container
    spacing={0}
    direction="column"
    alignItems="center"
    justify="center">
      <TextBoxAction 
        description="Shorten a URL to a go.short one" 
        submitURL="short" 
        placeholder="Enter Full URL" 
        buttonText="GO SHORT"></TextBoxAction>
    </Grid>
  );
}

export default GoShortPage;