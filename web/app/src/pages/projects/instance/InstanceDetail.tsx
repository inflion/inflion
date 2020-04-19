import React from 'react';
import { useParams } from 'react-router-dom';
import { Chart } from './chart';
import { useInstanceQuery } from '../../../graphql/generated';
import { Typography } from '@material-ui/core';

export const InstanceDetail = () => {
  const { instanceId: paramsInstanceId } = useParams();
  const instanceId = paramsInstanceId ?? '';

  const { data, loading } = useInstanceQuery({
    variables: { projectId: 1, instanceId: instanceId },
  });

  if (loading) {
    return <>loading</>;
  }

  return (
    <>
      <Typography variant="h4">{data?.instance[0].name}</Typography>
      <Typography variant="caption" display="block" gutterBottom>
        {instanceId}
      </Typography>

      <Chart instanceId={instanceId} type="CPUUtilization" />
      <Chart instanceId={instanceId} type="NetworkOut" />
      <Chart instanceId={instanceId} type="NetworkIn" />
      <Chart instanceId={instanceId} type="DiskWriteBytes" />
      <Chart instanceId={instanceId} type="DiskReadBytes" />
    </>
  );
};
