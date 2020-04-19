import React from 'react';
import { ActionContext, Action } from './Action';

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
  const { data, loading, setAction, editing, setEdit, setEditingId } = React.useContext(
    ActionContext,
  );

  const handleEdit = (row: Action) => {
    setAction && setAction(row);
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
          {data.action.map((row) => (
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
