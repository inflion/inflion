import React from 'react';
import { Typography, Grid } from '@material-ui/core';
import { Status } from './Status';
import styled from 'styled-components';
import { useGetServicesQuery } from '../../graphql/generated';

const StatusGridContainer = styled.div`
  padding: 8px;
  padding-top: 16px;
`;

export const Dashboard = () => {
  const { data, loading } = useGetServicesQuery();

  if (loading) {
    return <>loading</>;
  }

  return (
    <>
      <Typography variant="h5">Status Dashboard</Typography>
      <StatusGridContainer>
        {data?.service.length === 0 && <p>You have no services yet</p>}
        {data?.service.length !== 0 && (
          <Grid container spacing={2}>
            {data?.service.map((service) => (
              <Grid key={service.id} item>
                <Status name={service.name} type="ok"></Status>
              </Grid>
            ))}
          </Grid>
        )}
      </StatusGridContainer>
    </>
  );
};
