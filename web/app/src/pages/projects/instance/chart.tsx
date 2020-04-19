import React from 'react';
import { Line } from '@nivo/line';
import { useGetMetricQuery } from '../../../graphql';
import { Datum } from '@nivo/line';

type ChartProps = {
  instanceId: string;
  type: string;
};

export const Chart: React.FC<ChartProps> = ({ instanceId, type }) => {
  const queryParam = {
    pollInterval: 10000,
    variables: { type: type, instanceId: instanceId, limit: 30 },
  };

  const { data, loading } = useGetMetricQuery(queryParam);

  if (loading) return <p>loading</p>;

  let metricData: Datum[] = [];

  if (data && data.metrics) {
    metricData = data.metrics.map(({ time, value }) => {
      return {
        x: new Date(time),
        y: value,
      };
    });
  }

  return (
    <Line
      animate={true}
      width={800}
      height={300}
      data={[
        {
          id: type,
          data: metricData,
          color: '#B00',
        },
      ]}
      margin={{
        top: 50,
        right: 50,
        bottom: 50,
        left: 50,
      }}
      axisTop={null}
      axisRight={null}
      xScale={{
        type: 'time',
        format: 'native',
      }}
      xFormat="time:%H:%M"
      yScale={{
        type: 'linear',
        stacked: false,
      }}
      axisLeft={{
        legend: type,
        legendOffset: 12,
      }}
      axisBottom={{
        format: '%m/%d %H:%M',
        tickValues: 5,
        legendOffset: 0,
      }}
      isInteractive={true}
      pointLabel={'y'}
      enablePoints={true}
      useMesh={true}
      curve={'linear'}
    />
  );
};
