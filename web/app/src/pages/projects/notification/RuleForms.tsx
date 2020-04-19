import React, { Fragment } from 'react';
import { RuleForm } from './RuleForm';
import { Button } from '@material-ui/core';
import { createEmptyRule, Rule } from './Notification';

type RuleFormsProps = {
  ruleId: number;
  rules: Rule[];
  handleRulesChange: (rules: Rule[]) => void;
};

export const RuleForms: React.FC<RuleFormsProps> = ({ ruleId, rules, handleRulesChange }) => {
  const handleAdd = () => handleRulesChange([...rules, createEmptyRule(rules.length)]);

  const handleRemove = (index: number) => {
    return () => {
      handleRulesChange([...rules.filter((_, i) => i !== index)]);
    };
  };

  const handleChange = (rule: Rule) => {
    rules[rule.index] = rule;
    handleRulesChange(rules);
  };

  return (
    <>
      {rules.map((rule, i) => (
        <Fragment key={`${ruleId}-${rule.index}-${i}`}>
          <RuleForm index={rule.index} notifyToParent={handleChange} rule={rule}>
            {i !== 0 ? <Button onClick={handleRemove(i)}>-</Button> : undefined}
            {i === rules.length - 1 ? <Button onClick={handleAdd}>+</Button> : undefined}
          </RuleForm>
        </Fragment>
      ))}
    </>
  );
};
