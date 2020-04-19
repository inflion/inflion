import { useMutation } from '@apollo/react-hooks';
import React, { useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import { gql } from 'apollo-boost';
import { useForm } from 'react-hook-form';
import { useAuth0 } from '../utils/react-auth0-spa';

interface UserRegistrationFormProps {
  created?: (createdUser: { username: string; email: string }) => void;
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

const CREATE_USER = gql`
  mutation($input: UserAccountCreationInput!) {
    createUserAccount(userAccount: $input) {
      username
      email
      sub
    }
  }
`;

type FormData = {
  username: string;
  email: string;
};

type User = { sub: string; email: string; nickname: string };
const judgeUserType = (user: any): user is User =>
  typeof user.sub === 'string' && typeof user.email === 'string' && typeof user.nickname === 'string';

export const UserRegistration: React.FC<UserRegistrationFormProps> = props => {
  const classes = useStyles();

  const [sub, setSub] = useState<string>('');
  const [redirect, setRedirect] = useState<string>('');

  const { register, handleSubmit, setValue, errors } = useForm<FormData>();

  const [createUser, { loading, error }] = useMutation(CREATE_USER);

  const onSubmit = handleSubmit(({ username, email }) => {
    createUser({
      variables: { input: { username, email, sub } },
    }).catch(reason => console.log(reason));

    if (props.created !== undefined) {
      props.created({ username, email });
    }

    setValue('username', '');
    setValue('email', '');
  });

  const { user } = useAuth0();

  useEffect(() => {
    if (!judgeUserType(user)) return;

    setSub(user.sub);
    setValue('username', user.nickname);
    setValue('email', user.email);

    if (loading) {
      setRedirect('/');
    }
  }, [user, loading, error, setValue]);

  if (redirect !== '') {
    return <Redirect to={redirect} />;
  }

  return (
    <form className={classes.root} noValidate autoComplete="off" onSubmit={onSubmit}>
      <TextField
        id="standard-basic"
        label="UserName"
        name="username"
        error={errors.username !== undefined}
        helperText={
          errors.username &&
          'The name may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen. '
        }
        inputRef={register({
          required: true,
          pattern: /^[^-][A-Za-z-]+[^-]$/i,
        })}
      />
      <TextField
        id="standard-basic"
        label="Email"
        name="email"
        error={errors.email !== undefined}
        helperText={errors.email && "This field mustn't be empty"}
        inputRef={register({ required: true })}
      />
      <Button color="primary" variant="contained" type="submit">
        save
      </Button>
    </form>
  );
};
