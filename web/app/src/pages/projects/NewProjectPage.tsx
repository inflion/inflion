import React, { useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { useForm } from 'react-hook-form';
import { Container, Divider, Grid, Typography } from '@material-ui/core';

import { useCountProjectLazyQuery, useCreateProjectMutation } from '../../graphql';

type FormData = {
  name: string;
  description: string;
};

export const NewProjectPage: React.FC = () => {
  const [redirect] = useState<string>('');
  const [projectAlreadyExists, setProjectAlreadyExists] = useState<boolean>(false);
  const [createProjectMutation, { data, loading, error }] = useCreateProjectMutation();
  const [countProject, { data: count }] = useCountProjectLazyQuery();

  const { register, handleSubmit, errors } = useForm<FormData>();

  const onSubmit = handleSubmit(({ name, description }) => {
    createProjectMutation({
      variables: { input: { name: name, description: description } },
    }).then((r) => console.log(r));
  });

  const onChaneName = (event: React.ChangeEvent<HTMLInputElement>) => {
    countProject({ variables: { name: event.target.value } });
  };

  useEffect(() => {
    if (count === undefined) {
      return;
    }

    setProjectAlreadyExists(count.project_aggregate.aggregate?.count === 1);
  }, [count]);

  if (data !== undefined) {
    return <div>done</div>;
  }

  if (redirect !== '') {
    return <Redirect to={redirect} />;
  }

  return (
    <Container maxWidth="lg">
      <div>{loading}</div>
      <div>{error}</div>
      <Typography variant="h4">Create a new project</Typography>
      <Typography variant="caption" display="block" gutterBottom>
        A project contains all project files, including cloud settings.
      </Typography>

      <Divider variant="middle" />

      <form noValidate autoComplete="off" onSubmit={onSubmit}>
        <Grid container spacing={3}>
          <Grid container item spacing={3}>
            <Grid item>
              <TextField
                id="standard-basic"
                label="Name"
                name="name"
                error={errors.name !== undefined}
                onChange={onChaneName}
                helperText={
                  errors.name &&
                  'The name may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen. '
                }
                inputRef={register({
                  required: true,
                  pattern: /^[^-][A-Za-z-]+[^-]$/i,
                })}
              />
            </Grid>
          </Grid>

          <Grid item>
            <TextField
              id="standard-basic"
              label="Description(optional)"
              name="description"
              error={errors.description !== undefined}
              helperText={errors.description && 'This field is optional.'}
              inputRef={register({ required: false })}
            />
          </Grid>

          <Grid item xs={12}>
            <Divider />
          </Grid>

          <Grid item xs={12}>
            <Button
              color="primary"
              variant="contained"
              type="submit"
              disabled={projectAlreadyExists}
            >
              Create project
            </Button>
            {projectAlreadyExists && <p>You already have project</p>}
          </Grid>
        </Grid>
      </form>
    </Container>
  );
};
