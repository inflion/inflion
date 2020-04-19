import React from 'react';
import { SlackSettingContext, Setting } from './Slack';

import {
  Button,
  Paper,
  Table,
  TableContainer,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
} from '@material-ui/core';

export const Settings = () => {
  const { data, loading, setSetting, editing, setEdit, setEditingId } = React.useContext(
    SlackSettingContext,
  );

  const handleEdit = (row: Setting) => {
    setSetting && setSetting(row);
    setEdit && setEdit(true);
    row.id && setEditingId && setEditingId(row.id);
  };

  if (loading || !data) {
    return <></>;
  }

  return (
    <TableContainer component={Paper}>
      <Table aria-label="table">
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell align="right"></TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {data.slack_webhook.map((row) => (
            <TableRow key={row.id}>
              <TableCell component="th" scope="row">
                {row.name}
              </TableCell>
              <TableCell align="right">
                {editing && (
                  <Button disabled={true} onClick={() => handleEdit(row)}>
                    edit
                  </Button>
                )}
                {!editing && <Button onClick={() => handleEdit(row)}>edit</Button>}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};
