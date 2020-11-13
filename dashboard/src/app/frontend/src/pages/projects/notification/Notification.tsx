import React, { useState } from 'react';

import { makeStyles, Theme } from '@material-ui/core/styles';
import styled from 'styled-components';

import { useForm } from 'react-hook-form';

import {
  useCreateNotificationRuleMutation,
  useNotificationRulesQuery,
  useNotificationRuleLazyQuery,
} from '../../../graphql';
import { useDeleteNotificationRuleMutation } from '../../../graphql/generated';

import { RuleForms } from './RuleForms';
import { ProjectContext } from '../ProjectPage';
import { Typography } from '@material-ui/core';

import {
  TextField,
  Button,
  Grid,
  Paper,
  TableContainer,
  Table,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  Divider,
} from '@material-ui/core';

export type Rule = {
  index: number;
  event_type: string;
  target_attr: string;
  match_type: string;
  match_value: string;
};

const ActionButton = styled(Button)`
  margin: 0.2em;
`;

const Container = styled.div`
  padding: 0.2em;
`;

const Field = styled(TextField)`
  width: 150px;
  padding: 0.2em;
`;

const useStyles = makeStyles((theme: Theme) => ({
  root: {
    display: 'flex',
  },
  paper: {
    padding: theme.spacing(3, 2),
  },
}));

type FormData = {
  rule_name: string | undefined;
};

export const createEmptyRule = (index: number) => {
  return { index: index } as Rule;
};

export const Notification = () => {
  const { project } = React.useContext(ProjectContext);

  const classes = useStyles();

  const [editMode, setEditMode] = useState<boolean>(false);
  const [inEditMode, setInEditMode] = useState<boolean>(false);
  const [ruleId, setRuleId] = useState<number>(0);

  const [rules, setRules] = useState<Rule[]>([createEmptyRule(0)]);

  const { register, handleSubmit, setValue, errors } = useForm<FormData>();

  const [createNotificationRule] = useCreateNotificationRuleMutation();
  const [deleteNotificationRule] = useDeleteNotificationRuleMutation();

  const { data, loading, refetch } = useNotificationRulesQuery({
    variables: { projectId: project?.id },
  });

  const [
    fetchNotificationRule,
    { data: notificationRule, loading: ruleLoading },
  ] = useNotificationRuleLazyQuery();

  const handleChange = (rules: Rule[]) => {
    setRules(rules);
  };

  const onSubmit = handleSubmit(({ rule_name }) => {
    let input = {
      id: undefined,
      project_id: project?.id,
      rule_name: rule_name,
      rules: rules,
    };

    if (inEditMode) {
      input.id = notificationRule?.notification_rule[0].id;
    }

    createNotificationRule({ variables: { input } }).then(() => {
      reset();
      refetch();
    });
  });

  const reset = () => {
    setValue('rule_name', '');
    setRules([createEmptyRule(0)]);
    setEditMode(false);
    setInEditMode(false);
  };

  const handleEdit = (id: number) => {
    setEditMode(true);
    fetchNotificationRule({ variables: { ruleId: id } });
  };

  const handleEditCancel = () => {
    reset();
  };

  const handleRemove = () => {
    deleteNotificationRule({ variables: { ruleId: ruleId } }).then(() => {
      reset();
      refetch();
    });
  };

  if (loading || !data) {
    return <p>loading</p>;
  }

  if (editMode && !ruleLoading && !inEditMode) {
    const rule = notificationRule?.notification_rule[0];
    setValue('rule_name', rule?.rule_name);
    setRules(rule?.rules as Rule[]);
    setRuleId(rule?.id);
    setInEditMode(true);
  }

  return (
    <>
      <Typography variant="h5">Notification Rules</Typography>
      <TableContainer component={Paper}>
        <Table aria-label="table">
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell align="right"></TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.notification_rule.map((row) => (
              <TableRow key={row.id}>
                <TableCell component="th" scope="row">
                  {row.rule_name}
                </TableCell>
                <TableCell align="right">
                  <Button disabled={editMode} onClick={() => handleEdit(row.id)}>
                    edit
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      <Divider />
      {editMode && <Typography variant="h5">Save the rule</Typography>}
      {!editMode && <Typography variant="h5">Add a new rule</Typography>}
      <Paper className={classes.paper}>
        <form noValidate autoComplete="off" onSubmit={onSubmit}>
          <Grid container alignItems="flex-start" spacing={2}>
            <Grid item xs={12}>
              <Container>
                <Field
                  id="rule_name"
                  label="rule name"
                  name="rule_name"
                  variant="outlined"
                  size="small"
                  InputLabelProps={{ shrink: true }}
                  error={errors.rule_name !== undefined}
                  helperText={errors.rule_name && 'Rule name is invalid'}
                  inputRef={register({
                    required: true,
                    pattern: /^[\s\w]+$/,
                  })}
                />
              </Container>
            </Grid>
            <Grid item xs={12}>
              <RuleForms ruleId={ruleId} rules={rules} handleRulesChange={handleChange}></RuleForms>
            </Grid>
            <Grid item style={{ marginTop: 16 }}>
              <Container>
                {editMode && (
                  <>
                    <ActionButton variant="contained" color="primary" type="submit">
                      save
                    </ActionButton>
                    <ActionButton variant="contained" onClick={handleEditCancel}>
                      cancel
                    </ActionButton>
                    <ActionButton variant="contained" color="secondary" onClick={handleRemove}>
                      delete
                    </ActionButton>
                  </>
                )}
                {!editMode && (
                  <ActionButton variant="contained" color="primary" type="submit">
                    add
                  </ActionButton>
                )}
              </Container>
            </Grid>
          </Grid>
        </form>
      </Paper>
    </>
  );
};
