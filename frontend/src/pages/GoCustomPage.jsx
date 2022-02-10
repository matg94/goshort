import React from 'react';
import { Grid} from "@material-ui/core";
import CustomTextBoxAction from '../components/CustomTextBoxAction';


function GoCustomPage() {
  return (
    <Grid
    container
    spacing={0}
    direction="column"
    alignItems="center"
    justify="center">
      <CustomTextBoxAction 
        description="Shorten a URL to a custom go short one" 
        submitURL="custom" 
        placeholderOne="Enter Full URL" 
        placeholderTwo="Enter Desired URL"
        buttonText="GO CUSTOM"></CustomTextBoxAction>
    </Grid>
  );
}

export default GoCustomPage;