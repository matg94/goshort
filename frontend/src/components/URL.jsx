import { MDBCol } from "mdbreact";
import { Card, CardContent, Button, CardActions, Typography } from '@material-ui/core'
import {makeStyles} from '@material-ui/core';

const useStyles = makeStyles({
    root: {
      maxWidth: "40vw",
      minWidth: "20vw",
      maxHeight: "8vh",
      margin: 5,
      display: "flex",
      justifyContent: "space-between",
      alignItems: "center"
    },
    media: {
      height: 140,
    },
  });

function URL(props) {
    const classes = useStyles();
    return (
        <Card className={classes.root}>
            <CardContent>
                <Typography display="inline" component="p" color="textPrimary">{props.val}</Typography>
            </CardContent>
            <CardContent>
                <CardActions>
                    <Button 
                        variant="outlined"
                        color="secondary"
                        size="small"
                        >
                            Delete
                    </Button>
                </CardActions>
            </CardContent>
        </Card>
    );
}

export default URL;