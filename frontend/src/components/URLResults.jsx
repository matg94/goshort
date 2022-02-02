import React from 'react';
import URL from "./URL";
import {Grid, Card} from '@material-ui/core'

function URLResults(props) {

    const Results = props.URLs
    
    return (
        <Grid
            container
            spacing={1}
            direction="column"
            alignItems="center"
            justify="center">
        {Results.map(url => {
            return  (
                    <URL val={url}/>
            )
        })}
        </Grid>
    );
}

export default URLResults;