import React from 'react';
import { Grid} from "@material-ui/core";
import TextBoxAction from '../components/TextBoxAction';


function GoLongPage() {
  return (
    <Grid
    container
    spacing={0}
    direction="column"
    alignItems="center"
    justify="center">
        <TextBoxAction 
        description="Find where a go.short URL points to" 
        submitURL="long" 
        placeholder="Enter Full Go.Short URL" 
        buttonText="GO LONG"></TextBoxAction>
    </Grid>
  );
}

export default GoLongPage;
