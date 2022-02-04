import React from 'react';
import { Button, Card, Grid} from "@material-ui/core";
import TextBoxAction from '../components/TextBoxAction';


function GoLongPage() {
  return (
    <TextBoxAction 
    description="Find where a go.short URL points to" 
    submitURL="golong" 
    placeholder="Enter Full Go.Short URL" 
    buttonText="GO LONG"></TextBoxAction>
  );
}

export default GoLongPage;
