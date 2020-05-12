import {
  Box,
  Button,
  Container,
  Grid,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import AddIcon from '@material-ui/icons/Add';
import clsx from 'clsx';
import React, { useEffect } from 'react';
import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom';
import { useProjectsQuery } from '../../graphql';

const ProjectLink = React.forwardRef<HTMLAnchorElement, RouterLinkProps>((props, ref) => (
  <RouterLink innerRef={ref} {...props} />
));

const useStyles = makeStyles((theme) => ({
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  fixedHeight: {
    height: 240,
  },
  table: {
    minWidth: 650,
  },
}));

type Project = {
  name: string;
  project_id: bigint;
  description: string | null | undefined;
};

export const Projects = () => {
  const classes = useStyles();
  const [projects, setProjects] = React.useState<Project[]>([]);
  const { data, loading } = useProjectsQuery();
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  useEffect(() => {
    if (!data) {
      return;
    }
    const projects: Array<Project> = [];
    data.project.forEach((e) => {
      projects.push({ name: e.name, project_id: e.id, description: e.description });
    });
    setProjects(projects);
  }, [data]);

  if (loading) {
    return <>loading</>;
  }

  return (
    <>
      <Box color="text.primary" fontSize={30} margin={4}>
        Projects
        <hr />
      </Box>

      <Container maxWidth="lg" className={classes.container}>
        <Grid container spacing={3}>
          <Grid item xs={12}>
            <Grid justify="flex-end" container>
              <Box mb={2}>
                <Button variant="contained" color="primary">
                  <AddIcon /> Create
                </Button>
              </Box>
            </Grid>
            <Paper className={classes.paper}>
              <Table className={classes.table} aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell>Name</TableCell>
                    <TableCell>Description</TableCell>
                    <TableCell align="right">Actions</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {projects.map((row) => (
                    <TableRow key={row.project_id.toString()}>
                      <TableCell component="th" scope="row">
                        {row.name}
                      </TableCell>
                      <TableCell>{row.description}</TableCell>
                      <TableCell align="right">
                        <Button
                          variant="contained"
                          color="primary"
                          component={ProjectLink}
                          to={`/projects/${row.name}`}
                        >
                          Show
                        </Button>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </Paper>
          </Grid>
        </Grid>
      </Container>
    </>
  );
};
