import { useQuery } from '@apollo/react-hooks';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { gql } from 'apollo-boost';
import React, { useEffect, useState } from 'react';
import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom';

interface Data {
  name: string;
}

interface OrgListProps {
  update: (notify: () => void) => void;
}

const OrgLink = React.forwardRef<HTMLAnchorElement, RouterLinkProps>((props, ref) => (
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

function createData(name: string): Data {
  return { name };
}

export const OrgsList: React.FC<OrgListProps> = props => {
  const classes = useStyles();

  const [rows, setRows] = useState<Data[]>([]);

  const { loading, data, refetch } = useQuery(gql`
    {
      organizations {
        name
      }
    }
  `);

  props.update(() => {
    refetch();
  });

  useEffect(() => {
    if (loading) {
      return;
    }

    const mappedData: Data[] = [];

    if (data === undefined) {
      return;
    }

    data.organizations.forEach((org: { name: string }) => {
      mappedData.push(createData(org.name));
    });

    setRows(mappedData);
  }, [loading, data]);

  return (
    <>
      <Paper className={classes.root}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
              <TableCell></TableCell>
              <TableCell align="right">Name</TableCell>
              <TableCell align="right">Actions</TableCell>
            </TableRow>
          </TableHead>

          <TableBody>
            {rows.map(row => (
              <TableRow key={row.name}>
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="right">{row.name}</TableCell>
                <TableCell align="right">
                  <Button variant="contained" color="secondary" component={OrgLink} to={`/${row.name}`}>
                    Show
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>
    </>
  );
};
