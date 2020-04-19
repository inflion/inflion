import React, { useEffect, useState, useContext } from 'react';

import {
  Button,
  Checkbox,
  Paper,
  Table,
  TableBody,
  TableCell,
  TablePagination,
  TableRow,
  Typography,
} from '@material-ui/core';

import { Link as RouterLink } from 'react-router-dom';

import { useAwsInstancesQuery } from '../../../graphql';
import { ProjectContext } from '../ProjectPage';
import { useStyles } from './styles';
import { EnhancedTableToolbar } from './EnhancedTableToolbar';
import { EnhancedTableHead } from './EnhancedTableHead';
import { Tags } from '../../../graphql/generated';

export interface Data {
  instanceId: string;
  name: string;
  status: string;
  privateAddress: string;
  publicAddress: string;
  service: string;
}

function createData(
  instanceId: string,
  name: string,
  status: string,
  privateAddress: string,
  publicAddress: string,
  service: string,
): Data {
  return { instanceId, name, status, privateAddress, publicAddress, service };
}

function desc<T>(a: T, b: T, orderBy: keyof T) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

function stableSort<T>(array: T[], cmp: (a: T, b: T) => number) {
  const stabilizedThis = array.map((el, index) => [el, index] as [T, number]);
  stabilizedThis.sort((a, b) => {
    const order = cmp(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  return stabilizedThis.map((el) => el[0]);
}

export type Order = 'asc' | 'desc';

function getSorting<K extends keyof any>(
  order: Order,
  orderBy: K,
): (a: { [key in K]: number | string }, b: { [key in K]: number | string }) => number {
  return order === 'desc' ? (a, b) => desc(a, b, orderBy) : (a, b) => -desc(a, b, orderBy);
}

export const Instances = () => {
  const { project } = useContext(ProjectContext);

  const classes = useStyles();

  const [order, setOrder] = useState<Order>('asc');
  const [orderBy, setOrderBy] = useState<keyof Data>('name');
  const [selected, setSelected] = useState<string[]>([]);
  const [page, setPage] = useState(0);
  const [dense] = useState(false);
  const [rowsPerPage, setRowsPerPage] = useState(5);
  const [rows, setRows] = useState<Data[]>([]);

  const { loading, data, refetch } = useAwsInstancesQuery({
    variables: { projectId: Number(project?.id) },
  });

  const reloadData = () => {
    refetch();
  };

  useEffect(() => {
    const mappedData: Data[] = [];

    if (loading) {
      return;
    }

    if (data === undefined) {
      return;
    }

    data.aws_instances?.forEach(
      (instance: {
        instanceId: string;
        name: string;
        status: string;
        privateAddress: string;
        publicAddress: string;
        tags: Tags;
      }) => {
        const service = instance.tags.tags.find((x) => x?.key === 'Service')?.value;
        mappedData.push(
          createData(
            instance.instanceId,
            instance.name,
            instance.status,
            instance.privateAddress,
            instance.publicAddress,
            service ?? '',
          ),
        );
      },
    );

    setRows(mappedData);
  }, [loading, data]);

  const handleRequestSort = (event: React.MouseEvent<unknown>, property: keyof Data) => {
    const isDesc = orderBy === property && order === 'desc';
    setOrder(isDesc ? 'asc' : 'desc');
    setOrderBy(property);
  };

  const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      const newSelecteds = rows.map((n) => n.name);
      setSelected(newSelecteds);
      return;
    }
    setSelected([]);
  };

  const handleClick = (event: React.MouseEvent<unknown>, instanceId: string) => {
    const selectedIndex = selected.indexOf(instanceId);
    let newSelected: string[] = [];

    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selected, instanceId);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selected.slice(1));
    } else if (selectedIndex === selected.length - 1) {
      newSelected = newSelected.concat(selected.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selected.slice(0, selectedIndex),
        selected.slice(selectedIndex + 1),
      );
    }

    setSelected(newSelected);
  };

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const isSelected = (name: string) => selected.indexOf(name) !== -1;

  const emptyRows = rowsPerPage - Math.min(rowsPerPage, rows.length - page * rowsPerPage);

  return (
    <div className={classes.root}>
      <Typography variant="h4">Instances</Typography>
      <Paper className={classes.paper}>
        <EnhancedTableToolbar numSelected={selected.length} />
        <div className={classes.tableWrapper}>
          <Table
            className={classes.table}
            aria-labelledby="tableTitle"
            size={dense ? 'small' : 'medium'}
            aria-label="enhanced table"
          >
            <EnhancedTableHead
              classes={classes}
              numSelected={selected.length}
              order={order}
              orderBy={orderBy}
              onSelectAllClick={handleSelectAllClick}
              onRequestSort={handleRequestSort}
              rowCount={rows.length}
            />
            <TableBody>
              {stableSort(rows, getSorting(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map((row, index) => {
                  const isItemSelected = isSelected(row.instanceId);
                  const labelId = `enhanced-table-checkbox-${index}`;

                  return (
                    <TableRow
                      hover
                      onClick={(event) => handleClick(event, row.instanceId)}
                      role="checkbox"
                      aria-checked={isItemSelected}
                      tabIndex={-1}
                      key={row.instanceId}
                      selected={isItemSelected}
                    >
                      <TableCell padding="checkbox">
                        <Checkbox
                          checked={isItemSelected}
                          inputProps={{ 'aria-labelledby': labelId }}
                        />
                      </TableCell>
                      <TableCell component="th" id={labelId} scope="row" padding="none">
                        <RouterLink to={`/instance/${row.instanceId}`}>{row.instanceId}</RouterLink>
                      </TableCell>
                      <TableCell align="right">{row.name}</TableCell>
                      <TableCell align="right">{row.service}</TableCell>
                      <TableCell align="right">{row.status}</TableCell>
                      <TableCell align="right">{row.privateAddress}</TableCell>
                      <TableCell align="right">{row.publicAddress}</TableCell>
                    </TableRow>
                  );
                })}
              {emptyRows > 0 && (
                <TableRow style={{ height: (dense ? 33 : 53) * emptyRows }}>
                  <TableCell colSpan={6} />
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
        <TablePagination
          rowsPerPageOptions={[5, 10, 25]}
          component="div"
          count={rows.length}
          rowsPerPage={rowsPerPage}
          page={page}
          backIconButtonProps={{
            'aria-label': 'previous page',
          }}
          nextIconButtonProps={{
            'aria-label': 'next page',
          }}
          onChangePage={handleChangePage}
          onChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Paper>

      <Button variant="contained" className={classes.button} onClick={reloadData}>
        reload
      </Button>
    </div>
  );
};
