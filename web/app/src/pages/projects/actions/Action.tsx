import React from 'react';
import styled from 'styled-components';
import { useForm } from 'react-hook-form';
import { Button, Divider, Grid, Paper, TextField, Typography } from '@material-ui/core';
import {
  useGetActionsQuery,
  GetActionsQuery,
  useCreateActionMutation,
  ActionInsertInput,
  useDeleteActionMutation,
} from '../../../graphql/generated';
import { ProjectContext } from '../ProjectPage';
import { Settings } from './Settings';

const SpacedPaper = styled(Paper)`
  padding: ${(props) => props.theme.spacing(3, 2)};
`;

const Field = styled(TextField)`
  width: 100%;
`;

const Spacer = styled.div`
  height: 1em;
`;

const CancelButton = styled(Button)`
  margin-left: 5px;
`;

const DeleteButton = styled(Button)`
  margin-left: 5px;
`;

export type Action = {
  id: bigint | undefined;
  name: string;
  body: string;
};

type OptionalId = bigint | undefined;

export type ActionContextProps = {
  data: GetActionsQuery | undefined;
  loading: boolean;
  action: Action;
  setAction: (action: Action) => void;
  editing: boolean;
  setEdit: (editing: boolean) => void;
  editingId: bigint;
  setEditingId: (id: OptionalId) => void;
};

const emptyAction: Action = { id: undefined, name: '', body: '' };

export const ActionContext = React.createContext<Partial<ActionContextProps>>({});

export const Action = () => {
  const { project } = React.useContext(ProjectContext);

  const { data, loading, refetch } = useGetActionsQuery({
    variables: { projectId: project?.id },
  });

  const [action, setStateAction] = React.useState<Action>({
    id: undefined,
    name: '',
    body: '',
  });
  const [editing, setEdit] = React.useState(false);
  const [editingId, setEditingId] = React.useState<OptionalId>(undefined);

  const {
    register,
    setValue,
    handleSubmit,
    getValues,
    errors,
    clearError,
    reset: resetForm,
    setError,
  } = useForm<Action>();

  const isObject = (x: unknown): x is object =>
    x !== null && (typeof x === 'object' || typeof x === 'function');

  const setAction = (action: Action) => {
    setStateAction(action);
    setValue('name', action.name);

    if (isObject(action.body)) {
      setValue('body', JSON.stringify(action.body, null, 2));
    } else {
      setValue('body', action.body);
    }

    clearError();
  };

  const [createAction] = useCreateActionMutation();
  const [deleteAction] = useDeleteActionMutation();

  const onSubmit = handleSubmit((input: ActionInsertInput) => {
    const params = { id: editingId, project_id: project?.id, ...input };

    try {
      params.body = JSON.parse(input.body);
      createAction({ variables: { input: params } }).then(() => reset());
    } catch (e) {
      console.log(e);
      setError('body', 'syntax error', e.message);
    }
  });

  const handleOnChange = () => {
    setAction(getValues() as Action);
  };

  const onCancel = () => {
    reset();
  };

  const onDelete = () => {
    deleteAction({ variables: { id: editingId } }).then(() => reset());
  };

  const reset = () => {
    resetForm();
    setAction(emptyAction);
    setEdit(false);
    setEditingId(undefined);
    refetch();
  };

  return (
    <>
      <Typography variant="h4">Actions</Typography>
      <Typography variant="caption" display="block" gutterBottom>
        Actions.
      </Typography>

      <Divider variant="middle" />

      <Spacer />

      <Typography variant="h5">Actions</Typography>

      <ActionContext.Provider
        value={{
          data,
          loading,
          action,
          setAction,
          editing,
          setEdit,
          setEditingId,
        }}
      >
        {data?.action.length === 0 && <p>you have no actions yet.</p>}
        {data?.action.length !== 0 && <Settings />}
      </ActionContext.Provider>

      <Spacer />

      <Typography variant="h5">Add a new action</Typography>
      <SpacedPaper>
        <form noValidate autoComplete="off" onSubmit={onSubmit}>
          <Grid container alignItems="flex-start" spacing={2}>
            <Grid item xs={12}>
              <Field
                id="name"
                label="name"
                variant="outlined"
                name="name"
                size="small"
                error={errors.name !== undefined}
                helperText={
                  errors.name &&
                  'The name may only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen. '
                }
                InputLabelProps={{ shrink: action?.name !== '' }}
                onChange={handleOnChange}
                inputRef={register({
                  required: true,
                  pattern: /^[^-][A-Za-z0-9-_]+[^-]$/i,
                })}
              ></Field>
            </Grid>
            <Grid item xs={12}>
              <Field
                multiline
                rows="12"
                rowsMax="30"
                id="body"
                label="body"
                variant="outlined"
                name="body"
                size="small"
                error={errors.body !== undefined}
                helperText={
                  errors.body && errors.body.message
                  // 'The name may only contain alphanumeric characters or single hyphens or underscore, and cannot begin or end with a hyphen. '
                }
                InputLabelProps={{ shrink: action?.body !== '' }}
                onChange={handleOnChange}
                inputRef={register({
                  required: true,
                })}
              ></Field>
            </Grid>
            <Grid item style={{ marginTop: 16 }}>
              {editing && (
                <>
                  <Button variant="contained" color="primary" type="submit">
                    save
                  </Button>
                  <CancelButton variant="contained" onClick={onCancel}>
                    cancel
                  </CancelButton>
                  <DeleteButton variant="contained" color="secondary" onClick={onDelete}>
                    delete
                  </DeleteButton>
                </>
              )}
              {!editing && (
                <Button variant="contained" color="primary" type="submit">
                  add
                </Button>
              )}
            </Grid>
          </Grid>
        </form>
      </SpacedPaper>
    </>
  );
};
