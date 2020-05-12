import React from 'react';
import { makeStyles, Theme } from '@material-ui/core/styles';
import styled from 'styled-components';
import { Paper, Grid, TextField, Button, Typography } from '@material-ui/core';
import { useForm } from 'react-hook-form';
import {
  useCreateAwsAccountMutation,
  useAwsAccountByProjectIdQuery,
  AwsAccount as GraphqlAwsAccount,
} from '../../../graphql';

import { ProjectContext } from '../ProjectPage';

type SelectedAwsAccount = Pick<
  GraphqlAwsAccount,
  'id' | 'account_id' | 'role_name' | 'external_id'
>;

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
  awsAccountId: string;
  awsExternalId: string;
  awsRoleName: string;
};

const randomString = (length: number) => {
  let result = '';
  const characters = 'abcdefghijklmnopqrstuvwxyz0123456789';
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
};

const IdGeneratorTextField = styled(TextField)`
  width: 300px;
  font-family: 'Lucida Console', 'Monaco', 'monospace';
`;

const StyledButton = styled(Button)``;
const AwsExternalIdLength = 32;

export const AwsAccount = () => {
  const { project } = React.useContext(ProjectContext);

  const classes = useStyles();

  const [selectedAwsAccount, setSelectedAwsAccount] = React.useState<SelectedAwsAccount>({
    id: '',
    account_id: '',
    external_id: '',
    role_name: '',
  });

  const { register, handleSubmit, setValue, errors } = useForm<FormData>();

  const [createAwsAccount] = useCreateAwsAccountMutation();
  const { data: awsAccount } = useAwsAccountByProjectIdQuery({
    variables: { projectId: project?.id },
  });

  const handleGenerateId = () => {
    setValue('awsExternalId', randomString(AwsExternalIdLength));
  };

  const onSubmit = handleSubmit(({ awsAccountId, awsRoleName, awsExternalId }) => {
    const input = {
      id: selectedAwsAccount.id,
      account_id: awsAccountId,
      role_name: awsRoleName,
      external_id: awsExternalId,
      project_id: project?.id,
    };

    if (selectedAwsAccount.id === '') {
      delete input.id;
    }

    createAwsAccount({ variables: { input: input } });
  });

  if (awsAccount === undefined) {
    return <>loading</>;
  }

  if (selectedAwsAccount.id === '' && awsAccount.aws_account.length !== 0) {
    setSelectedAwsAccount(awsAccount.aws_account[0]);
  }

  return (
    <>
      <Typography variant="h5">AWS Integration</Typography>
      <Paper className={classes.paper}>
        <form noValidate autoComplete="off" onSubmit={onSubmit}>
          <Grid container alignItems="flex-start" spacing={2}>
            <Grid item xs={12}>
              <TextField
                id="aws-account-id"
                label="AWS Account ID"
                name="awsAccountId"
                defaultValue={selectedAwsAccount.account_id}
                error={errors.awsAccountId !== undefined}
                helperText={
                  errors.awsAccountId &&
                  'The name may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen. '
                }
                inputRef={register({
                  required: true,
                  pattern: /^[1-9]+$/i,
                })}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                id="aws-role-name"
                label="AWS Role Name"
                name="awsRoleName"
                defaultValue={selectedAwsAccount.role_name}
                error={errors.awsRoleName !== undefined}
                helperText={
                  errors.awsRoleName &&
                  'The name may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen. '
                }
                inputRef={register({
                  required: true,
                  pattern: /^[^-][A-Za-z0-9-]+[^-]$/i,
                })}
              />
            </Grid>
            <Grid item xs={12} className={classes.idGenerator}>
              <IdGeneratorTextField
                className={classes.idGeneratorTextField}
                id="aws-external-id"
                label="AWS External ID"
                name="awsExternalId"
                defaultValue={selectedAwsAccount.external_id}
                InputLabelProps={{ shrink: true }}
                disabled={true}
                error={errors.awsExternalId !== undefined}
                helperText={
                  errors.awsExternalId &&
                  'The name may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen. '
                }
                inputRef={register({
                  required: true,
                })}
              />
              <StyledButton
                variant="contained"
                className={classes.idGeneratorButton}
                onClick={handleGenerateId}
              >
                Generate ID
              </StyledButton>
            </Grid>
            <Grid item style={{ marginTop: 16 }}>
              <Button variant="contained" color="primary" type="submit">
                {awsAccount.aws_account.length === 0 ? 'Save' : 'Update'}
              </Button>
            </Grid>
          </Grid>
        </form>
      </Paper>
    </>
  );
};
