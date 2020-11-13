import React from 'react';
import styled from 'styled-components';
import { useForm } from 'react-hook-form';
import { Button, Divider, Grid, Paper, TextField, Typography } from '@material-ui/core';
import {
  useGetSlackWebHookSettingsQuery,
  GetSlackWebHookSettingsQuery,
} from '../../../../graphql/generated';
import {
  useCreateSlackWebhookMutation,
  SlackWebhookInsertInput,
} from '../../../../graphql/generated';
import { ProjectContext } from '../../ProjectPage';
import { Settings } from './Settings';
import { useDeleteSlackWebHookMutation } from '../../../../graphql/generated';

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

export type Setting = {
  id: bigint | undefined;
  name: string;
  webhook_url: string;
  channel: string;
};

type OptionalId = bigint | undefined;

export type SlackSettingContextProps = {
  data: GetSlackWebHookSettingsQuery | undefined;
  loading: boolean;
  setting: Setting;
  setSetting: (setting: Setting) => void;
  editing: boolean;
  setEdit: (editing: boolean) => void;
  editingId: bigint;
  setEditingId: (id: OptionalId) => void;
};

const emptySetting: Setting = { id: undefined, name: '', webhook_url: '', channel: '' };

export const SlackSettingContext = React.createContext<Partial<SlackSettingContextProps>>({});

export const Slack = () => {
  const { project } = React.useContext(ProjectContext);

  const { data, loading, refetch } = useGetSlackWebHookSettingsQuery({
    variables: { projectId: project?.id },
  });

  const [setting, setStateSetting] = React.useState<Setting>({
    id: undefined,
    name: '',
    webhook_url: '',
    channel: '',
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
  } = useForm<SlackWebhookInsertInput>();

  const setSetting = (setting: Setting) => {
    setStateSetting(setting);
    setValue('name', setting.name);
    setValue('webhook_url', setting.webhook_url);
    setValue('channel', setting.channel);
    clearError();
  };

  const [createSlackWebhook] = useCreateSlackWebhookMutation();
  const [deleteSlackWebhook] = useDeleteSlackWebHookMutation();

  const onSubmit = handleSubmit((input: SlackWebhookInsertInput) => {
    const params = { id: editingId, project_id: project?.id, ...input };
    createSlackWebhook({ variables: { input: params } }).then(() => reset());
  });

  const handleOnChange = () => {
    setSetting(getValues() as Setting);
  };

  const onCancel = () => {
    reset();
  };

  const onDelete = () => {
    deleteSlackWebhook({ variables: { id: editingId } }).then(() => reset());
  };

  const reset = () => {
    resetForm();
    setSetting(emptySetting);
    setEdit(false);
    setEditingId(undefined);
    refetch();
  };

  return (
    <>
      <Typography variant="h4">Slack webhook</Typography>
      <Typography variant="caption" display="block" gutterBottom>
        Settings of slack webhook.
      </Typography>

      <Divider variant="middle" />

      <Spacer />

      <Typography variant="h5">Settings</Typography>

      <SlackSettingContext.Provider
        value={{
          data,
          loading,
          setting,
          setSetting,
          editing,
          setEdit,
          setEditingId,
        }}
      >
        {data?.slack_webhook.length === 0 && <p>you have no settings yet.</p>}
        {data?.slack_webhook.length !== 0 && <Settings />}
      </SlackSettingContext.Provider>

      <Spacer />

      <Typography variant="h5">Add a new setting</Typography>
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
                InputLabelProps={{ shrink: setting?.name !== '' }}
                onChange={handleOnChange}
                inputRef={register({
                  required: true,
                  pattern: /^[^-][A-Za-z0-9-_]+[^-]$/i,
                })}
              ></Field>
            </Grid>
            <Grid item xs={12}>
              <Field
                id="webhook_url"
                label="webhook url"
                variant="outlined"
                name="webhook_url"
                size="small"
                error={errors.webhook_url !== undefined}
                helperText={errors.webhook_url && 'The webhook url must contain url'}
                InputLabelProps={{ shrink: setting?.webhook_url !== '' }}
                onChange={handleOnChange}
                inputRef={register({
                  required: true,
                  pattern: /^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_+.~#?&//=]*)$/i,
                })}
              ></Field>
            </Grid>
            <Grid item xs={12}>
              <Field
                id="channel"
                label="channel"
                variant="outlined"
                name="channel"
                size="small"
                error={errors.channel !== undefined}
                helperText={
                  errors.channel &&
                  'The name may only contain alphanumeric characters or single hyphens or underscore, and cannot begin or end with a hyphen. '
                }
                InputLabelProps={{ shrink: setting?.channel !== '' }}
                onChange={handleOnChange}
                inputRef={register({
                  required: true,
                  pattern: /^[^-][A-Za-z0-9-_]+[^-]$/i,
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
