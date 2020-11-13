import React from 'react';

import { useForm } from 'react-hook-form';
import { makeStyles, Theme } from '@material-ui/core/styles';

import { useCreateProjectInvitationMutation, useInvitationQuery } from '../../../graphql';

import { ProjectContext } from '../ProjectPage';
import { Typography } from '@material-ui/core';
import {
  Button,
  Grid,
  Paper,
  TextField,
  TableContainer,
  Table,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  Divider,
} from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    display: 'flex',
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
  paper: {
    padding: theme.spacing(3, 2),
  },
  idGenerator: {
    display: 'flex',
    alignItems: 'center',
  },
  idGeneratorButton: {
    marginLeft: 30,
  },
  idGeneratorTextField: {
    width: 300,
    fontFamily: '"Lucida Console", "Monaco", "monospace"',
  },
}));

type FormData = {
  email: string;
};

export const Invitation = () => {
  const { project } = React.useContext(ProjectContext);

  const classes = useStyles();

  const { register, handleSubmit, setValue, errors } = useForm<FormData>();

  const [createProjectInvitation] = useCreateProjectInvitationMutation();
  const { data, loading, refetch } = useInvitationQuery({ variables: { projectId: project?.id } });

  const onSubmit = handleSubmit(({ email }) => {
    const input = {
      mail_address: email,
      project_id: project?.id,
    };

    createProjectInvitation({ variables: { input } }).then(() => {
      setValue('email', '');
      refetch();
    });
  });

  if (loading || !data) {
    return <p>loading</p>;
  }

  return (
    <>
      <Typography variant="h5">Invitations</Typography>
      <TableContainer component={Paper}>
        <Table aria-label="table">
          <TableHead>
            <TableRow>
              <TableCell></TableCell>
              <TableCell align="right"></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.project_invitation.map((row) => (
              <TableRow key={row.mail_address}>
                <TableCell component="th" scope="row">
                  {row.mail_address}
                </TableCell>
                <TableCell align="right"></TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      <Divider />

      <Typography variant="h5">Add a new member</Typography>
      <Paper className={classes.paper}>
        <form noValidate autoComplete="off" onSubmit={onSubmit}>
          <Grid container alignItems="flex-start" spacing={2}>
            <Grid item xs={12}>
              <TextField
                id="email"
                label="email"
                name="email"
                error={errors.email !== undefined}
                helperText={errors.email && 'Email is invalid or already taken'}
                inputRef={register({
                  required: true,
                  pattern: /^\S+@\S+$/,
                })}
              />
            </Grid>

            <Grid item style={{ marginTop: 16 }}>
              <Button variant="contained" color="primary" type="submit">
                add
              </Button>
            </Grid>
          </Grid>
        </form>
      </Paper>
    </>
  );
};
