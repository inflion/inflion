import React from 'react';
import { useParams } from 'react-router-dom';
import { Chart } from './chart';
import { useInstanceQuery } from '../../../graphql/generated';
import { Typography } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

// const useStyles = makeStyles({
//   table: {
//     minWidth: 650,
//   },
// });

// function createData(sid: string, name: string) {
//   return { sid, name };
// }

// const rows = [
//   createData('sg-05a8a24adb2084952', 'c16e'),
//   createData('sg-05a8a24adb2084952', 'c16e'),
//   createData('sg-05a8a24adb2084952', 'c16e'),
//   createData('sg-05a8a24adb2084952', 'c16e'),
// ];

export const InstanceDetail = () => {
  // const classes = useStyles();

  // const { projectId, instanceId } = useParams();

  // const { data, loading } = useInstanceQuery({
  //   variables: { projectId: '', instanceId: '' },
  // });

  // if (loading) {
  //   return <>loading</>;
  // }
  // return (
  //   <>
  //     {data && data.instance.length > 0 ? (
  //       <>
  {
    /* <Typography variant="h4">{data?.instance[0].name}</Typography>
          <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
              <TableHead>
                <TableRow>
                  <TableCell>Security group ID</TableCell>
                  <TableCell>Security group name</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {rows.map((row, index) => (
                  <TableRow key={index}>
                    <TableCell>{row.sid}</TableCell>
                    <TableCell>{row.name}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          <Typography variant="caption" display="block" gutterBottom>
            {instanceId}
          </Typography>
          <Chart instanceId={instanceId} type="CPUUtilization" />
          <Chart instanceId={instanceId} type="NetworkOut" />
          <Chart instanceId={instanceId} type="NetworkIn" />
          <Chart instanceId={instanceId} type="DiskWriteBytes" />
          <Chart instanceId={instanceId} type="DiskReadBytes" /> */
  }
  //       </>
  //     ) : null}
  //   </>
  // );
};
