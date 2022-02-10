import React from 'react';
import URL from "./URL";
import {Grid} from '@material-ui/core'

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
                    <URL onClick={props.handleRemoveUrl} val={url}/>
            )
        })}
        </Grid>
    );
}

export default URLResults;