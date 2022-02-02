import URL from "./URL";
import {Grid} from '@material-ui/core'

function URLResults() {

    const Results = ["1123","12.12.12", "1412"]
    
    return (
        <Grid
            container
            spacing={0}
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