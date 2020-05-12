import { Box, Container, Tab, Tabs, Typography } from '@material-ui/core';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import React, { useState } from 'react';
import { useParams } from 'react-router-dom';
import { TabPanel } from '../../components/TabPanel';
import { useProjectByNameQuery } from '../../graphql';
import { Action } from './actions/Action';
import { Instances } from './instances/Instances';
import { AwsAccount } from './integration/AwsAccount';
import { Invitation } from './invitation/Invitation';
import { Notification } from './notification/Notification';
import { Settings } from './settings/Settings';

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
}));

export type Project = {
  id: bigint;
  name: string;
};

type ProjectContextProps = {
  project: Project;
};

export const ProjectContext = React.createContext<Partial<ProjectContextProps>>({});

export const ProjectPage = () => {
  const { projectName } = useParams();
  const [value, setValue] = useState(0);
  const { data } = useProjectByNameQuery({ variables: { name: projectName ?? '' } });
  const classes = useStyles();
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  const handleChange = (event: React.ChangeEvent<{}>, newValue: number) => {
    setValue(newValue);
  };

  if (data === undefined) {
    return <>loading</>;
  }

  const project: Project = { id: data.project[0].id, name: data.project[0].id };

  if (!project) {
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
          {/* Chart */}
          <Grid item xs={12}>
            <Paper className={classes.paper}></Paper>
          </Grid>
          <Grid item xs={12} md={8} lg={9}>
            <Paper className={fixedHeightPaper}></Paper>
          </Grid>
          {/* Recent Deposits */}
          <Grid item xs={12} md={4} lg={3}>
            <Paper className={fixedHeightPaper}></Paper>
          </Grid>
          {/* Recent Orders */}
          <Grid item xs={12}>
            <Paper className={classes.paper}></Paper>
          </Grid>
        </Grid>
        <Box pt={4}></Box>
      </Container>
      <Typography variant="h3">Responsive h3</Typography>

      <Tabs value={value} onChange={handleChange} indicatorColor="primary" textColor="primary">
        <Tab label="Instances" />
        <Tab label="Notifications" />
        <Tab label="Actions" />
        <Tab label="AwsAccount" />
        <Tab label="Invitations" />
        <Tab label="Settings" />
      </Tabs>

      <ProjectContext.Provider value={{ project: project }}>
        <TabPanel value={value} index={0}>
          <Instances />
        </TabPanel>

        <TabPanel value={value} index={1}>
          <Notification />
        </TabPanel>

        <TabPanel value={value} index={2}>
          <Action />
        </TabPanel>

        <TabPanel value={value} index={3}>
          <AwsAccount />
        </TabPanel>

        <TabPanel value={value} index={4}>
          <Invitation />
        </TabPanel>

        <TabPanel value={value} index={5}>
          <Settings />
        </TabPanel>
      </ProjectContext.Provider>
    </>
  );
};
