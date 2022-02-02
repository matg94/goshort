import React from 'react';
import { Card, CardContent, Button, CardActions, Typography } from '@material-ui/core'
import {makeStyles} from '@material-ui/core';

const useStyles = makeStyles({
    root: {
      maxWidth: "80%",
      minWidth: "80%",
      maxHeight: "10vh",
      minHeight: "8vh",
      margin: 5,
      display: "flex",
      justifyContent: "space-between",
      alignItems: "center"
    },
    media: {
      height: 140,
    },
    Original: {
        fontSize: 12,
        justifyContent: "left"
    },
    Short: {
        fontSize: 16,
        justifyContent: "left"
    },
    Tag: {
        fontSize: 12,
        justifyContent: "right"
    }
  });

function URL(props) {
    const classes = useStyles();
    return (
        <Card className={classes.root}>
            <CardContent>
                <Typography className={classes.Short} component="p" color="textPrimary">{props.val.short}</Typography>
                <Typography className={classes.Original} component="p" color="textSecondary">{props.val.original}</Typography>
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