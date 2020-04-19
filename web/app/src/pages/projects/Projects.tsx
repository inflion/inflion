import React, { useEffect } from 'react';

import { makeStyles } from '@material-ui/core/styles';

import { Button, Paper, Table, TableBody, TableCell, TableHead, TableRow } from '@material-ui/core';
import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom';

import { useProjectsQuery } from '../../graphql';

const ProjectLink = React.forwardRef<HTMLAnchorElement, RouterLinkProps>((props, ref) => (
  <RouterLink innerRef={ref} {...props} />
));

const useStyles = makeStyles({
  root: {
    width: '100%',
    overflowX: 'auto',
  },
  table: {
    minWidth: 650,
  },
});

type Project = {
  name: string;
  project_id: bigint;
  description: string | null | undefined;
};

export const Projects = () => {
  const classes = useStyles();
  const [projects, setProjects] = React.useState<Project[]>([]);
  const { data, loading } = useProjectsQuery();

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
    <Paper className={classes.root}>
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
  );
};
