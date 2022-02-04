import React from 'react';
import { Button, Card, Grid} from "@material-ui/core";
import TextBoxAction from '../components/TextBoxAction';


function GoShortPage() {
  return (
    <TextBoxAction 
      description="Shorten a URL to a go.short one" 
      submitURL="goshort" 
      placeholder="Enter Full URL" 
      buttonText="GO SHORT"></TextBoxAction>
  );
}

export default GoShortPage;