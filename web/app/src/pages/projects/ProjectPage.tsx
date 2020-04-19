import React, { useState } from 'react';

import { TabPanel } from '../../components/TabPanel';
import { Tabs, Tab } from '@material-ui/core';

import { Instances } from './instances/Instances';
import { Notification } from './notification/Notification';
import { Action } from './actions/Action';
import { AwsAccount } from './integration/AwsAccount';
import { Invitation } from './invitation/Invitation';
import { Settings } from './settings/Settings';

import { useParams } from 'react-router-dom';
import { useProjectByNameQuery } from '../../graphql';

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
