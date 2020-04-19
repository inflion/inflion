import { useMutation } from '@apollo/react-hooks';
import React from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { gql } from 'apollo-boost';
import { useForm } from 'react-hook-form';

interface NewOrgFormProps {
  created: (createdOrganization: { name: string; displayName: string }) => void;
}

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      '& > *': {
        margin: theme.spacing(1),
        width: 200,
      },
    },
  }),
);

const CREATE_ORGANIZATION = gql`
  mutation($input: OrganizationInput!) {
    createOrganization(organization: $input) {
      name
    }
  }
`;

type FormData = {
  name: string;
  displayName: string;
};

export const NewOrgForm: React.FC<NewOrgFormProps> = props => {
  const classes = useStyles();

  const { register, handleSubmit, setValue, errors } = useForm<FormData>();

  const [createOrg] = useMutation(CREATE_ORGANIZATION);

  const onSubmit = handleSubmit(({ name, displayName }) => {
    createOrg({
      variables: { input: { name: name, displayName: displayName } },
    });

    props.created({ name: name, displayName: displayName });

    setValue('name', '');
    setValue('displayName', '');
  });

  return (
    <form className={classes.root} noValidate autoComplete="off" onSubmit={onSubmit}>
      <TextField
        id="standard-basic"
        label="Name"
        name="name"
        error={errors.name !== undefined}
        helperText={
          errors.name &&
          'The name may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen. '
        }
        inputRef={register({
          required: true,
          pattern: /^[^-][A-Za-z-]+[^-]$/i,
        })}
      />
      <TextField
        id="standard-basic"
        label="Display Name"
        name="displayName"
        error={errors.displayName !== undefined}
        helperText={errors.displayName && 'This field is required.'}
        inputRef={register({ required: true })}
      />
      <Button color="primary" variant="contained" type="submit">
        Add
      </Button>
    </form>
  );
};
