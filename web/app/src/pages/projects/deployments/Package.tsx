import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
  card: {
    maxWidth: 345,
  },
  media: {
    height: 140,
  },
});

export const Package = () => {
  const classes = useStyles();

  return (
    <Card className={classes.card}>
      <CardActionArea>
        <CardMedia className={classes.media} image="/images/package.jpg" title="Image" />
        <CardContent>
          <Typography gutterBottom variant="h5" component="h2">
            Nomad Cluster
          </Typography>
          <Typography variant="body2" color="textSecondary" component="p">
            This package deploy Nomad and Consul cluster to AWS.
          </Typography>
        </CardContent>
      </CardActionArea>
      <CardActions>
        <Button size="small" color="primary">
          Deploy
        </Button>
        <Button size="small" color="primary">
          Details
        </Button>
      </CardActions>
    </Card>
  );
};
