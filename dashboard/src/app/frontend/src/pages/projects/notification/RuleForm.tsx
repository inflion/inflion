import React from 'react';
import styled from 'styled-components';
import { useForm } from 'react-hook-form';
import { TextField, MenuItem } from '@material-ui/core';
import { Rule } from './Notification';

type RuleFormProps = {
  index: number;
  rule: Rule;
  notifyToParent: (rule: Rule) => void;
};

const Container = styled.div`
  padding: 0.2em;
`;

const Field = styled(TextField)`
  width: 150px;
  padding: 0.2em;
`;

const SelectEventType = styled(Field)`
  width: 180px;
`;

type FormData = {
  event_type: string;
  target_attr: string;
  match_type: string;
  match_value: string;
};

const ruleOptions = [
  {
    name: 'CPUUtilization',
    attributes: ['Message', 'InstanceId', 'Value'],
  },
  {
    name: 'SecurityGroup',
    attributes: ['Message', 'SecurityGroupId', 'SecurityGroupName', 'OpenPorts'],
  },
];

const getAttributesByEventType = (eventType: string): string[] => {
  const found = ruleOptions.find((rule) => rule.name === eventType);

  if (found) {
    return found.attributes;
  } else {
    return [];
  }
};

export const RuleForm: React.FC<RuleFormProps> = ({ index, notifyToParent, rule, children }) => {
  const [eventType, setEventType] = React.useState(rule.event_type);
  const [attributes, setAttributes] = React.useState(getAttributesByEventType(rule.event_type));
  const [selectedAttr, setSelectedAttr] = React.useState(rule.target_attr);
  const [matchType, setMatchType] = React.useState(rule.match_type);
  const [matchValue, setMatchValue] = React.useState(rule.match_value);

  const { register, setValue, errors } = useForm<FormData>();

  const change = () => {
    notifyToParent({
      index: index,
      event_type: eventType,
      target_attr: selectedAttr,
      match_type: matchType,
      match_value: matchValue,
    });
  };

  const handleEventTypeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    setEventType(value);
    setValue('event_type', value);
    const rule = ruleOptions.find((rule) => rule.name === value);
    if (rule) {
      setAttributes(rule.attributes);
      setSelectedAttr(rule.attributes[0]);
      change();
    }
  };

  const handleChangeAttr = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    setSelectedAttr(value);
    setValue('target_attr', value);
    change();
  };

  const handleMatchTypeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    setMatchType(value);
    setValue('match_type', value);
    change();
  };

  const handleMatchValueChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = event.target.value;
    setMatchValue(value);
    setValue('match_value', value);
    change();
  };

  return (
    <Container>
      <SelectEventType
        id={`event_type_${index}`}
        select
        label="event type"
        variant="outlined"
        name="event_type"
        size="small"
        value={eventType}
        onChange={handleEventTypeChange}
        helperText={errors.event_type && 'Event type is invalid'}
      >
        {ruleOptions.map((option) => (
          <MenuItem key={option.name} value={option.name}>
            {option.name}
          </MenuItem>
        ))}
      </SelectEventType>
      <Field
        id={`target_attr_${index}`}
        select
        label="target attr"
        variant="outlined"
        name="target_attr"
        size="small"
        value={selectedAttr}
        onChange={handleChangeAttr}
        error={errors.target_attr !== undefined}
        helperText={errors.target_attr && 'Target attribute is invalid'}
      >
        {attributes.map((attr) => (
          <MenuItem key={attr} value={attr}>
            {attr}
          </MenuItem>
        ))}
      </Field>

      <Field
        id={`match_type_${index}`}
        select
        label="match type"
        variant="outlined"
        name="match_type"
        size="small"
        value={matchType}
        onChange={handleMatchTypeChange}
        error={errors.match_type !== undefined}
        helperText={errors.match_type && 'Match type is invalid'}
      >
        {['contains', 'exact'].map((type) => (
          <MenuItem key={type} value={type}>
            {type}
          </MenuItem>
        ))}
      </Field>

      <Field
        id={`match_value_${index}`}
        label="match value"
        variant="outlined"
        name="match_value"
        size="small"
        value={matchValue}
        error={errors.match_value !== undefined}
        helperText={errors.match_value && 'Match value is invalid'}
        onChange={handleMatchValueChange}
        inputRef={register({
          required: true,
          pattern: /^[\w\s]+$/,
        })}
      />

      {children}
    </Container>
  );
};
