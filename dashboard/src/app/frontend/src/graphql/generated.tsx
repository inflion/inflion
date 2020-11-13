import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/react-common';
import * as ApolloReactHooks from '@apollo/react-hooks';
export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: number;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Map: any;
  Time: any;
  bigint: any;
  float8: any;
  json: any;
  jsonb: any;
  timestamp: any;
  timestamptz: any;
  uuid: any;
};

export type AwsInstance = {
   __typename?: 'AwsInstance';
  instanceId: Scalars['String'];
  name: Scalars['String'];
  privateAddress: Scalars['String'];
  publicAddress: Scalars['String'];
  securityGroupIds: Array<Scalars['String']>;
  status: Scalars['String'];
  tags: Tags;
};

/** expression to compare columns of type Boolean. All fields are combined with logical 'AND'. */
export type BooleanComparisonExp = {
  _eq?: Maybe<Scalars['Boolean']>;
  _gt?: Maybe<Scalars['Boolean']>;
  _gte?: Maybe<Scalars['Boolean']>;
  _in?: Maybe<Array<Scalars['Boolean']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['Boolean']>;
  _lte?: Maybe<Scalars['Boolean']>;
  _neq?: Maybe<Scalars['Boolean']>;
  _nin?: Maybe<Array<Scalars['Boolean']>>;
};

export type ConfirmInvitationInput = {
  token: Scalars['String'];
};

export type ConfirmInvitationOutput = {
   __typename?: 'ConfirmInvitationOutput';
  result: Scalars['Boolean'];
};


export type Query = {
   __typename?: 'Query';
  aws_instance?: Maybe<AwsInstance>;
  aws_instances: Array<AwsInstance>;
};


export type QueryAwsInstanceArgs = {
  id: Scalars['ID'];
};


export type QueryAwsInstancesArgs = {
  projectId: Scalars['ID'];
};

/** expression to compare columns of type String. All fields are combined with logical 'AND'. */
export type StringComparisonExp = {
  _eq?: Maybe<Scalars['String']>;
  _gt?: Maybe<Scalars['String']>;
  _gte?: Maybe<Scalars['String']>;
  _ilike?: Maybe<Scalars['String']>;
  _in?: Maybe<Array<Scalars['String']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _like?: Maybe<Scalars['String']>;
  _lt?: Maybe<Scalars['String']>;
  _lte?: Maybe<Scalars['String']>;
  _neq?: Maybe<Scalars['String']>;
  _nilike?: Maybe<Scalars['String']>;
  _nin?: Maybe<Array<Scalars['String']>>;
  _nlike?: Maybe<Scalars['String']>;
  _nsimilar?: Maybe<Scalars['String']>;
  _similar?: Maybe<Scalars['String']>;
};

export type Tag = {
   __typename?: 'Tag';
  key: Scalars['String'];
  value: Scalars['String'];
};

export type Tags = {
   __typename?: 'Tags';
  tags: Array<Maybe<Tag>>;
};


/** columns and relationships of "action" */
export type Action = {
   __typename?: 'action';
  body: Scalars['jsonb'];
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  name: Scalars['String'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  updated_at: Scalars['timestamp'];
  user_id: Scalars['String'];
};


/** columns and relationships of "action" */
export type ActionBodyArgs = {
  path?: Maybe<Scalars['String']>;
};

/** aggregated selection of "action" */
export type ActionAggregate = {
   __typename?: 'action_aggregate';
  aggregate?: Maybe<ActionAggregateFields>;
  nodes: Array<Action>;
};

/** aggregate fields of "action" */
export type ActionAggregateFields = {
   __typename?: 'action_aggregate_fields';
  avg?: Maybe<ActionAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<ActionMaxFields>;
  min?: Maybe<ActionMinFields>;
  stddev?: Maybe<ActionStddevFields>;
  stddev_pop?: Maybe<ActionStddevPopFields>;
  stddev_samp?: Maybe<ActionStddevSampFields>;
  sum?: Maybe<ActionSumFields>;
  var_pop?: Maybe<ActionVarPopFields>;
  var_samp?: Maybe<ActionVarSampFields>;
  variance?: Maybe<ActionVarianceFields>;
};


/** aggregate fields of "action" */
export type ActionAggregateFieldsCountArgs = {
  columns?: Maybe<Array<ActionSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "action" */
export type ActionAggregateOrderBy = {
  avg?: Maybe<ActionAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<ActionMaxOrderBy>;
  min?: Maybe<ActionMinOrderBy>;
  stddev?: Maybe<ActionStddevOrderBy>;
  stddev_pop?: Maybe<ActionStddevPopOrderBy>;
  stddev_samp?: Maybe<ActionStddevSampOrderBy>;
  sum?: Maybe<ActionSumOrderBy>;
  var_pop?: Maybe<ActionVarPopOrderBy>;
  var_samp?: Maybe<ActionVarSampOrderBy>;
  variance?: Maybe<ActionVarianceOrderBy>;
};

/** append existing jsonb value of filtered columns with new jsonb value */
export type ActionAppendInput = {
  body?: Maybe<Scalars['jsonb']>;
};

/** input type for inserting array relation for remote table "action" */
export type ActionArrRelInsertInput = {
  data: Array<ActionInsertInput>;
  on_conflict?: Maybe<ActionOnConflict>;
};

/** aggregate avg on columns */
export type ActionAvgFields = {
   __typename?: 'action_avg_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "action" */
export type ActionAvgOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "action". All fields are combined with a logical 'AND'. */
export type ActionBoolExp = {
  _and?: Maybe<Array<Maybe<ActionBoolExp>>>;
  _not?: Maybe<ActionBoolExp>;
  _or?: Maybe<Array<Maybe<ActionBoolExp>>>;
  body?: Maybe<JsonbComparisonExp>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  name?: Maybe<StringComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
  user_id?: Maybe<StringComparisonExp>;
};

/** unique or primary key constraints on table "action" */
export enum ActionConstraint {
  /** unique or primary key constraint */
  ACTION_PKEY = 'action_pkey'
}

/** delete the field or element with specified path (for JSON arrays, negative integers count from the end) */
export type ActionDeleteAtPathInput = {
  body?: Maybe<Array<Maybe<Scalars['String']>>>;
};

/** delete the array element with specified index (negative integers count from the end). throws an error if top level container is not an array */
export type ActionDeleteElemInput = {
  body?: Maybe<Scalars['Int']>;
};

/** delete key/value pair or string element. key/value pairs are matched based on their key value */
export type ActionDeleteKeyInput = {
  body?: Maybe<Scalars['String']>;
};

/** input type for incrementing integer columne in table "action" */
export type ActionIncInput = {
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "action" */
export type ActionInsertInput = {
  body?: Maybe<Scalars['jsonb']>;
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type ActionMaxFields = {
   __typename?: 'action_max_fields';
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "action" */
export type ActionMaxOrderBy = {
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type ActionMinFields = {
   __typename?: 'action_min_fields';
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "action" */
export type ActionMinOrderBy = {
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** response of any mutation on the table "action" */
export type ActionMutationResponse = {
   __typename?: 'action_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Action>;
};

/** input type for inserting object relation for remote table "action" */
export type ActionObjRelInsertInput = {
  data: ActionInsertInput;
  on_conflict?: Maybe<ActionOnConflict>;
};

/** on conflict condition type for table "action" */
export type ActionOnConflict = {
  constraint: ActionConstraint;
  update_columns: Array<ActionUpdateColumn>;
  where?: Maybe<ActionBoolExp>;
};

/** ordering options when selecting data from "action" */
export type ActionOrderBy = {
  body?: Maybe<OrderBy>;
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** primary key columns input for table: "action" */
export type ActionPkColumnsInput = {
  id: Scalars['bigint'];
};

/** prepend existing jsonb value of filtered columns with new jsonb value */
export type ActionPrependInput = {
  body?: Maybe<Scalars['jsonb']>;
};

/** select columns of table "action" */
export enum ActionSelectColumn {
  /** column name */
  BODY = 'body',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id'
}

/** input type for updating data in table "action" */
export type ActionSetInput = {
  body?: Maybe<Scalars['jsonb']>;
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type ActionStddevFields = {
   __typename?: 'action_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "action" */
export type ActionStddevOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type ActionStddevPopFields = {
   __typename?: 'action_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "action" */
export type ActionStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type ActionStddevSampFields = {
   __typename?: 'action_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "action" */
export type ActionStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type ActionSumFields = {
   __typename?: 'action_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "action" */
export type ActionSumOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "action" */
export enum ActionUpdateColumn {
  /** column name */
  BODY = 'body',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id'
}

/** aggregate var_pop on columns */
export type ActionVarPopFields = {
   __typename?: 'action_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "action" */
export type ActionVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type ActionVarSampFields = {
   __typename?: 'action_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "action" */
export type ActionVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type ActionVarianceFields = {
   __typename?: 'action_variance_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "action" */
export type ActionVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** columns and relationships of "aws_account" */
export type AwsAccount = {
   __typename?: 'aws_account';
  account_id: Scalars['String'];
  created_at: Scalars['timestamp'];
  external_id: Scalars['String'];
  id: Scalars['bigint'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  role_name: Scalars['String'];
  updated_at: Scalars['timestamp'];
  /** An object relationship */
  user_account: UserAccount;
  user_id: Scalars['String'];
};

/** aggregated selection of "aws_account" */
export type AwsAccountAggregate = {
   __typename?: 'aws_account_aggregate';
  aggregate?: Maybe<AwsAccountAggregateFields>;
  nodes: Array<AwsAccount>;
};

/** aggregate fields of "aws_account" */
export type AwsAccountAggregateFields = {
   __typename?: 'aws_account_aggregate_fields';
  avg?: Maybe<AwsAccountAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<AwsAccountMaxFields>;
  min?: Maybe<AwsAccountMinFields>;
  stddev?: Maybe<AwsAccountStddevFields>;
  stddev_pop?: Maybe<AwsAccountStddevPopFields>;
  stddev_samp?: Maybe<AwsAccountStddevSampFields>;
  sum?: Maybe<AwsAccountSumFields>;
  var_pop?: Maybe<AwsAccountVarPopFields>;
  var_samp?: Maybe<AwsAccountVarSampFields>;
  variance?: Maybe<AwsAccountVarianceFields>;
};


/** aggregate fields of "aws_account" */
export type AwsAccountAggregateFieldsCountArgs = {
  columns?: Maybe<Array<AwsAccountSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "aws_account" */
export type AwsAccountAggregateOrderBy = {
  avg?: Maybe<AwsAccountAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<AwsAccountMaxOrderBy>;
  min?: Maybe<AwsAccountMinOrderBy>;
  stddev?: Maybe<AwsAccountStddevOrderBy>;
  stddev_pop?: Maybe<AwsAccountStddevPopOrderBy>;
  stddev_samp?: Maybe<AwsAccountStddevSampOrderBy>;
  sum?: Maybe<AwsAccountSumOrderBy>;
  var_pop?: Maybe<AwsAccountVarPopOrderBy>;
  var_samp?: Maybe<AwsAccountVarSampOrderBy>;
  variance?: Maybe<AwsAccountVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "aws_account" */
export type AwsAccountArrRelInsertInput = {
  data: Array<AwsAccountInsertInput>;
  on_conflict?: Maybe<AwsAccountOnConflict>;
};

/** aggregate avg on columns */
export type AwsAccountAvgFields = {
   __typename?: 'aws_account_avg_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "aws_account" */
export type AwsAccountAvgOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "aws_account". All fields are combined with a logical 'AND'. */
export type AwsAccountBoolExp = {
  _and?: Maybe<Array<Maybe<AwsAccountBoolExp>>>;
  _not?: Maybe<AwsAccountBoolExp>;
  _or?: Maybe<Array<Maybe<AwsAccountBoolExp>>>;
  account_id?: Maybe<StringComparisonExp>;
  created_at?: Maybe<TimestampComparisonExp>;
  external_id?: Maybe<StringComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  role_name?: Maybe<StringComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
  user_account?: Maybe<UserAccountBoolExp>;
  user_id?: Maybe<StringComparisonExp>;
};

/** unique or primary key constraints on table "aws_account" */
export enum AwsAccountConstraint {
  /** unique or primary key constraint */
  AWS_ACCOUNT_PKEY = 'aws_account_pkey'
}

/** input type for incrementing integer columne in table "aws_account" */
export type AwsAccountIncInput = {
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "aws_account" */
export type AwsAccountInsertInput = {
  account_id?: Maybe<Scalars['String']>;
  created_at?: Maybe<Scalars['timestamp']>;
  external_id?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  role_name?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_account?: Maybe<UserAccountObjRelInsertInput>;
  user_id?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type AwsAccountMaxFields = {
   __typename?: 'aws_account_max_fields';
  account_id?: Maybe<Scalars['String']>;
  external_id?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
  role_name?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "aws_account" */
export type AwsAccountMaxOrderBy = {
  account_id?: Maybe<OrderBy>;
  external_id?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  role_name?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type AwsAccountMinFields = {
   __typename?: 'aws_account_min_fields';
  account_id?: Maybe<Scalars['String']>;
  external_id?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
  role_name?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "aws_account" */
export type AwsAccountMinOrderBy = {
  account_id?: Maybe<OrderBy>;
  external_id?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  role_name?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** response of any mutation on the table "aws_account" */
export type AwsAccountMutationResponse = {
   __typename?: 'aws_account_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<AwsAccount>;
};

/** input type for inserting object relation for remote table "aws_account" */
export type AwsAccountObjRelInsertInput = {
  data: AwsAccountInsertInput;
  on_conflict?: Maybe<AwsAccountOnConflict>;
};

/** on conflict condition type for table "aws_account" */
export type AwsAccountOnConflict = {
  constraint: AwsAccountConstraint;
  update_columns: Array<AwsAccountUpdateColumn>;
  where?: Maybe<AwsAccountBoolExp>;
};

/** ordering options when selecting data from "aws_account" */
export type AwsAccountOrderBy = {
  account_id?: Maybe<OrderBy>;
  created_at?: Maybe<OrderBy>;
  external_id?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  role_name?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
  user_account?: Maybe<UserAccountOrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** primary key columns input for table: "aws_account" */
export type AwsAccountPkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "aws_account" */
export enum AwsAccountSelectColumn {
  /** column name */
  ACCOUNT_ID = 'account_id',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  EXTERNAL_ID = 'external_id',
  /** column name */
  ID = 'id',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  ROLE_NAME = 'role_name',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id'
}

/** input type for updating data in table "aws_account" */
export type AwsAccountSetInput = {
  account_id?: Maybe<Scalars['String']>;
  created_at?: Maybe<Scalars['timestamp']>;
  external_id?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
  role_name?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type AwsAccountStddevFields = {
   __typename?: 'aws_account_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "aws_account" */
export type AwsAccountStddevOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type AwsAccountStddevPopFields = {
   __typename?: 'aws_account_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "aws_account" */
export type AwsAccountStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type AwsAccountStddevSampFields = {
   __typename?: 'aws_account_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "aws_account" */
export type AwsAccountStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type AwsAccountSumFields = {
   __typename?: 'aws_account_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "aws_account" */
export type AwsAccountSumOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "aws_account" */
export enum AwsAccountUpdateColumn {
  /** column name */
  ACCOUNT_ID = 'account_id',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  EXTERNAL_ID = 'external_id',
  /** column name */
  ID = 'id',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  ROLE_NAME = 'role_name',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id'
}

/** aggregate var_pop on columns */
export type AwsAccountVarPopFields = {
   __typename?: 'aws_account_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "aws_account" */
export type AwsAccountVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type AwsAccountVarSampFields = {
   __typename?: 'aws_account_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "aws_account" */
export type AwsAccountVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type AwsAccountVarianceFields = {
   __typename?: 'aws_account_variance_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "aws_account" */
export type AwsAccountVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};


/** expression to compare columns of type bigint. All fields are combined with logical 'AND'. */
export type BigintComparisonExp = {
  _eq?: Maybe<Scalars['bigint']>;
  _gt?: Maybe<Scalars['bigint']>;
  _gte?: Maybe<Scalars['bigint']>;
  _in?: Maybe<Array<Scalars['bigint']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['bigint']>;
  _lte?: Maybe<Scalars['bigint']>;
  _neq?: Maybe<Scalars['bigint']>;
  _nin?: Maybe<Array<Scalars['bigint']>>;
};


/** expression to compare columns of type float8. All fields are combined with logical 'AND'. */
export type Float8ComparisonExp = {
  _eq?: Maybe<Scalars['float8']>;
  _gt?: Maybe<Scalars['float8']>;
  _gte?: Maybe<Scalars['float8']>;
  _in?: Maybe<Array<Scalars['float8']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['float8']>;
  _lte?: Maybe<Scalars['float8']>;
  _neq?: Maybe<Scalars['float8']>;
  _nin?: Maybe<Array<Scalars['float8']>>;
};

/** columns and relationships of "instance" */
export type Instance = {
   __typename?: 'instance';
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  instance_id: Scalars['String'];
  name: Scalars['String'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  status?: Maybe<Scalars['String']>;
  updated_at: Scalars['timestamp'];
};

/** aggregated selection of "instance" */
export type InstanceAggregate = {
   __typename?: 'instance_aggregate';
  aggregate?: Maybe<InstanceAggregateFields>;
  nodes: Array<Instance>;
};

/** aggregate fields of "instance" */
export type InstanceAggregateFields = {
   __typename?: 'instance_aggregate_fields';
  avg?: Maybe<InstanceAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<InstanceMaxFields>;
  min?: Maybe<InstanceMinFields>;
  stddev?: Maybe<InstanceStddevFields>;
  stddev_pop?: Maybe<InstanceStddevPopFields>;
  stddev_samp?: Maybe<InstanceStddevSampFields>;
  sum?: Maybe<InstanceSumFields>;
  var_pop?: Maybe<InstanceVarPopFields>;
  var_samp?: Maybe<InstanceVarSampFields>;
  variance?: Maybe<InstanceVarianceFields>;
};


/** aggregate fields of "instance" */
export type InstanceAggregateFieldsCountArgs = {
  columns?: Maybe<Array<InstanceSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "instance" */
export type InstanceAggregateOrderBy = {
  avg?: Maybe<InstanceAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<InstanceMaxOrderBy>;
  min?: Maybe<InstanceMinOrderBy>;
  stddev?: Maybe<InstanceStddevOrderBy>;
  stddev_pop?: Maybe<InstanceStddevPopOrderBy>;
  stddev_samp?: Maybe<InstanceStddevSampOrderBy>;
  sum?: Maybe<InstanceSumOrderBy>;
  var_pop?: Maybe<InstanceVarPopOrderBy>;
  var_samp?: Maybe<InstanceVarSampOrderBy>;
  variance?: Maybe<InstanceVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "instance" */
export type InstanceArrRelInsertInput = {
  data: Array<InstanceInsertInput>;
  on_conflict?: Maybe<InstanceOnConflict>;
};

/** columns and relationships of "instance_at_service" */
export type InstanceAtService = {
   __typename?: 'instance_at_service';
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  instance_id: Scalars['bigint'];
  service_id: Scalars['bigint'];
  updated_at: Scalars['timestamp'];
};

/** aggregated selection of "instance_at_service" */
export type InstanceAtServiceAggregate = {
   __typename?: 'instance_at_service_aggregate';
  aggregate?: Maybe<InstanceAtServiceAggregateFields>;
  nodes: Array<InstanceAtService>;
};

/** aggregate fields of "instance_at_service" */
export type InstanceAtServiceAggregateFields = {
   __typename?: 'instance_at_service_aggregate_fields';
  avg?: Maybe<InstanceAtServiceAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<InstanceAtServiceMaxFields>;
  min?: Maybe<InstanceAtServiceMinFields>;
  stddev?: Maybe<InstanceAtServiceStddevFields>;
  stddev_pop?: Maybe<InstanceAtServiceStddevPopFields>;
  stddev_samp?: Maybe<InstanceAtServiceStddevSampFields>;
  sum?: Maybe<InstanceAtServiceSumFields>;
  var_pop?: Maybe<InstanceAtServiceVarPopFields>;
  var_samp?: Maybe<InstanceAtServiceVarSampFields>;
  variance?: Maybe<InstanceAtServiceVarianceFields>;
};


/** aggregate fields of "instance_at_service" */
export type InstanceAtServiceAggregateFieldsCountArgs = {
  columns?: Maybe<Array<InstanceAtServiceSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "instance_at_service" */
export type InstanceAtServiceAggregateOrderBy = {
  avg?: Maybe<InstanceAtServiceAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<InstanceAtServiceMaxOrderBy>;
  min?: Maybe<InstanceAtServiceMinOrderBy>;
  stddev?: Maybe<InstanceAtServiceStddevOrderBy>;
  stddev_pop?: Maybe<InstanceAtServiceStddevPopOrderBy>;
  stddev_samp?: Maybe<InstanceAtServiceStddevSampOrderBy>;
  sum?: Maybe<InstanceAtServiceSumOrderBy>;
  var_pop?: Maybe<InstanceAtServiceVarPopOrderBy>;
  var_samp?: Maybe<InstanceAtServiceVarSampOrderBy>;
  variance?: Maybe<InstanceAtServiceVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "instance_at_service" */
export type InstanceAtServiceArrRelInsertInput = {
  data: Array<InstanceAtServiceInsertInput>;
  on_conflict?: Maybe<InstanceAtServiceOnConflict>;
};

/** aggregate avg on columns */
export type InstanceAtServiceAvgFields = {
   __typename?: 'instance_at_service_avg_fields';
  id?: Maybe<Scalars['Float']>;
  instance_id?: Maybe<Scalars['Float']>;
  service_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "instance_at_service" */
export type InstanceAtServiceAvgOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "instance_at_service". All fields are combined with a logical 'AND'. */
export type InstanceAtServiceBoolExp = {
  _and?: Maybe<Array<Maybe<InstanceAtServiceBoolExp>>>;
  _not?: Maybe<InstanceAtServiceBoolExp>;
  _or?: Maybe<Array<Maybe<InstanceAtServiceBoolExp>>>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  instance_id?: Maybe<BigintComparisonExp>;
  service_id?: Maybe<BigintComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
};

/** unique or primary key constraints on table "instance_at_service" */
export enum InstanceAtServiceConstraint {
  /** unique or primary key constraint */
  INSTANCE_AT_SERVICE_PKEY = 'instance_at_service_pkey',
  /** unique or primary key constraint */
  INSTANCE_AT_SERVICE_SERVICE_ID_INSTANCE_ID_KEY = 'instance_at_service_service_id_instance_id_key'
}

/** input type for incrementing integer columne in table "instance_at_service" */
export type InstanceAtServiceIncInput = {
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['bigint']>;
  service_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "instance_at_service" */
export type InstanceAtServiceInsertInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['bigint']>;
  service_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate max on columns */
export type InstanceAtServiceMaxFields = {
   __typename?: 'instance_at_service_max_fields';
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['bigint']>;
  service_id?: Maybe<Scalars['bigint']>;
};

/** order by max() on columns of table "instance_at_service" */
export type InstanceAtServiceMaxOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type InstanceAtServiceMinFields = {
   __typename?: 'instance_at_service_min_fields';
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['bigint']>;
  service_id?: Maybe<Scalars['bigint']>;
};

/** order by min() on columns of table "instance_at_service" */
export type InstanceAtServiceMinOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** response of any mutation on the table "instance_at_service" */
export type InstanceAtServiceMutationResponse = {
   __typename?: 'instance_at_service_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<InstanceAtService>;
};

/** input type for inserting object relation for remote table "instance_at_service" */
export type InstanceAtServiceObjRelInsertInput = {
  data: InstanceAtServiceInsertInput;
  on_conflict?: Maybe<InstanceAtServiceOnConflict>;
};

/** on conflict condition type for table "instance_at_service" */
export type InstanceAtServiceOnConflict = {
  constraint: InstanceAtServiceConstraint;
  update_columns: Array<InstanceAtServiceUpdateColumn>;
  where?: Maybe<InstanceAtServiceBoolExp>;
};

/** ordering options when selecting data from "instance_at_service" */
export type InstanceAtServiceOrderBy = {
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
};

/** primary key columns input for table: "instance_at_service" */
export type InstanceAtServicePkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "instance_at_service" */
export enum InstanceAtServiceSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  INSTANCE_ID = 'instance_id',
  /** column name */
  SERVICE_ID = 'service_id',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** input type for updating data in table "instance_at_service" */
export type InstanceAtServiceSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['bigint']>;
  service_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate stddev on columns */
export type InstanceAtServiceStddevFields = {
   __typename?: 'instance_at_service_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  instance_id?: Maybe<Scalars['Float']>;
  service_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "instance_at_service" */
export type InstanceAtServiceStddevOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type InstanceAtServiceStddevPopFields = {
   __typename?: 'instance_at_service_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  instance_id?: Maybe<Scalars['Float']>;
  service_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "instance_at_service" */
export type InstanceAtServiceStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type InstanceAtServiceStddevSampFields = {
   __typename?: 'instance_at_service_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  instance_id?: Maybe<Scalars['Float']>;
  service_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "instance_at_service" */
export type InstanceAtServiceStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type InstanceAtServiceSumFields = {
   __typename?: 'instance_at_service_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['bigint']>;
  service_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "instance_at_service" */
export type InstanceAtServiceSumOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** update columns of table "instance_at_service" */
export enum InstanceAtServiceUpdateColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  INSTANCE_ID = 'instance_id',
  /** column name */
  SERVICE_ID = 'service_id',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** aggregate var_pop on columns */
export type InstanceAtServiceVarPopFields = {
   __typename?: 'instance_at_service_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  instance_id?: Maybe<Scalars['Float']>;
  service_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "instance_at_service" */
export type InstanceAtServiceVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type InstanceAtServiceVarSampFields = {
   __typename?: 'instance_at_service_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  instance_id?: Maybe<Scalars['Float']>;
  service_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "instance_at_service" */
export type InstanceAtServiceVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type InstanceAtServiceVarianceFields = {
   __typename?: 'instance_at_service_variance_fields';
  id?: Maybe<Scalars['Float']>;
  instance_id?: Maybe<Scalars['Float']>;
  service_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "instance_at_service" */
export type InstanceAtServiceVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  service_id?: Maybe<OrderBy>;
};

/** aggregate avg on columns */
export type InstanceAvgFields = {
   __typename?: 'instance_avg_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "instance" */
export type InstanceAvgOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "instance". All fields are combined with a logical 'AND'. */
export type InstanceBoolExp = {
  _and?: Maybe<Array<Maybe<InstanceBoolExp>>>;
  _not?: Maybe<InstanceBoolExp>;
  _or?: Maybe<Array<Maybe<InstanceBoolExp>>>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  instance_id?: Maybe<StringComparisonExp>;
  name?: Maybe<StringComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  status?: Maybe<StringComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
};

/** unique or primary key constraints on table "instance" */
export enum InstanceConstraint {
  /** unique or primary key constraint */
  INSTANCE_INSTANCE_ID_KEY = 'instance_instance_id_key',
  /** unique or primary key constraint */
  INSTANCE_PKEY = 'instance_pkey'
}

/** input type for incrementing integer columne in table "instance" */
export type InstanceIncInput = {
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "instance" */
export type InstanceInsertInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  status?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate max on columns */
export type InstanceMaxFields = {
   __typename?: 'instance_max_fields';
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  status?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "instance" */
export type InstanceMaxOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  status?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type InstanceMinFields = {
   __typename?: 'instance_min_fields';
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  status?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "instance" */
export type InstanceMinOrderBy = {
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  status?: Maybe<OrderBy>;
};

/** response of any mutation on the table "instance" */
export type InstanceMutationResponse = {
   __typename?: 'instance_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Instance>;
};

/** input type for inserting object relation for remote table "instance" */
export type InstanceObjRelInsertInput = {
  data: InstanceInsertInput;
  on_conflict?: Maybe<InstanceOnConflict>;
};

/** on conflict condition type for table "instance" */
export type InstanceOnConflict = {
  constraint: InstanceConstraint;
  update_columns: Array<InstanceUpdateColumn>;
  where?: Maybe<InstanceBoolExp>;
};

/** ordering options when selecting data from "instance" */
export type InstanceOrderBy = {
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  instance_id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  status?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
};

/** primary key columns input for table: "instance" */
export type InstancePkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "instance" */
export enum InstanceSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  INSTANCE_ID = 'instance_id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  STATUS = 'status',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** input type for updating data in table "instance" */
export type InstanceSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  instance_id?: Maybe<Scalars['String']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  status?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate stddev on columns */
export type InstanceStddevFields = {
   __typename?: 'instance_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "instance" */
export type InstanceStddevOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type InstanceStddevPopFields = {
   __typename?: 'instance_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "instance" */
export type InstanceStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type InstanceStddevSampFields = {
   __typename?: 'instance_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "instance" */
export type InstanceStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type InstanceSumFields = {
   __typename?: 'instance_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "instance" */
export type InstanceSumOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "instance" */
export enum InstanceUpdateColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  INSTANCE_ID = 'instance_id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  STATUS = 'status',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** aggregate var_pop on columns */
export type InstanceVarPopFields = {
   __typename?: 'instance_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "instance" */
export type InstanceVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type InstanceVarSampFields = {
   __typename?: 'instance_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "instance" */
export type InstanceVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type InstanceVarianceFields = {
   __typename?: 'instance_variance_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "instance" */
export type InstanceVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};


/** expression to compare columns of type json. All fields are combined with logical 'AND'. */
export type JsonComparisonExp = {
  _eq?: Maybe<Scalars['json']>;
  _gt?: Maybe<Scalars['json']>;
  _gte?: Maybe<Scalars['json']>;
  _in?: Maybe<Array<Scalars['json']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['json']>;
  _lte?: Maybe<Scalars['json']>;
  _neq?: Maybe<Scalars['json']>;
  _nin?: Maybe<Array<Scalars['json']>>;
};


/** expression to compare columns of type jsonb. All fields are combined with logical 'AND'. */
export type JsonbComparisonExp = {
  /** is the column contained in the given json value */
  _contained_in?: Maybe<Scalars['jsonb']>;
  /** does the column contain the given json value at the top level */
  _contains?: Maybe<Scalars['jsonb']>;
  _eq?: Maybe<Scalars['jsonb']>;
  _gt?: Maybe<Scalars['jsonb']>;
  _gte?: Maybe<Scalars['jsonb']>;
  /** does the string exist as a top-level key in the column */
  _has_key?: Maybe<Scalars['String']>;
  /** do all of these strings exist as top-level keys in the column */
  _has_keys_all?: Maybe<Array<Scalars['String']>>;
  /** do any of these strings exist as top-level keys in the column */
  _has_keys_any?: Maybe<Array<Scalars['String']>>;
  _in?: Maybe<Array<Scalars['jsonb']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['jsonb']>;
  _lte?: Maybe<Scalars['jsonb']>;
  _neq?: Maybe<Scalars['jsonb']>;
  _nin?: Maybe<Array<Scalars['jsonb']>>;
};

/** columns and relationships of "metrics" */
export type Metrics = {
   __typename?: 'metrics';
  instance_id: Scalars['String'];
  time: Scalars['timestamptz'];
  type: Scalars['String'];
  value?: Maybe<Scalars['float8']>;
};

/** aggregated selection of "metrics" */
export type MetricsAggregate = {
   __typename?: 'metrics_aggregate';
  aggregate?: Maybe<MetricsAggregateFields>;
  nodes: Array<Metrics>;
};

/** aggregate fields of "metrics" */
export type MetricsAggregateFields = {
   __typename?: 'metrics_aggregate_fields';
  avg?: Maybe<MetricsAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<MetricsMaxFields>;
  min?: Maybe<MetricsMinFields>;
  stddev?: Maybe<MetricsStddevFields>;
  stddev_pop?: Maybe<MetricsStddevPopFields>;
  stddev_samp?: Maybe<MetricsStddevSampFields>;
  sum?: Maybe<MetricsSumFields>;
  var_pop?: Maybe<MetricsVarPopFields>;
  var_samp?: Maybe<MetricsVarSampFields>;
  variance?: Maybe<MetricsVarianceFields>;
};


/** aggregate fields of "metrics" */
export type MetricsAggregateFieldsCountArgs = {
  columns?: Maybe<Array<MetricsSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "metrics" */
export type MetricsAggregateOrderBy = {
  avg?: Maybe<MetricsAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<MetricsMaxOrderBy>;
  min?: Maybe<MetricsMinOrderBy>;
  stddev?: Maybe<MetricsStddevOrderBy>;
  stddev_pop?: Maybe<MetricsStddevPopOrderBy>;
  stddev_samp?: Maybe<MetricsStddevSampOrderBy>;
  sum?: Maybe<MetricsSumOrderBy>;
  var_pop?: Maybe<MetricsVarPopOrderBy>;
  var_samp?: Maybe<MetricsVarSampOrderBy>;
  variance?: Maybe<MetricsVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "metrics" */
export type MetricsArrRelInsertInput = {
  data: Array<MetricsInsertInput>;
};

/** aggregate avg on columns */
export type MetricsAvgFields = {
   __typename?: 'metrics_avg_fields';
  value?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "metrics" */
export type MetricsAvgOrderBy = {
  value?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "metrics". All fields are combined with a logical 'AND'. */
export type MetricsBoolExp = {
  _and?: Maybe<Array<Maybe<MetricsBoolExp>>>;
  _not?: Maybe<MetricsBoolExp>;
  _or?: Maybe<Array<Maybe<MetricsBoolExp>>>;
  instance_id?: Maybe<StringComparisonExp>;
  time?: Maybe<TimestamptzComparisonExp>;
  type?: Maybe<StringComparisonExp>;
  value?: Maybe<Float8ComparisonExp>;
};

/** input type for inserting data into table "metrics" */
export type MetricsInsertInput = {
  instance_id?: Maybe<Scalars['String']>;
  time?: Maybe<Scalars['timestamptz']>;
  type?: Maybe<Scalars['String']>;
  value?: Maybe<Scalars['float8']>;
};

/** aggregate max on columns */
export type MetricsMaxFields = {
   __typename?: 'metrics_max_fields';
  instance_id?: Maybe<Scalars['String']>;
  time?: Maybe<Scalars['timestamptz']>;
  type?: Maybe<Scalars['String']>;
  value?: Maybe<Scalars['float8']>;
};

/** order by max() on columns of table "metrics" */
export type MetricsMaxOrderBy = {
  instance_id?: Maybe<OrderBy>;
  time?: Maybe<OrderBy>;
  type?: Maybe<OrderBy>;
  value?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type MetricsMinFields = {
   __typename?: 'metrics_min_fields';
  instance_id?: Maybe<Scalars['String']>;
  time?: Maybe<Scalars['timestamptz']>;
  type?: Maybe<Scalars['String']>;
  value?: Maybe<Scalars['float8']>;
};

/** order by min() on columns of table "metrics" */
export type MetricsMinOrderBy = {
  instance_id?: Maybe<OrderBy>;
  time?: Maybe<OrderBy>;
  type?: Maybe<OrderBy>;
  value?: Maybe<OrderBy>;
};

/** response of any mutation on the table "metrics" */
export type MetricsMutationResponse = {
   __typename?: 'metrics_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Metrics>;
};

/** input type for inserting object relation for remote table "metrics" */
export type MetricsObjRelInsertInput = {
  data: MetricsInsertInput;
};

/** ordering options when selecting data from "metrics" */
export type MetricsOrderBy = {
  instance_id?: Maybe<OrderBy>;
  time?: Maybe<OrderBy>;
  type?: Maybe<OrderBy>;
  value?: Maybe<OrderBy>;
};

/** select columns of table "metrics" */
export enum MetricsSelectColumn {
  /** column name */
  INSTANCE_ID = 'instance_id',
  /** column name */
  TIME = 'time',
  /** column name */
  TYPE = 'type',
  /** column name */
  VALUE = 'value'
}

/** input type for updating data in table "metrics" */
export type MetricsSetInput = {
  instance_id?: Maybe<Scalars['String']>;
  time?: Maybe<Scalars['timestamptz']>;
  type?: Maybe<Scalars['String']>;
  value?: Maybe<Scalars['float8']>;
};

/** aggregate stddev on columns */
export type MetricsStddevFields = {
   __typename?: 'metrics_stddev_fields';
  value?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "metrics" */
export type MetricsStddevOrderBy = {
  value?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type MetricsStddevPopFields = {
   __typename?: 'metrics_stddev_pop_fields';
  value?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "metrics" */
export type MetricsStddevPopOrderBy = {
  value?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type MetricsStddevSampFields = {
   __typename?: 'metrics_stddev_samp_fields';
  value?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "metrics" */
export type MetricsStddevSampOrderBy = {
  value?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type MetricsSumFields = {
   __typename?: 'metrics_sum_fields';
  value?: Maybe<Scalars['float8']>;
};

/** order by sum() on columns of table "metrics" */
export type MetricsSumOrderBy = {
  value?: Maybe<OrderBy>;
};

/** aggregate var_pop on columns */
export type MetricsVarPopFields = {
   __typename?: 'metrics_var_pop_fields';
  value?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "metrics" */
export type MetricsVarPopOrderBy = {
  value?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type MetricsVarSampFields = {
   __typename?: 'metrics_var_samp_fields';
  value?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "metrics" */
export type MetricsVarSampOrderBy = {
  value?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type MetricsVarianceFields = {
   __typename?: 'metrics_variance_fields';
  value?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "metrics" */
export type MetricsVarianceOrderBy = {
  value?: Maybe<OrderBy>;
};

/** mutation root */
export type MutationRoot = {
   __typename?: 'mutation_root';
  /** perform the action: "confirmInvitation" */
  confirmInvitation: ConfirmInvitationOutput;
  /** delete data from the table: "action" */
  delete_action?: Maybe<ActionMutationResponse>;
  /** delete single row from the table: "action" */
  delete_action_by_pk?: Maybe<Action>;
  /** delete data from the table: "aws_account" */
  delete_aws_account?: Maybe<AwsAccountMutationResponse>;
  /** delete single row from the table: "aws_account" */
  delete_aws_account_by_pk?: Maybe<AwsAccount>;
  /** delete data from the table: "instance" */
  delete_instance?: Maybe<InstanceMutationResponse>;
  /** delete data from the table: "instance_at_service" */
  delete_instance_at_service?: Maybe<InstanceAtServiceMutationResponse>;
  /** delete single row from the table: "instance_at_service" */
  delete_instance_at_service_by_pk?: Maybe<InstanceAtService>;
  /** delete single row from the table: "instance" */
  delete_instance_by_pk?: Maybe<Instance>;
  /** delete data from the table: "metrics" */
  delete_metrics?: Maybe<MetricsMutationResponse>;
  /** delete data from the table: "notification_rule" */
  delete_notification_rule?: Maybe<NotificationRuleMutationResponse>;
  /** delete single row from the table: "notification_rule" */
  delete_notification_rule_by_pk?: Maybe<NotificationRule>;
  /** delete data from the table: "organization" */
  delete_organization?: Maybe<OrganizationMutationResponse>;
  /** delete single row from the table: "organization" */
  delete_organization_by_pk?: Maybe<Organization>;
  /** delete data from the table: "project" */
  delete_project?: Maybe<ProjectMutationResponse>;
  /** delete single row from the table: "project" */
  delete_project_by_pk?: Maybe<Project>;
  /** delete data from the table: "project_collaborator" */
  delete_project_collaborator?: Maybe<ProjectCollaboratorMutationResponse>;
  /** delete data from the table: "project_in_organization" */
  delete_project_in_organization?: Maybe<ProjectInOrganizationMutationResponse>;
  /** delete single row from the table: "project_in_organization" */
  delete_project_in_organization_by_pk?: Maybe<ProjectInOrganization>;
  /** delete data from the table: "project_invitation" */
  delete_project_invitation?: Maybe<ProjectInvitationMutationResponse>;
  /** delete single row from the table: "project_invitation" */
  delete_project_invitation_by_pk?: Maybe<ProjectInvitation>;
  /** delete data from the table: "service" */
  delete_service?: Maybe<ServiceMutationResponse>;
  /** delete single row from the table: "service" */
  delete_service_by_pk?: Maybe<Service>;
  /** delete data from the table: "slack_webhook" */
  delete_slack_webhook?: Maybe<SlackWebhookMutationResponse>;
  /** delete single row from the table: "slack_webhook" */
  delete_slack_webhook_by_pk?: Maybe<SlackWebhook>;
  /** delete data from the table: "user_account" */
  delete_user_account?: Maybe<UserAccountMutationResponse>;
  /** delete single row from the table: "user_account" */
  delete_user_account_by_pk?: Maybe<UserAccount>;
  /** insert data into the table: "action" */
  insert_action?: Maybe<ActionMutationResponse>;
  /** insert a single row into the table: "action" */
  insert_action_one?: Maybe<Action>;
  /** insert data into the table: "aws_account" */
  insert_aws_account?: Maybe<AwsAccountMutationResponse>;
  /** insert a single row into the table: "aws_account" */
  insert_aws_account_one?: Maybe<AwsAccount>;
  /** insert data into the table: "instance" */
  insert_instance?: Maybe<InstanceMutationResponse>;
  /** insert data into the table: "instance_at_service" */
  insert_instance_at_service?: Maybe<InstanceAtServiceMutationResponse>;
  /** insert a single row into the table: "instance_at_service" */
  insert_instance_at_service_one?: Maybe<InstanceAtService>;
  /** insert a single row into the table: "instance" */
  insert_instance_one?: Maybe<Instance>;
  /** insert data into the table: "metrics" */
  insert_metrics?: Maybe<MetricsMutationResponse>;
  /** insert a single row into the table: "metrics" */
  insert_metrics_one?: Maybe<Metrics>;
  /** insert data into the table: "notification_rule" */
  insert_notification_rule?: Maybe<NotificationRuleMutationResponse>;
  /** insert a single row into the table: "notification_rule" */
  insert_notification_rule_one?: Maybe<NotificationRule>;
  /** insert data into the table: "organization" */
  insert_organization?: Maybe<OrganizationMutationResponse>;
  /** insert a single row into the table: "organization" */
  insert_organization_one?: Maybe<Organization>;
  /** insert data into the table: "project" */
  insert_project?: Maybe<ProjectMutationResponse>;
  /** insert data into the table: "project_collaborator" */
  insert_project_collaborator?: Maybe<ProjectCollaboratorMutationResponse>;
  /** insert a single row into the table: "project_collaborator" */
  insert_project_collaborator_one?: Maybe<ProjectCollaborator>;
  /** insert data into the table: "project_in_organization" */
  insert_project_in_organization?: Maybe<ProjectInOrganizationMutationResponse>;
  /** insert a single row into the table: "project_in_organization" */
  insert_project_in_organization_one?: Maybe<ProjectInOrganization>;
  /** insert data into the table: "project_invitation" */
  insert_project_invitation?: Maybe<ProjectInvitationMutationResponse>;
  /** insert a single row into the table: "project_invitation" */
  insert_project_invitation_one?: Maybe<ProjectInvitation>;
  /** insert a single row into the table: "project" */
  insert_project_one?: Maybe<Project>;
  /** insert data into the table: "service" */
  insert_service?: Maybe<ServiceMutationResponse>;
  /** insert a single row into the table: "service" */
  insert_service_one?: Maybe<Service>;
  /** insert data into the table: "slack_webhook" */
  insert_slack_webhook?: Maybe<SlackWebhookMutationResponse>;
  /** insert a single row into the table: "slack_webhook" */
  insert_slack_webhook_one?: Maybe<SlackWebhook>;
  /** insert data into the table: "user_account" */
  insert_user_account?: Maybe<UserAccountMutationResponse>;
  /** insert a single row into the table: "user_account" */
  insert_user_account_one?: Maybe<UserAccount>;
  /** update data of the table: "action" */
  update_action?: Maybe<ActionMutationResponse>;
  /** update single row of the table: "action" */
  update_action_by_pk?: Maybe<Action>;
  /** update data of the table: "aws_account" */
  update_aws_account?: Maybe<AwsAccountMutationResponse>;
  /** update single row of the table: "aws_account" */
  update_aws_account_by_pk?: Maybe<AwsAccount>;
  /** update data of the table: "instance" */
  update_instance?: Maybe<InstanceMutationResponse>;
  /** update data of the table: "instance_at_service" */
  update_instance_at_service?: Maybe<InstanceAtServiceMutationResponse>;
  /** update single row of the table: "instance_at_service" */
  update_instance_at_service_by_pk?: Maybe<InstanceAtService>;
  /** update single row of the table: "instance" */
  update_instance_by_pk?: Maybe<Instance>;
  /** update data of the table: "metrics" */
  update_metrics?: Maybe<MetricsMutationResponse>;
  /** update data of the table: "notification_rule" */
  update_notification_rule?: Maybe<NotificationRuleMutationResponse>;
  /** update single row of the table: "notification_rule" */
  update_notification_rule_by_pk?: Maybe<NotificationRule>;
  /** update data of the table: "organization" */
  update_organization?: Maybe<OrganizationMutationResponse>;
  /** update single row of the table: "organization" */
  update_organization_by_pk?: Maybe<Organization>;
  /** update data of the table: "project" */
  update_project?: Maybe<ProjectMutationResponse>;
  /** update single row of the table: "project" */
  update_project_by_pk?: Maybe<Project>;
  /** update data of the table: "project_collaborator" */
  update_project_collaborator?: Maybe<ProjectCollaboratorMutationResponse>;
  /** update data of the table: "project_in_organization" */
  update_project_in_organization?: Maybe<ProjectInOrganizationMutationResponse>;
  /** update single row of the table: "project_in_organization" */
  update_project_in_organization_by_pk?: Maybe<ProjectInOrganization>;
  /** update data of the table: "project_invitation" */
  update_project_invitation?: Maybe<ProjectInvitationMutationResponse>;
  /** update single row of the table: "project_invitation" */
  update_project_invitation_by_pk?: Maybe<ProjectInvitation>;
  /** update data of the table: "service" */
  update_service?: Maybe<ServiceMutationResponse>;
  /** update single row of the table: "service" */
  update_service_by_pk?: Maybe<Service>;
  /** update data of the table: "slack_webhook" */
  update_slack_webhook?: Maybe<SlackWebhookMutationResponse>;
  /** update single row of the table: "slack_webhook" */
  update_slack_webhook_by_pk?: Maybe<SlackWebhook>;
  /** update data of the table: "user_account" */
  update_user_account?: Maybe<UserAccountMutationResponse>;
  /** update single row of the table: "user_account" */
  update_user_account_by_pk?: Maybe<UserAccount>;
};


/** mutation root */
export type MutationRootConfirmInvitationArgs = {
  input: ConfirmInvitationInput;
};


/** mutation root */
export type MutationRootDeleteActionArgs = {
  where: ActionBoolExp;
};


/** mutation root */
export type MutationRootDeleteActionByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteAwsAccountArgs = {
  where: AwsAccountBoolExp;
};


/** mutation root */
export type MutationRootDeleteAwsAccountByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteInstanceArgs = {
  where: InstanceBoolExp;
};


/** mutation root */
export type MutationRootDeleteInstanceAtServiceArgs = {
  where: InstanceAtServiceBoolExp;
};


/** mutation root */
export type MutationRootDeleteInstanceAtServiceByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteInstanceByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteMetricsArgs = {
  where: MetricsBoolExp;
};


/** mutation root */
export type MutationRootDeleteNotificationRuleArgs = {
  where: NotificationRuleBoolExp;
};


/** mutation root */
export type MutationRootDeleteNotificationRuleByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteOrganizationArgs = {
  where: OrganizationBoolExp;
};


/** mutation root */
export type MutationRootDeleteOrganizationByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteProjectArgs = {
  where: ProjectBoolExp;
};


/** mutation root */
export type MutationRootDeleteProjectByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteProjectCollaboratorArgs = {
  where: ProjectCollaboratorBoolExp;
};


/** mutation root */
export type MutationRootDeleteProjectInOrganizationArgs = {
  where: ProjectInOrganizationBoolExp;
};


/** mutation root */
export type MutationRootDeleteProjectInOrganizationByPkArgs = {
  organization_id: Scalars['bigint'];
  project_id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteProjectInvitationArgs = {
  where: ProjectInvitationBoolExp;
};


/** mutation root */
export type MutationRootDeleteProjectInvitationByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteServiceArgs = {
  where: ServiceBoolExp;
};


/** mutation root */
export type MutationRootDeleteServiceByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteSlackWebhookArgs = {
  where: SlackWebhookBoolExp;
};


/** mutation root */
export type MutationRootDeleteSlackWebhookByPkArgs = {
  id: Scalars['bigint'];
};


/** mutation root */
export type MutationRootDeleteUserAccountArgs = {
  where: UserAccountBoolExp;
};


/** mutation root */
export type MutationRootDeleteUserAccountByPkArgs = {
  id: Scalars['String'];
};


/** mutation root */
export type MutationRootInsertActionArgs = {
  objects: Array<ActionInsertInput>;
  on_conflict?: Maybe<ActionOnConflict>;
};


/** mutation root */
export type MutationRootInsertActionOneArgs = {
  object: ActionInsertInput;
  on_conflict?: Maybe<ActionOnConflict>;
};


/** mutation root */
export type MutationRootInsertAwsAccountArgs = {
  objects: Array<AwsAccountInsertInput>;
  on_conflict?: Maybe<AwsAccountOnConflict>;
};


/** mutation root */
export type MutationRootInsertAwsAccountOneArgs = {
  object: AwsAccountInsertInput;
  on_conflict?: Maybe<AwsAccountOnConflict>;
};


/** mutation root */
export type MutationRootInsertInstanceArgs = {
  objects: Array<InstanceInsertInput>;
  on_conflict?: Maybe<InstanceOnConflict>;
};


/** mutation root */
export type MutationRootInsertInstanceAtServiceArgs = {
  objects: Array<InstanceAtServiceInsertInput>;
  on_conflict?: Maybe<InstanceAtServiceOnConflict>;
};


/** mutation root */
export type MutationRootInsertInstanceAtServiceOneArgs = {
  object: InstanceAtServiceInsertInput;
  on_conflict?: Maybe<InstanceAtServiceOnConflict>;
};


/** mutation root */
export type MutationRootInsertInstanceOneArgs = {
  object: InstanceInsertInput;
  on_conflict?: Maybe<InstanceOnConflict>;
};


/** mutation root */
export type MutationRootInsertMetricsArgs = {
  objects: Array<MetricsInsertInput>;
};


/** mutation root */
export type MutationRootInsertMetricsOneArgs = {
  object: MetricsInsertInput;
};


/** mutation root */
export type MutationRootInsertNotificationRuleArgs = {
  objects: Array<NotificationRuleInsertInput>;
  on_conflict?: Maybe<NotificationRuleOnConflict>;
};


/** mutation root */
export type MutationRootInsertNotificationRuleOneArgs = {
  object: NotificationRuleInsertInput;
  on_conflict?: Maybe<NotificationRuleOnConflict>;
};


/** mutation root */
export type MutationRootInsertOrganizationArgs = {
  objects: Array<OrganizationInsertInput>;
  on_conflict?: Maybe<OrganizationOnConflict>;
};


/** mutation root */
export type MutationRootInsertOrganizationOneArgs = {
  object: OrganizationInsertInput;
  on_conflict?: Maybe<OrganizationOnConflict>;
};


/** mutation root */
export type MutationRootInsertProjectArgs = {
  objects: Array<ProjectInsertInput>;
  on_conflict?: Maybe<ProjectOnConflict>;
};


/** mutation root */
export type MutationRootInsertProjectCollaboratorArgs = {
  objects: Array<ProjectCollaboratorInsertInput>;
};


/** mutation root */
export type MutationRootInsertProjectCollaboratorOneArgs = {
  object: ProjectCollaboratorInsertInput;
};


/** mutation root */
export type MutationRootInsertProjectInOrganizationArgs = {
  objects: Array<ProjectInOrganizationInsertInput>;
  on_conflict?: Maybe<ProjectInOrganizationOnConflict>;
};


/** mutation root */
export type MutationRootInsertProjectInOrganizationOneArgs = {
  object: ProjectInOrganizationInsertInput;
  on_conflict?: Maybe<ProjectInOrganizationOnConflict>;
};


/** mutation root */
export type MutationRootInsertProjectInvitationArgs = {
  objects: Array<ProjectInvitationInsertInput>;
  on_conflict?: Maybe<ProjectInvitationOnConflict>;
};


/** mutation root */
export type MutationRootInsertProjectInvitationOneArgs = {
  object: ProjectInvitationInsertInput;
  on_conflict?: Maybe<ProjectInvitationOnConflict>;
};


/** mutation root */
export type MutationRootInsertProjectOneArgs = {
  object: ProjectInsertInput;
  on_conflict?: Maybe<ProjectOnConflict>;
};


/** mutation root */
export type MutationRootInsertServiceArgs = {
  objects: Array<ServiceInsertInput>;
  on_conflict?: Maybe<ServiceOnConflict>;
};


/** mutation root */
export type MutationRootInsertServiceOneArgs = {
  object: ServiceInsertInput;
  on_conflict?: Maybe<ServiceOnConflict>;
};


/** mutation root */
export type MutationRootInsertSlackWebhookArgs = {
  objects: Array<SlackWebhookInsertInput>;
  on_conflict?: Maybe<SlackWebhookOnConflict>;
};


/** mutation root */
export type MutationRootInsertSlackWebhookOneArgs = {
  object: SlackWebhookInsertInput;
  on_conflict?: Maybe<SlackWebhookOnConflict>;
};


/** mutation root */
export type MutationRootInsertUserAccountArgs = {
  objects: Array<UserAccountInsertInput>;
  on_conflict?: Maybe<UserAccountOnConflict>;
};


/** mutation root */
export type MutationRootInsertUserAccountOneArgs = {
  object: UserAccountInsertInput;
  on_conflict?: Maybe<UserAccountOnConflict>;
};


/** mutation root */
export type MutationRootUpdateActionArgs = {
  _append?: Maybe<ActionAppendInput>;
  _delete_at_path?: Maybe<ActionDeleteAtPathInput>;
  _delete_elem?: Maybe<ActionDeleteElemInput>;
  _delete_key?: Maybe<ActionDeleteKeyInput>;
  _inc?: Maybe<ActionIncInput>;
  _prepend?: Maybe<ActionPrependInput>;
  _set?: Maybe<ActionSetInput>;
  where: ActionBoolExp;
};


/** mutation root */
export type MutationRootUpdateActionByPkArgs = {
  _append?: Maybe<ActionAppendInput>;
  _delete_at_path?: Maybe<ActionDeleteAtPathInput>;
  _delete_elem?: Maybe<ActionDeleteElemInput>;
  _delete_key?: Maybe<ActionDeleteKeyInput>;
  _inc?: Maybe<ActionIncInput>;
  _prepend?: Maybe<ActionPrependInput>;
  _set?: Maybe<ActionSetInput>;
  pk_columns: ActionPkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateAwsAccountArgs = {
  _inc?: Maybe<AwsAccountIncInput>;
  _set?: Maybe<AwsAccountSetInput>;
  where: AwsAccountBoolExp;
};


/** mutation root */
export type MutationRootUpdateAwsAccountByPkArgs = {
  _inc?: Maybe<AwsAccountIncInput>;
  _set?: Maybe<AwsAccountSetInput>;
  pk_columns: AwsAccountPkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateInstanceArgs = {
  _inc?: Maybe<InstanceIncInput>;
  _set?: Maybe<InstanceSetInput>;
  where: InstanceBoolExp;
};


/** mutation root */
export type MutationRootUpdateInstanceAtServiceArgs = {
  _inc?: Maybe<InstanceAtServiceIncInput>;
  _set?: Maybe<InstanceAtServiceSetInput>;
  where: InstanceAtServiceBoolExp;
};


/** mutation root */
export type MutationRootUpdateInstanceAtServiceByPkArgs = {
  _inc?: Maybe<InstanceAtServiceIncInput>;
  _set?: Maybe<InstanceAtServiceSetInput>;
  pk_columns: InstanceAtServicePkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateInstanceByPkArgs = {
  _inc?: Maybe<InstanceIncInput>;
  _set?: Maybe<InstanceSetInput>;
  pk_columns: InstancePkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateMetricsArgs = {
  _set?: Maybe<MetricsSetInput>;
  where: MetricsBoolExp;
};


/** mutation root */
export type MutationRootUpdateNotificationRuleArgs = {
  _append?: Maybe<NotificationRuleAppendInput>;
  _delete_at_path?: Maybe<NotificationRuleDeleteAtPathInput>;
  _delete_elem?: Maybe<NotificationRuleDeleteElemInput>;
  _delete_key?: Maybe<NotificationRuleDeleteKeyInput>;
  _inc?: Maybe<NotificationRuleIncInput>;
  _prepend?: Maybe<NotificationRulePrependInput>;
  _set?: Maybe<NotificationRuleSetInput>;
  where: NotificationRuleBoolExp;
};


/** mutation root */
export type MutationRootUpdateNotificationRuleByPkArgs = {
  _append?: Maybe<NotificationRuleAppendInput>;
  _delete_at_path?: Maybe<NotificationRuleDeleteAtPathInput>;
  _delete_elem?: Maybe<NotificationRuleDeleteElemInput>;
  _delete_key?: Maybe<NotificationRuleDeleteKeyInput>;
  _inc?: Maybe<NotificationRuleIncInput>;
  _prepend?: Maybe<NotificationRulePrependInput>;
  _set?: Maybe<NotificationRuleSetInput>;
  pk_columns: NotificationRulePkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateOrganizationArgs = {
  _inc?: Maybe<OrganizationIncInput>;
  _set?: Maybe<OrganizationSetInput>;
  where: OrganizationBoolExp;
};


/** mutation root */
export type MutationRootUpdateOrganizationByPkArgs = {
  _inc?: Maybe<OrganizationIncInput>;
  _set?: Maybe<OrganizationSetInput>;
  pk_columns: OrganizationPkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateProjectArgs = {
  _inc?: Maybe<ProjectIncInput>;
  _set?: Maybe<ProjectSetInput>;
  where: ProjectBoolExp;
};


/** mutation root */
export type MutationRootUpdateProjectByPkArgs = {
  _inc?: Maybe<ProjectIncInput>;
  _set?: Maybe<ProjectSetInput>;
  pk_columns: ProjectPkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateProjectCollaboratorArgs = {
  _inc?: Maybe<ProjectCollaboratorIncInput>;
  _set?: Maybe<ProjectCollaboratorSetInput>;
  where: ProjectCollaboratorBoolExp;
};


/** mutation root */
export type MutationRootUpdateProjectInOrganizationArgs = {
  _inc?: Maybe<ProjectInOrganizationIncInput>;
  _set?: Maybe<ProjectInOrganizationSetInput>;
  where: ProjectInOrganizationBoolExp;
};


/** mutation root */
export type MutationRootUpdateProjectInOrganizationByPkArgs = {
  _inc?: Maybe<ProjectInOrganizationIncInput>;
  _set?: Maybe<ProjectInOrganizationSetInput>;
  pk_columns: ProjectInOrganizationPkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateProjectInvitationArgs = {
  _inc?: Maybe<ProjectInvitationIncInput>;
  _set?: Maybe<ProjectInvitationSetInput>;
  where: ProjectInvitationBoolExp;
};


/** mutation root */
export type MutationRootUpdateProjectInvitationByPkArgs = {
  _inc?: Maybe<ProjectInvitationIncInput>;
  _set?: Maybe<ProjectInvitationSetInput>;
  pk_columns: ProjectInvitationPkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateServiceArgs = {
  _inc?: Maybe<ServiceIncInput>;
  _set?: Maybe<ServiceSetInput>;
  where: ServiceBoolExp;
};


/** mutation root */
export type MutationRootUpdateServiceByPkArgs = {
  _inc?: Maybe<ServiceIncInput>;
  _set?: Maybe<ServiceSetInput>;
  pk_columns: ServicePkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateSlackWebhookArgs = {
  _inc?: Maybe<SlackWebhookIncInput>;
  _set?: Maybe<SlackWebhookSetInput>;
  where: SlackWebhookBoolExp;
};


/** mutation root */
export type MutationRootUpdateSlackWebhookByPkArgs = {
  _inc?: Maybe<SlackWebhookIncInput>;
  _set?: Maybe<SlackWebhookSetInput>;
  pk_columns: SlackWebhookPkColumnsInput;
};


/** mutation root */
export type MutationRootUpdateUserAccountArgs = {
  _set?: Maybe<UserAccountSetInput>;
  where: UserAccountBoolExp;
};


/** mutation root */
export type MutationRootUpdateUserAccountByPkArgs = {
  _set?: Maybe<UserAccountSetInput>;
  pk_columns: UserAccountPkColumnsInput;
};

/** columns and relationships of "notification_rule" */
export type NotificationRule = {
   __typename?: 'notification_rule';
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  rule_name: Scalars['String'];
  rules: Scalars['jsonb'];
  updated_at: Scalars['timestamp'];
};


/** columns and relationships of "notification_rule" */
export type NotificationRuleRulesArgs = {
  path?: Maybe<Scalars['String']>;
};

/** aggregated selection of "notification_rule" */
export type NotificationRuleAggregate = {
   __typename?: 'notification_rule_aggregate';
  aggregate?: Maybe<NotificationRuleAggregateFields>;
  nodes: Array<NotificationRule>;
};

/** aggregate fields of "notification_rule" */
export type NotificationRuleAggregateFields = {
   __typename?: 'notification_rule_aggregate_fields';
  avg?: Maybe<NotificationRuleAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<NotificationRuleMaxFields>;
  min?: Maybe<NotificationRuleMinFields>;
  stddev?: Maybe<NotificationRuleStddevFields>;
  stddev_pop?: Maybe<NotificationRuleStddevPopFields>;
  stddev_samp?: Maybe<NotificationRuleStddevSampFields>;
  sum?: Maybe<NotificationRuleSumFields>;
  var_pop?: Maybe<NotificationRuleVarPopFields>;
  var_samp?: Maybe<NotificationRuleVarSampFields>;
  variance?: Maybe<NotificationRuleVarianceFields>;
};


/** aggregate fields of "notification_rule" */
export type NotificationRuleAggregateFieldsCountArgs = {
  columns?: Maybe<Array<NotificationRuleSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "notification_rule" */
export type NotificationRuleAggregateOrderBy = {
  avg?: Maybe<NotificationRuleAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<NotificationRuleMaxOrderBy>;
  min?: Maybe<NotificationRuleMinOrderBy>;
  stddev?: Maybe<NotificationRuleStddevOrderBy>;
  stddev_pop?: Maybe<NotificationRuleStddevPopOrderBy>;
  stddev_samp?: Maybe<NotificationRuleStddevSampOrderBy>;
  sum?: Maybe<NotificationRuleSumOrderBy>;
  var_pop?: Maybe<NotificationRuleVarPopOrderBy>;
  var_samp?: Maybe<NotificationRuleVarSampOrderBy>;
  variance?: Maybe<NotificationRuleVarianceOrderBy>;
};

/** append existing jsonb value of filtered columns with new jsonb value */
export type NotificationRuleAppendInput = {
  rules?: Maybe<Scalars['jsonb']>;
};

/** input type for inserting array relation for remote table "notification_rule" */
export type NotificationRuleArrRelInsertInput = {
  data: Array<NotificationRuleInsertInput>;
  on_conflict?: Maybe<NotificationRuleOnConflict>;
};

/** aggregate avg on columns */
export type NotificationRuleAvgFields = {
   __typename?: 'notification_rule_avg_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "notification_rule" */
export type NotificationRuleAvgOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "notification_rule". All fields are combined with a logical 'AND'. */
export type NotificationRuleBoolExp = {
  _and?: Maybe<Array<Maybe<NotificationRuleBoolExp>>>;
  _not?: Maybe<NotificationRuleBoolExp>;
  _or?: Maybe<Array<Maybe<NotificationRuleBoolExp>>>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  rule_name?: Maybe<StringComparisonExp>;
  rules?: Maybe<JsonbComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
};

/** unique or primary key constraints on table "notification_rule" */
export enum NotificationRuleConstraint {
  /** unique or primary key constraint */
  NOTIFICATION_RULE_PKEY = 'notification_rule_pkey'
}

/** delete the field or element with specified path (for JSON arrays, negative integers count from the end) */
export type NotificationRuleDeleteAtPathInput = {
  rules?: Maybe<Array<Maybe<Scalars['String']>>>;
};

/** delete the array element with specified index (negative integers count from the end). throws an error if top level container is not an array */
export type NotificationRuleDeleteElemInput = {
  rules?: Maybe<Scalars['Int']>;
};

/** delete key/value pair or string element. key/value pairs are matched based on their key value */
export type NotificationRuleDeleteKeyInput = {
  rules?: Maybe<Scalars['String']>;
};

/** input type for incrementing integer columne in table "notification_rule" */
export type NotificationRuleIncInput = {
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "notification_rule" */
export type NotificationRuleInsertInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  rule_name?: Maybe<Scalars['String']>;
  rules?: Maybe<Scalars['jsonb']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate max on columns */
export type NotificationRuleMaxFields = {
   __typename?: 'notification_rule_max_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
  rule_name?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "notification_rule" */
export type NotificationRuleMaxOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  rule_name?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type NotificationRuleMinFields = {
   __typename?: 'notification_rule_min_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
  rule_name?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "notification_rule" */
export type NotificationRuleMinOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  rule_name?: Maybe<OrderBy>;
};

/** response of any mutation on the table "notification_rule" */
export type NotificationRuleMutationResponse = {
   __typename?: 'notification_rule_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<NotificationRule>;
};

/** input type for inserting object relation for remote table "notification_rule" */
export type NotificationRuleObjRelInsertInput = {
  data: NotificationRuleInsertInput;
  on_conflict?: Maybe<NotificationRuleOnConflict>;
};

/** on conflict condition type for table "notification_rule" */
export type NotificationRuleOnConflict = {
  constraint: NotificationRuleConstraint;
  update_columns: Array<NotificationRuleUpdateColumn>;
  where?: Maybe<NotificationRuleBoolExp>;
};

/** ordering options when selecting data from "notification_rule" */
export type NotificationRuleOrderBy = {
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  rule_name?: Maybe<OrderBy>;
  rules?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
};

/** primary key columns input for table: "notification_rule" */
export type NotificationRulePkColumnsInput = {
  id: Scalars['bigint'];
};

/** prepend existing jsonb value of filtered columns with new jsonb value */
export type NotificationRulePrependInput = {
  rules?: Maybe<Scalars['jsonb']>;
};

/** select columns of table "notification_rule" */
export enum NotificationRuleSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  RULE_NAME = 'rule_name',
  /** column name */
  RULES = 'rules',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** input type for updating data in table "notification_rule" */
export type NotificationRuleSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
  rule_name?: Maybe<Scalars['String']>;
  rules?: Maybe<Scalars['jsonb']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate stddev on columns */
export type NotificationRuleStddevFields = {
   __typename?: 'notification_rule_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "notification_rule" */
export type NotificationRuleStddevOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type NotificationRuleStddevPopFields = {
   __typename?: 'notification_rule_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "notification_rule" */
export type NotificationRuleStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type NotificationRuleStddevSampFields = {
   __typename?: 'notification_rule_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "notification_rule" */
export type NotificationRuleStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type NotificationRuleSumFields = {
   __typename?: 'notification_rule_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "notification_rule" */
export type NotificationRuleSumOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "notification_rule" */
export enum NotificationRuleUpdateColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  RULE_NAME = 'rule_name',
  /** column name */
  RULES = 'rules',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** aggregate var_pop on columns */
export type NotificationRuleVarPopFields = {
   __typename?: 'notification_rule_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "notification_rule" */
export type NotificationRuleVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type NotificationRuleVarSampFields = {
   __typename?: 'notification_rule_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "notification_rule" */
export type NotificationRuleVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type NotificationRuleVarianceFields = {
   __typename?: 'notification_rule_variance_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "notification_rule" */
export type NotificationRuleVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** column ordering options */
export enum OrderBy {
  /** in the ascending order, nulls last */
  ASC = 'asc',
  /** in the ascending order, nulls first */
  ASC_NULLS_FIRST = 'asc_nulls_first',
  /** in the ascending order, nulls last */
  ASC_NULLS_LAST = 'asc_nulls_last',
  /** in the descending order, nulls first */
  DESC = 'desc',
  /** in the descending order, nulls first */
  DESC_NULLS_FIRST = 'desc_nulls_first',
  /** in the descending order, nulls last */
  DESC_NULLS_LAST = 'desc_nulls_last'
}

/** columns and relationships of "organization" */
export type Organization = {
   __typename?: 'organization';
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  name: Scalars['String'];
  /** An array relationship */
  project_in_organizations: Array<ProjectInOrganization>;
  /** An aggregated array relationship */
  project_in_organizations_aggregate: ProjectInOrganizationAggregate;
  unique_name: Scalars['String'];
  updated_at: Scalars['timestamp'];
};


/** columns and relationships of "organization" */
export type OrganizationProjectInOrganizationsArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};


/** columns and relationships of "organization" */
export type OrganizationProjectInOrganizationsAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};

/** aggregated selection of "organization" */
export type OrganizationAggregate = {
   __typename?: 'organization_aggregate';
  aggregate?: Maybe<OrganizationAggregateFields>;
  nodes: Array<Organization>;
};

/** aggregate fields of "organization" */
export type OrganizationAggregateFields = {
   __typename?: 'organization_aggregate_fields';
  avg?: Maybe<OrganizationAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<OrganizationMaxFields>;
  min?: Maybe<OrganizationMinFields>;
  stddev?: Maybe<OrganizationStddevFields>;
  stddev_pop?: Maybe<OrganizationStddevPopFields>;
  stddev_samp?: Maybe<OrganizationStddevSampFields>;
  sum?: Maybe<OrganizationSumFields>;
  var_pop?: Maybe<OrganizationVarPopFields>;
  var_samp?: Maybe<OrganizationVarSampFields>;
  variance?: Maybe<OrganizationVarianceFields>;
};


/** aggregate fields of "organization" */
export type OrganizationAggregateFieldsCountArgs = {
  columns?: Maybe<Array<OrganizationSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "organization" */
export type OrganizationAggregateOrderBy = {
  avg?: Maybe<OrganizationAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<OrganizationMaxOrderBy>;
  min?: Maybe<OrganizationMinOrderBy>;
  stddev?: Maybe<OrganizationStddevOrderBy>;
  stddev_pop?: Maybe<OrganizationStddevPopOrderBy>;
  stddev_samp?: Maybe<OrganizationStddevSampOrderBy>;
  sum?: Maybe<OrganizationSumOrderBy>;
  var_pop?: Maybe<OrganizationVarPopOrderBy>;
  var_samp?: Maybe<OrganizationVarSampOrderBy>;
  variance?: Maybe<OrganizationVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "organization" */
export type OrganizationArrRelInsertInput = {
  data: Array<OrganizationInsertInput>;
  on_conflict?: Maybe<OrganizationOnConflict>;
};

/** aggregate avg on columns */
export type OrganizationAvgFields = {
   __typename?: 'organization_avg_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "organization" */
export type OrganizationAvgOrderBy = {
  id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "organization". All fields are combined with a logical 'AND'. */
export type OrganizationBoolExp = {
  _and?: Maybe<Array<Maybe<OrganizationBoolExp>>>;
  _not?: Maybe<OrganizationBoolExp>;
  _or?: Maybe<Array<Maybe<OrganizationBoolExp>>>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  name?: Maybe<StringComparisonExp>;
  project_in_organizations?: Maybe<ProjectInOrganizationBoolExp>;
  unique_name?: Maybe<StringComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
};

/** unique or primary key constraints on table "organization" */
export enum OrganizationConstraint {
  /** unique or primary key constraint */
  ORGANIZATION_PKEY = 'organization_pkey'
}

/** input type for incrementing integer columne in table "organization" */
export type OrganizationIncInput = {
  id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "organization" */
export type OrganizationInsertInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_in_organizations?: Maybe<ProjectInOrganizationArrRelInsertInput>;
  unique_name?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate max on columns */
export type OrganizationMaxFields = {
   __typename?: 'organization_max_fields';
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  unique_name?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "organization" */
export type OrganizationMaxOrderBy = {
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  unique_name?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type OrganizationMinFields = {
   __typename?: 'organization_min_fields';
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  unique_name?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "organization" */
export type OrganizationMinOrderBy = {
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  unique_name?: Maybe<OrderBy>;
};

/** response of any mutation on the table "organization" */
export type OrganizationMutationResponse = {
   __typename?: 'organization_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Organization>;
};

/** input type for inserting object relation for remote table "organization" */
export type OrganizationObjRelInsertInput = {
  data: OrganizationInsertInput;
  on_conflict?: Maybe<OrganizationOnConflict>;
};

/** on conflict condition type for table "organization" */
export type OrganizationOnConflict = {
  constraint: OrganizationConstraint;
  update_columns: Array<OrganizationUpdateColumn>;
  where?: Maybe<OrganizationBoolExp>;
};

/** ordering options when selecting data from "organization" */
export type OrganizationOrderBy = {
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_in_organizations_aggregate?: Maybe<ProjectInOrganizationAggregateOrderBy>;
  unique_name?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
};

/** primary key columns input for table: "organization" */
export type OrganizationPkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "organization" */
export enum OrganizationSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  UNIQUE_NAME = 'unique_name',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** input type for updating data in table "organization" */
export type OrganizationSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  unique_name?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate stddev on columns */
export type OrganizationStddevFields = {
   __typename?: 'organization_stddev_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "organization" */
export type OrganizationStddevOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type OrganizationStddevPopFields = {
   __typename?: 'organization_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "organization" */
export type OrganizationStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type OrganizationStddevSampFields = {
   __typename?: 'organization_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "organization" */
export type OrganizationStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type OrganizationSumFields = {
   __typename?: 'organization_sum_fields';
  id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "organization" */
export type OrganizationSumOrderBy = {
  id?: Maybe<OrderBy>;
};

/** update columns of table "organization" */
export enum OrganizationUpdateColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  UNIQUE_NAME = 'unique_name',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** aggregate var_pop on columns */
export type OrganizationVarPopFields = {
   __typename?: 'organization_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "organization" */
export type OrganizationVarPopOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type OrganizationVarSampFields = {
   __typename?: 'organization_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "organization" */
export type OrganizationVarSampOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type OrganizationVarianceFields = {
   __typename?: 'organization_variance_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "organization" */
export type OrganizationVarianceOrderBy = {
  id?: Maybe<OrderBy>;
};

/** columns and relationships of "project" */
export type Project = {
   __typename?: 'project';
  /** An array relationship */
  aws_accounts: Array<AwsAccount>;
  /** An aggregated array relationship */
  aws_accounts_aggregate: AwsAccountAggregate;
  created_at: Scalars['timestamp'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['bigint'];
  /** An array relationship */
  instances: Array<Instance>;
  /** An aggregated array relationship */
  instances_aggregate: InstanceAggregate;
  name: Scalars['String'];
  /** An array relationship */
  project_collaborators: Array<ProjectCollaborator>;
  /** An aggregated array relationship */
  project_collaborators_aggregate: ProjectCollaboratorAggregate;
  /** An array relationship */
  project_in_organizations: Array<ProjectInOrganization>;
  /** An aggregated array relationship */
  project_in_organizations_aggregate: ProjectInOrganizationAggregate;
  updated_at: Scalars['timestamp'];
  /** An object relationship */
  user_account: UserAccount;
  user_id: Scalars['String'];
};


/** columns and relationships of "project" */
export type ProjectAwsAccountsArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** columns and relationships of "project" */
export type ProjectAwsAccountsAggregateArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** columns and relationships of "project" */
export type ProjectInstancesArgs = {
  distinct_on?: Maybe<Array<InstanceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceOrderBy>>;
  where?: Maybe<InstanceBoolExp>;
};


/** columns and relationships of "project" */
export type ProjectInstancesAggregateArgs = {
  distinct_on?: Maybe<Array<InstanceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceOrderBy>>;
  where?: Maybe<InstanceBoolExp>;
};


/** columns and relationships of "project" */
export type ProjectProjectCollaboratorsArgs = {
  distinct_on?: Maybe<Array<ProjectCollaboratorSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectCollaboratorOrderBy>>;
  where?: Maybe<ProjectCollaboratorBoolExp>;
};


/** columns and relationships of "project" */
export type ProjectProjectCollaboratorsAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectCollaboratorSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectCollaboratorOrderBy>>;
  where?: Maybe<ProjectCollaboratorBoolExp>;
};


/** columns and relationships of "project" */
export type ProjectProjectInOrganizationsArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};


/** columns and relationships of "project" */
export type ProjectProjectInOrganizationsAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};

/** aggregated selection of "project" */
export type ProjectAggregate = {
   __typename?: 'project_aggregate';
  aggregate?: Maybe<ProjectAggregateFields>;
  nodes: Array<Project>;
};

/** aggregate fields of "project" */
export type ProjectAggregateFields = {
   __typename?: 'project_aggregate_fields';
  avg?: Maybe<ProjectAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<ProjectMaxFields>;
  min?: Maybe<ProjectMinFields>;
  stddev?: Maybe<ProjectStddevFields>;
  stddev_pop?: Maybe<ProjectStddevPopFields>;
  stddev_samp?: Maybe<ProjectStddevSampFields>;
  sum?: Maybe<ProjectSumFields>;
  var_pop?: Maybe<ProjectVarPopFields>;
  var_samp?: Maybe<ProjectVarSampFields>;
  variance?: Maybe<ProjectVarianceFields>;
};


/** aggregate fields of "project" */
export type ProjectAggregateFieldsCountArgs = {
  columns?: Maybe<Array<ProjectSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "project" */
export type ProjectAggregateOrderBy = {
  avg?: Maybe<ProjectAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<ProjectMaxOrderBy>;
  min?: Maybe<ProjectMinOrderBy>;
  stddev?: Maybe<ProjectStddevOrderBy>;
  stddev_pop?: Maybe<ProjectStddevPopOrderBy>;
  stddev_samp?: Maybe<ProjectStddevSampOrderBy>;
  sum?: Maybe<ProjectSumOrderBy>;
  var_pop?: Maybe<ProjectVarPopOrderBy>;
  var_samp?: Maybe<ProjectVarSampOrderBy>;
  variance?: Maybe<ProjectVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "project" */
export type ProjectArrRelInsertInput = {
  data: Array<ProjectInsertInput>;
  on_conflict?: Maybe<ProjectOnConflict>;
};

/** aggregate avg on columns */
export type ProjectAvgFields = {
   __typename?: 'project_avg_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "project" */
export type ProjectAvgOrderBy = {
  id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "project". All fields are combined with a logical 'AND'. */
export type ProjectBoolExp = {
  _and?: Maybe<Array<Maybe<ProjectBoolExp>>>;
  _not?: Maybe<ProjectBoolExp>;
  _or?: Maybe<Array<Maybe<ProjectBoolExp>>>;
  aws_accounts?: Maybe<AwsAccountBoolExp>;
  created_at?: Maybe<TimestampComparisonExp>;
  description?: Maybe<StringComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  instances?: Maybe<InstanceBoolExp>;
  name?: Maybe<StringComparisonExp>;
  project_collaborators?: Maybe<ProjectCollaboratorBoolExp>;
  project_in_organizations?: Maybe<ProjectInOrganizationBoolExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
  user_account?: Maybe<UserAccountBoolExp>;
  user_id?: Maybe<StringComparisonExp>;
};

/** columns and relationships of "project_collaborator" */
export type ProjectCollaborator = {
   __typename?: 'project_collaborator';
  created_at: Scalars['timestamp'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  updated_at: Scalars['timestamp'];
  user_id: Scalars['String'];
};

/** aggregated selection of "project_collaborator" */
export type ProjectCollaboratorAggregate = {
   __typename?: 'project_collaborator_aggregate';
  aggregate?: Maybe<ProjectCollaboratorAggregateFields>;
  nodes: Array<ProjectCollaborator>;
};

/** aggregate fields of "project_collaborator" */
export type ProjectCollaboratorAggregateFields = {
   __typename?: 'project_collaborator_aggregate_fields';
  avg?: Maybe<ProjectCollaboratorAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<ProjectCollaboratorMaxFields>;
  min?: Maybe<ProjectCollaboratorMinFields>;
  stddev?: Maybe<ProjectCollaboratorStddevFields>;
  stddev_pop?: Maybe<ProjectCollaboratorStddevPopFields>;
  stddev_samp?: Maybe<ProjectCollaboratorStddevSampFields>;
  sum?: Maybe<ProjectCollaboratorSumFields>;
  var_pop?: Maybe<ProjectCollaboratorVarPopFields>;
  var_samp?: Maybe<ProjectCollaboratorVarSampFields>;
  variance?: Maybe<ProjectCollaboratorVarianceFields>;
};


/** aggregate fields of "project_collaborator" */
export type ProjectCollaboratorAggregateFieldsCountArgs = {
  columns?: Maybe<Array<ProjectCollaboratorSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "project_collaborator" */
export type ProjectCollaboratorAggregateOrderBy = {
  avg?: Maybe<ProjectCollaboratorAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<ProjectCollaboratorMaxOrderBy>;
  min?: Maybe<ProjectCollaboratorMinOrderBy>;
  stddev?: Maybe<ProjectCollaboratorStddevOrderBy>;
  stddev_pop?: Maybe<ProjectCollaboratorStddevPopOrderBy>;
  stddev_samp?: Maybe<ProjectCollaboratorStddevSampOrderBy>;
  sum?: Maybe<ProjectCollaboratorSumOrderBy>;
  var_pop?: Maybe<ProjectCollaboratorVarPopOrderBy>;
  var_samp?: Maybe<ProjectCollaboratorVarSampOrderBy>;
  variance?: Maybe<ProjectCollaboratorVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "project_collaborator" */
export type ProjectCollaboratorArrRelInsertInput = {
  data: Array<ProjectCollaboratorInsertInput>;
};

/** aggregate avg on columns */
export type ProjectCollaboratorAvgFields = {
   __typename?: 'project_collaborator_avg_fields';
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "project_collaborator" */
export type ProjectCollaboratorAvgOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "project_collaborator". All fields are combined with a logical 'AND'. */
export type ProjectCollaboratorBoolExp = {
  _and?: Maybe<Array<Maybe<ProjectCollaboratorBoolExp>>>;
  _not?: Maybe<ProjectCollaboratorBoolExp>;
  _or?: Maybe<Array<Maybe<ProjectCollaboratorBoolExp>>>;
  created_at?: Maybe<TimestampComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
  user_id?: Maybe<StringComparisonExp>;
};

/** input type for incrementing integer columne in table "project_collaborator" */
export type ProjectCollaboratorIncInput = {
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "project_collaborator" */
export type ProjectCollaboratorInsertInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type ProjectCollaboratorMaxFields = {
   __typename?: 'project_collaborator_max_fields';
  project_id?: Maybe<Scalars['bigint']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "project_collaborator" */
export type ProjectCollaboratorMaxOrderBy = {
  project_id?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type ProjectCollaboratorMinFields = {
   __typename?: 'project_collaborator_min_fields';
  project_id?: Maybe<Scalars['bigint']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "project_collaborator" */
export type ProjectCollaboratorMinOrderBy = {
  project_id?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** response of any mutation on the table "project_collaborator" */
export type ProjectCollaboratorMutationResponse = {
   __typename?: 'project_collaborator_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<ProjectCollaborator>;
};

/** input type for inserting object relation for remote table "project_collaborator" */
export type ProjectCollaboratorObjRelInsertInput = {
  data: ProjectCollaboratorInsertInput;
};

/** ordering options when selecting data from "project_collaborator" */
export type ProjectCollaboratorOrderBy = {
  created_at?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** select columns of table "project_collaborator" */
export enum ProjectCollaboratorSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id'
}

/** input type for updating data in table "project_collaborator" */
export type ProjectCollaboratorSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type ProjectCollaboratorStddevFields = {
   __typename?: 'project_collaborator_stddev_fields';
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "project_collaborator" */
export type ProjectCollaboratorStddevOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type ProjectCollaboratorStddevPopFields = {
   __typename?: 'project_collaborator_stddev_pop_fields';
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "project_collaborator" */
export type ProjectCollaboratorStddevPopOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type ProjectCollaboratorStddevSampFields = {
   __typename?: 'project_collaborator_stddev_samp_fields';
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "project_collaborator" */
export type ProjectCollaboratorStddevSampOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type ProjectCollaboratorSumFields = {
   __typename?: 'project_collaborator_sum_fields';
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "project_collaborator" */
export type ProjectCollaboratorSumOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_pop on columns */
export type ProjectCollaboratorVarPopFields = {
   __typename?: 'project_collaborator_var_pop_fields';
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "project_collaborator" */
export type ProjectCollaboratorVarPopOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type ProjectCollaboratorVarSampFields = {
   __typename?: 'project_collaborator_var_samp_fields';
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "project_collaborator" */
export type ProjectCollaboratorVarSampOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type ProjectCollaboratorVarianceFields = {
   __typename?: 'project_collaborator_variance_fields';
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "project_collaborator" */
export type ProjectCollaboratorVarianceOrderBy = {
  project_id?: Maybe<OrderBy>;
};

/** unique or primary key constraints on table "project" */
export enum ProjectConstraint {
  /** unique or primary key constraint */
  PROJECT_PKEY = 'project_pkey'
}

/** columns and relationships of "project_in_organization" */
export type ProjectInOrganization = {
   __typename?: 'project_in_organization';
  /** An object relationship */
  organization: Organization;
  organization_id: Scalars['bigint'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
};

/** aggregated selection of "project_in_organization" */
export type ProjectInOrganizationAggregate = {
   __typename?: 'project_in_organization_aggregate';
  aggregate?: Maybe<ProjectInOrganizationAggregateFields>;
  nodes: Array<ProjectInOrganization>;
};

/** aggregate fields of "project_in_organization" */
export type ProjectInOrganizationAggregateFields = {
   __typename?: 'project_in_organization_aggregate_fields';
  avg?: Maybe<ProjectInOrganizationAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<ProjectInOrganizationMaxFields>;
  min?: Maybe<ProjectInOrganizationMinFields>;
  stddev?: Maybe<ProjectInOrganizationStddevFields>;
  stddev_pop?: Maybe<ProjectInOrganizationStddevPopFields>;
  stddev_samp?: Maybe<ProjectInOrganizationStddevSampFields>;
  sum?: Maybe<ProjectInOrganizationSumFields>;
  var_pop?: Maybe<ProjectInOrganizationVarPopFields>;
  var_samp?: Maybe<ProjectInOrganizationVarSampFields>;
  variance?: Maybe<ProjectInOrganizationVarianceFields>;
};


/** aggregate fields of "project_in_organization" */
export type ProjectInOrganizationAggregateFieldsCountArgs = {
  columns?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "project_in_organization" */
export type ProjectInOrganizationAggregateOrderBy = {
  avg?: Maybe<ProjectInOrganizationAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<ProjectInOrganizationMaxOrderBy>;
  min?: Maybe<ProjectInOrganizationMinOrderBy>;
  stddev?: Maybe<ProjectInOrganizationStddevOrderBy>;
  stddev_pop?: Maybe<ProjectInOrganizationStddevPopOrderBy>;
  stddev_samp?: Maybe<ProjectInOrganizationStddevSampOrderBy>;
  sum?: Maybe<ProjectInOrganizationSumOrderBy>;
  var_pop?: Maybe<ProjectInOrganizationVarPopOrderBy>;
  var_samp?: Maybe<ProjectInOrganizationVarSampOrderBy>;
  variance?: Maybe<ProjectInOrganizationVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "project_in_organization" */
export type ProjectInOrganizationArrRelInsertInput = {
  data: Array<ProjectInOrganizationInsertInput>;
  on_conflict?: Maybe<ProjectInOrganizationOnConflict>;
};

/** aggregate avg on columns */
export type ProjectInOrganizationAvgFields = {
   __typename?: 'project_in_organization_avg_fields';
  organization_id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "project_in_organization" */
export type ProjectInOrganizationAvgOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "project_in_organization". All fields are combined with a logical 'AND'. */
export type ProjectInOrganizationBoolExp = {
  _and?: Maybe<Array<Maybe<ProjectInOrganizationBoolExp>>>;
  _not?: Maybe<ProjectInOrganizationBoolExp>;
  _or?: Maybe<Array<Maybe<ProjectInOrganizationBoolExp>>>;
  organization?: Maybe<OrganizationBoolExp>;
  organization_id?: Maybe<BigintComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
};

/** unique or primary key constraints on table "project_in_organization" */
export enum ProjectInOrganizationConstraint {
  /** unique or primary key constraint */
  PROJECT_IN_ORGANIZATION_PKEY = 'project_in_organization_pkey'
}

/** input type for incrementing integer columne in table "project_in_organization" */
export type ProjectInOrganizationIncInput = {
  organization_id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "project_in_organization" */
export type ProjectInOrganizationInsertInput = {
  organization?: Maybe<OrganizationObjRelInsertInput>;
  organization_id?: Maybe<Scalars['bigint']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** aggregate max on columns */
export type ProjectInOrganizationMaxFields = {
   __typename?: 'project_in_organization_max_fields';
  organization_id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by max() on columns of table "project_in_organization" */
export type ProjectInOrganizationMaxOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type ProjectInOrganizationMinFields = {
   __typename?: 'project_in_organization_min_fields';
  organization_id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by min() on columns of table "project_in_organization" */
export type ProjectInOrganizationMinOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** response of any mutation on the table "project_in_organization" */
export type ProjectInOrganizationMutationResponse = {
   __typename?: 'project_in_organization_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<ProjectInOrganization>;
};

/** input type for inserting object relation for remote table "project_in_organization" */
export type ProjectInOrganizationObjRelInsertInput = {
  data: ProjectInOrganizationInsertInput;
  on_conflict?: Maybe<ProjectInOrganizationOnConflict>;
};

/** on conflict condition type for table "project_in_organization" */
export type ProjectInOrganizationOnConflict = {
  constraint: ProjectInOrganizationConstraint;
  update_columns: Array<ProjectInOrganizationUpdateColumn>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};

/** ordering options when selecting data from "project_in_organization" */
export type ProjectInOrganizationOrderBy = {
  organization?: Maybe<OrganizationOrderBy>;
  organization_id?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** primary key columns input for table: "project_in_organization" */
export type ProjectInOrganizationPkColumnsInput = {
  organization_id: Scalars['bigint'];
  project_id: Scalars['bigint'];
};

/** select columns of table "project_in_organization" */
export enum ProjectInOrganizationSelectColumn {
  /** column name */
  ORGANIZATION_ID = 'organization_id',
  /** column name */
  PROJECT_ID = 'project_id'
}

/** input type for updating data in table "project_in_organization" */
export type ProjectInOrganizationSetInput = {
  organization_id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** aggregate stddev on columns */
export type ProjectInOrganizationStddevFields = {
   __typename?: 'project_in_organization_stddev_fields';
  organization_id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "project_in_organization" */
export type ProjectInOrganizationStddevOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type ProjectInOrganizationStddevPopFields = {
   __typename?: 'project_in_organization_stddev_pop_fields';
  organization_id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "project_in_organization" */
export type ProjectInOrganizationStddevPopOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type ProjectInOrganizationStddevSampFields = {
   __typename?: 'project_in_organization_stddev_samp_fields';
  organization_id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "project_in_organization" */
export type ProjectInOrganizationStddevSampOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type ProjectInOrganizationSumFields = {
   __typename?: 'project_in_organization_sum_fields';
  organization_id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "project_in_organization" */
export type ProjectInOrganizationSumOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "project_in_organization" */
export enum ProjectInOrganizationUpdateColumn {
  /** column name */
  ORGANIZATION_ID = 'organization_id',
  /** column name */
  PROJECT_ID = 'project_id'
}

/** aggregate var_pop on columns */
export type ProjectInOrganizationVarPopFields = {
   __typename?: 'project_in_organization_var_pop_fields';
  organization_id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "project_in_organization" */
export type ProjectInOrganizationVarPopOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type ProjectInOrganizationVarSampFields = {
   __typename?: 'project_in_organization_var_samp_fields';
  organization_id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "project_in_organization" */
export type ProjectInOrganizationVarSampOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type ProjectInOrganizationVarianceFields = {
   __typename?: 'project_in_organization_variance_fields';
  organization_id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "project_in_organization" */
export type ProjectInOrganizationVarianceOrderBy = {
  organization_id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** input type for incrementing integer columne in table "project" */
export type ProjectIncInput = {
  id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "project" */
export type ProjectInsertInput = {
  aws_accounts?: Maybe<AwsAccountArrRelInsertInput>;
  created_at?: Maybe<Scalars['timestamp']>;
  description?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  instances?: Maybe<InstanceArrRelInsertInput>;
  name?: Maybe<Scalars['String']>;
  project_collaborators?: Maybe<ProjectCollaboratorArrRelInsertInput>;
  project_in_organizations?: Maybe<ProjectInOrganizationArrRelInsertInput>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_account?: Maybe<UserAccountObjRelInsertInput>;
  user_id?: Maybe<Scalars['String']>;
};

/** columns and relationships of "project_invitation" */
export type ProjectInvitation = {
   __typename?: 'project_invitation';
  confirmed: Scalars['Boolean'];
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  invitee_user_id: Scalars['String'];
  mail_address: Scalars['String'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  token: Scalars['String'];
  updated_at: Scalars['timestamp'];
};

/** aggregated selection of "project_invitation" */
export type ProjectInvitationAggregate = {
   __typename?: 'project_invitation_aggregate';
  aggregate?: Maybe<ProjectInvitationAggregateFields>;
  nodes: Array<ProjectInvitation>;
};

/** aggregate fields of "project_invitation" */
export type ProjectInvitationAggregateFields = {
   __typename?: 'project_invitation_aggregate_fields';
  avg?: Maybe<ProjectInvitationAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<ProjectInvitationMaxFields>;
  min?: Maybe<ProjectInvitationMinFields>;
  stddev?: Maybe<ProjectInvitationStddevFields>;
  stddev_pop?: Maybe<ProjectInvitationStddevPopFields>;
  stddev_samp?: Maybe<ProjectInvitationStddevSampFields>;
  sum?: Maybe<ProjectInvitationSumFields>;
  var_pop?: Maybe<ProjectInvitationVarPopFields>;
  var_samp?: Maybe<ProjectInvitationVarSampFields>;
  variance?: Maybe<ProjectInvitationVarianceFields>;
};


/** aggregate fields of "project_invitation" */
export type ProjectInvitationAggregateFieldsCountArgs = {
  columns?: Maybe<Array<ProjectInvitationSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "project_invitation" */
export type ProjectInvitationAggregateOrderBy = {
  avg?: Maybe<ProjectInvitationAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<ProjectInvitationMaxOrderBy>;
  min?: Maybe<ProjectInvitationMinOrderBy>;
  stddev?: Maybe<ProjectInvitationStddevOrderBy>;
  stddev_pop?: Maybe<ProjectInvitationStddevPopOrderBy>;
  stddev_samp?: Maybe<ProjectInvitationStddevSampOrderBy>;
  sum?: Maybe<ProjectInvitationSumOrderBy>;
  var_pop?: Maybe<ProjectInvitationVarPopOrderBy>;
  var_samp?: Maybe<ProjectInvitationVarSampOrderBy>;
  variance?: Maybe<ProjectInvitationVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "project_invitation" */
export type ProjectInvitationArrRelInsertInput = {
  data: Array<ProjectInvitationInsertInput>;
  on_conflict?: Maybe<ProjectInvitationOnConflict>;
};

/** aggregate avg on columns */
export type ProjectInvitationAvgFields = {
   __typename?: 'project_invitation_avg_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "project_invitation" */
export type ProjectInvitationAvgOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "project_invitation". All fields are combined with a logical 'AND'. */
export type ProjectInvitationBoolExp = {
  _and?: Maybe<Array<Maybe<ProjectInvitationBoolExp>>>;
  _not?: Maybe<ProjectInvitationBoolExp>;
  _or?: Maybe<Array<Maybe<ProjectInvitationBoolExp>>>;
  confirmed?: Maybe<BooleanComparisonExp>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  invitee_user_id?: Maybe<StringComparisonExp>;
  mail_address?: Maybe<StringComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  token?: Maybe<StringComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
};

/** unique or primary key constraints on table "project_invitation" */
export enum ProjectInvitationConstraint {
  /** unique or primary key constraint */
  PROJECT_INVITATION_PKEY = 'project_invitation_pkey',
  /** unique or primary key constraint */
  PROJECT_INVITATION_TOKEN_KEY = 'project_invitation_token_key'
}

/** input type for incrementing integer columne in table "project_invitation" */
export type ProjectInvitationIncInput = {
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "project_invitation" */
export type ProjectInvitationInsertInput = {
  confirmed?: Maybe<Scalars['Boolean']>;
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  invitee_user_id?: Maybe<Scalars['String']>;
  mail_address?: Maybe<Scalars['String']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  token?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate max on columns */
export type ProjectInvitationMaxFields = {
   __typename?: 'project_invitation_max_fields';
  id?: Maybe<Scalars['bigint']>;
  invitee_user_id?: Maybe<Scalars['String']>;
  mail_address?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  token?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "project_invitation" */
export type ProjectInvitationMaxOrderBy = {
  id?: Maybe<OrderBy>;
  invitee_user_id?: Maybe<OrderBy>;
  mail_address?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  token?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type ProjectInvitationMinFields = {
   __typename?: 'project_invitation_min_fields';
  id?: Maybe<Scalars['bigint']>;
  invitee_user_id?: Maybe<Scalars['String']>;
  mail_address?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  token?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "project_invitation" */
export type ProjectInvitationMinOrderBy = {
  id?: Maybe<OrderBy>;
  invitee_user_id?: Maybe<OrderBy>;
  mail_address?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  token?: Maybe<OrderBy>;
};

/** response of any mutation on the table "project_invitation" */
export type ProjectInvitationMutationResponse = {
   __typename?: 'project_invitation_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<ProjectInvitation>;
};

/** input type for inserting object relation for remote table "project_invitation" */
export type ProjectInvitationObjRelInsertInput = {
  data: ProjectInvitationInsertInput;
  on_conflict?: Maybe<ProjectInvitationOnConflict>;
};

/** on conflict condition type for table "project_invitation" */
export type ProjectInvitationOnConflict = {
  constraint: ProjectInvitationConstraint;
  update_columns: Array<ProjectInvitationUpdateColumn>;
  where?: Maybe<ProjectInvitationBoolExp>;
};

/** ordering options when selecting data from "project_invitation" */
export type ProjectInvitationOrderBy = {
  confirmed?: Maybe<OrderBy>;
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  invitee_user_id?: Maybe<OrderBy>;
  mail_address?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  token?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
};

/** primary key columns input for table: "project_invitation" */
export type ProjectInvitationPkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "project_invitation" */
export enum ProjectInvitationSelectColumn {
  /** column name */
  CONFIRMED = 'confirmed',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  INVITEE_USER_ID = 'invitee_user_id',
  /** column name */
  MAIL_ADDRESS = 'mail_address',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  TOKEN = 'token',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** input type for updating data in table "project_invitation" */
export type ProjectInvitationSetInput = {
  confirmed?: Maybe<Scalars['Boolean']>;
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  invitee_user_id?: Maybe<Scalars['String']>;
  mail_address?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  token?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate stddev on columns */
export type ProjectInvitationStddevFields = {
   __typename?: 'project_invitation_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "project_invitation" */
export type ProjectInvitationStddevOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type ProjectInvitationStddevPopFields = {
   __typename?: 'project_invitation_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "project_invitation" */
export type ProjectInvitationStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type ProjectInvitationStddevSampFields = {
   __typename?: 'project_invitation_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "project_invitation" */
export type ProjectInvitationStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type ProjectInvitationSumFields = {
   __typename?: 'project_invitation_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "project_invitation" */
export type ProjectInvitationSumOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "project_invitation" */
export enum ProjectInvitationUpdateColumn {
  /** column name */
  CONFIRMED = 'confirmed',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  INVITEE_USER_ID = 'invitee_user_id',
  /** column name */
  MAIL_ADDRESS = 'mail_address',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  TOKEN = 'token',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** aggregate var_pop on columns */
export type ProjectInvitationVarPopFields = {
   __typename?: 'project_invitation_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "project_invitation" */
export type ProjectInvitationVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type ProjectInvitationVarSampFields = {
   __typename?: 'project_invitation_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "project_invitation" */
export type ProjectInvitationVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type ProjectInvitationVarianceFields = {
   __typename?: 'project_invitation_variance_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "project_invitation" */
export type ProjectInvitationVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate max on columns */
export type ProjectMaxFields = {
   __typename?: 'project_max_fields';
  description?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "project" */
export type ProjectMaxOrderBy = {
  description?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type ProjectMinFields = {
   __typename?: 'project_min_fields';
  description?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  user_id?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "project" */
export type ProjectMinOrderBy = {
  description?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** response of any mutation on the table "project" */
export type ProjectMutationResponse = {
   __typename?: 'project_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Project>;
};

/** input type for inserting object relation for remote table "project" */
export type ProjectObjRelInsertInput = {
  data: ProjectInsertInput;
  on_conflict?: Maybe<ProjectOnConflict>;
};

/** on conflict condition type for table "project" */
export type ProjectOnConflict = {
  constraint: ProjectConstraint;
  update_columns: Array<ProjectUpdateColumn>;
  where?: Maybe<ProjectBoolExp>;
};

/** ordering options when selecting data from "project" */
export type ProjectOrderBy = {
  aws_accounts_aggregate?: Maybe<AwsAccountAggregateOrderBy>;
  created_at?: Maybe<OrderBy>;
  description?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  instances_aggregate?: Maybe<InstanceAggregateOrderBy>;
  name?: Maybe<OrderBy>;
  project_collaborators_aggregate?: Maybe<ProjectCollaboratorAggregateOrderBy>;
  project_in_organizations_aggregate?: Maybe<ProjectInOrganizationAggregateOrderBy>;
  updated_at?: Maybe<OrderBy>;
  user_account?: Maybe<UserAccountOrderBy>;
  user_id?: Maybe<OrderBy>;
};

/** primary key columns input for table: "project" */
export type ProjectPkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "project" */
export enum ProjectSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  DESCRIPTION = 'description',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id'
}

/** input type for updating data in table "project" */
export type ProjectSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  description?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type ProjectStddevFields = {
   __typename?: 'project_stddev_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "project" */
export type ProjectStddevOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type ProjectStddevPopFields = {
   __typename?: 'project_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "project" */
export type ProjectStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type ProjectStddevSampFields = {
   __typename?: 'project_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "project" */
export type ProjectStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type ProjectSumFields = {
   __typename?: 'project_sum_fields';
  id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "project" */
export type ProjectSumOrderBy = {
  id?: Maybe<OrderBy>;
};

/** update columns of table "project" */
export enum ProjectUpdateColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  DESCRIPTION = 'description',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id'
}

/** aggregate var_pop on columns */
export type ProjectVarPopFields = {
   __typename?: 'project_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "project" */
export type ProjectVarPopOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type ProjectVarSampFields = {
   __typename?: 'project_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "project" */
export type ProjectVarSampOrderBy = {
  id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type ProjectVarianceFields = {
   __typename?: 'project_variance_fields';
  id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "project" */
export type ProjectVarianceOrderBy = {
  id?: Maybe<OrderBy>;
};

/** query root */
export type QueryRoot = {
   __typename?: 'query_root';
  /** fetch data from the table: "action" */
  action: Array<Action>;
  /** fetch aggregated fields from the table: "action" */
  action_aggregate: ActionAggregate;
  /** fetch data from the table: "action" using primary key columns */
  action_by_pk?: Maybe<Action>;
  /** fetch data from the table: "aws_account" */
  aws_account: Array<AwsAccount>;
  /** fetch aggregated fields from the table: "aws_account" */
  aws_account_aggregate: AwsAccountAggregate;
  /** fetch data from the table: "aws_account" using primary key columns */
  aws_account_by_pk?: Maybe<AwsAccount>;
  aws_instance?: Maybe<AwsInstance>;
  aws_instances: Array<AwsInstance>;
  /** fetch data from the table: "instance" */
  instance: Array<Instance>;
  /** fetch aggregated fields from the table: "instance" */
  instance_aggregate: InstanceAggregate;
  /** fetch data from the table: "instance_at_service" */
  instance_at_service: Array<InstanceAtService>;
  /** fetch aggregated fields from the table: "instance_at_service" */
  instance_at_service_aggregate: InstanceAtServiceAggregate;
  /** fetch data from the table: "instance_at_service" using primary key columns */
  instance_at_service_by_pk?: Maybe<InstanceAtService>;
  /** fetch data from the table: "instance" using primary key columns */
  instance_by_pk?: Maybe<Instance>;
  /** fetch data from the table: "metrics" */
  metrics: Array<Metrics>;
  /** fetch aggregated fields from the table: "metrics" */
  metrics_aggregate: MetricsAggregate;
  /** fetch data from the table: "notification_rule" */
  notification_rule: Array<NotificationRule>;
  /** fetch aggregated fields from the table: "notification_rule" */
  notification_rule_aggregate: NotificationRuleAggregate;
  /** fetch data from the table: "notification_rule" using primary key columns */
  notification_rule_by_pk?: Maybe<NotificationRule>;
  /** fetch data from the table: "organization" */
  organization: Array<Organization>;
  /** fetch aggregated fields from the table: "organization" */
  organization_aggregate: OrganizationAggregate;
  /** fetch data from the table: "organization" using primary key columns */
  organization_by_pk?: Maybe<Organization>;
  /** fetch data from the table: "project" */
  project: Array<Project>;
  /** fetch aggregated fields from the table: "project" */
  project_aggregate: ProjectAggregate;
  /** fetch data from the table: "project" using primary key columns */
  project_by_pk?: Maybe<Project>;
  /** fetch data from the table: "project_collaborator" */
  project_collaborator: Array<ProjectCollaborator>;
  /** fetch aggregated fields from the table: "project_collaborator" */
  project_collaborator_aggregate: ProjectCollaboratorAggregate;
  /** fetch data from the table: "project_in_organization" */
  project_in_organization: Array<ProjectInOrganization>;
  /** fetch aggregated fields from the table: "project_in_organization" */
  project_in_organization_aggregate: ProjectInOrganizationAggregate;
  /** fetch data from the table: "project_in_organization" using primary key columns */
  project_in_organization_by_pk?: Maybe<ProjectInOrganization>;
  /** fetch data from the table: "project_invitation" */
  project_invitation: Array<ProjectInvitation>;
  /** fetch aggregated fields from the table: "project_invitation" */
  project_invitation_aggregate: ProjectInvitationAggregate;
  /** fetch data from the table: "project_invitation" using primary key columns */
  project_invitation_by_pk?: Maybe<ProjectInvitation>;
  /** fetch data from the table: "service" */
  service: Array<Service>;
  /** fetch aggregated fields from the table: "service" */
  service_aggregate: ServiceAggregate;
  /** fetch data from the table: "service" using primary key columns */
  service_by_pk?: Maybe<Service>;
  /** fetch data from the table: "slack_webhook" */
  slack_webhook: Array<SlackWebhook>;
  /** fetch aggregated fields from the table: "slack_webhook" */
  slack_webhook_aggregate: SlackWebhookAggregate;
  /** fetch data from the table: "slack_webhook" using primary key columns */
  slack_webhook_by_pk?: Maybe<SlackWebhook>;
  /** fetch data from the table: "user_account" */
  user_account: Array<UserAccount>;
  /** fetch aggregated fields from the table: "user_account" */
  user_account_aggregate: UserAccountAggregate;
  /** fetch data from the table: "user_account" using primary key columns */
  user_account_by_pk?: Maybe<UserAccount>;
};


/** query root */
export type QueryRootActionArgs = {
  distinct_on?: Maybe<Array<ActionSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ActionOrderBy>>;
  where?: Maybe<ActionBoolExp>;
};


/** query root */
export type QueryRootActionAggregateArgs = {
  distinct_on?: Maybe<Array<ActionSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ActionOrderBy>>;
  where?: Maybe<ActionBoolExp>;
};


/** query root */
export type QueryRootActionByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootAwsAccountArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** query root */
export type QueryRootAwsAccountAggregateArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** query root */
export type QueryRootAwsAccountByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootAwsInstanceArgs = {
  id: Scalars['ID'];
};


/** query root */
export type QueryRootAwsInstancesArgs = {
  projectId: Scalars['ID'];
};


/** query root */
export type QueryRootInstanceArgs = {
  distinct_on?: Maybe<Array<InstanceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceOrderBy>>;
  where?: Maybe<InstanceBoolExp>;
};


/** query root */
export type QueryRootInstanceAggregateArgs = {
  distinct_on?: Maybe<Array<InstanceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceOrderBy>>;
  where?: Maybe<InstanceBoolExp>;
};


/** query root */
export type QueryRootInstanceAtServiceArgs = {
  distinct_on?: Maybe<Array<InstanceAtServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceAtServiceOrderBy>>;
  where?: Maybe<InstanceAtServiceBoolExp>;
};


/** query root */
export type QueryRootInstanceAtServiceAggregateArgs = {
  distinct_on?: Maybe<Array<InstanceAtServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceAtServiceOrderBy>>;
  where?: Maybe<InstanceAtServiceBoolExp>;
};


/** query root */
export type QueryRootInstanceAtServiceByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootInstanceByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootMetricsArgs = {
  distinct_on?: Maybe<Array<MetricsSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<MetricsOrderBy>>;
  where?: Maybe<MetricsBoolExp>;
};


/** query root */
export type QueryRootMetricsAggregateArgs = {
  distinct_on?: Maybe<Array<MetricsSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<MetricsOrderBy>>;
  where?: Maybe<MetricsBoolExp>;
};


/** query root */
export type QueryRootNotificationRuleArgs = {
  distinct_on?: Maybe<Array<NotificationRuleSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<NotificationRuleOrderBy>>;
  where?: Maybe<NotificationRuleBoolExp>;
};


/** query root */
export type QueryRootNotificationRuleAggregateArgs = {
  distinct_on?: Maybe<Array<NotificationRuleSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<NotificationRuleOrderBy>>;
  where?: Maybe<NotificationRuleBoolExp>;
};


/** query root */
export type QueryRootNotificationRuleByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootOrganizationArgs = {
  distinct_on?: Maybe<Array<OrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<OrganizationOrderBy>>;
  where?: Maybe<OrganizationBoolExp>;
};


/** query root */
export type QueryRootOrganizationAggregateArgs = {
  distinct_on?: Maybe<Array<OrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<OrganizationOrderBy>>;
  where?: Maybe<OrganizationBoolExp>;
};


/** query root */
export type QueryRootOrganizationByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootProjectArgs = {
  distinct_on?: Maybe<Array<ProjectSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectOrderBy>>;
  where?: Maybe<ProjectBoolExp>;
};


/** query root */
export type QueryRootProjectAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectOrderBy>>;
  where?: Maybe<ProjectBoolExp>;
};


/** query root */
export type QueryRootProjectByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootProjectCollaboratorArgs = {
  distinct_on?: Maybe<Array<ProjectCollaboratorSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectCollaboratorOrderBy>>;
  where?: Maybe<ProjectCollaboratorBoolExp>;
};


/** query root */
export type QueryRootProjectCollaboratorAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectCollaboratorSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectCollaboratorOrderBy>>;
  where?: Maybe<ProjectCollaboratorBoolExp>;
};


/** query root */
export type QueryRootProjectInOrganizationArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};


/** query root */
export type QueryRootProjectInOrganizationAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};


/** query root */
export type QueryRootProjectInOrganizationByPkArgs = {
  organization_id: Scalars['bigint'];
  project_id: Scalars['bigint'];
};


/** query root */
export type QueryRootProjectInvitationArgs = {
  distinct_on?: Maybe<Array<ProjectInvitationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInvitationOrderBy>>;
  where?: Maybe<ProjectInvitationBoolExp>;
};


/** query root */
export type QueryRootProjectInvitationAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectInvitationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInvitationOrderBy>>;
  where?: Maybe<ProjectInvitationBoolExp>;
};


/** query root */
export type QueryRootProjectInvitationByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootServiceArgs = {
  distinct_on?: Maybe<Array<ServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ServiceOrderBy>>;
  where?: Maybe<ServiceBoolExp>;
};


/** query root */
export type QueryRootServiceAggregateArgs = {
  distinct_on?: Maybe<Array<ServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ServiceOrderBy>>;
  where?: Maybe<ServiceBoolExp>;
};


/** query root */
export type QueryRootServiceByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootSlackWebhookArgs = {
  distinct_on?: Maybe<Array<SlackWebhookSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<SlackWebhookOrderBy>>;
  where?: Maybe<SlackWebhookBoolExp>;
};


/** query root */
export type QueryRootSlackWebhookAggregateArgs = {
  distinct_on?: Maybe<Array<SlackWebhookSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<SlackWebhookOrderBy>>;
  where?: Maybe<SlackWebhookBoolExp>;
};


/** query root */
export type QueryRootSlackWebhookByPkArgs = {
  id: Scalars['bigint'];
};


/** query root */
export type QueryRootUserAccountArgs = {
  distinct_on?: Maybe<Array<UserAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<UserAccountOrderBy>>;
  where?: Maybe<UserAccountBoolExp>;
};


/** query root */
export type QueryRootUserAccountAggregateArgs = {
  distinct_on?: Maybe<Array<UserAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<UserAccountOrderBy>>;
  where?: Maybe<UserAccountBoolExp>;
};


/** query root */
export type QueryRootUserAccountByPkArgs = {
  id: Scalars['String'];
};

/** columns and relationships of "service" */
export type Service = {
   __typename?: 'service';
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  name: Scalars['String'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  updated_at: Scalars['timestamp'];
};

/** aggregated selection of "service" */
export type ServiceAggregate = {
   __typename?: 'service_aggregate';
  aggregate?: Maybe<ServiceAggregateFields>;
  nodes: Array<Service>;
};

/** aggregate fields of "service" */
export type ServiceAggregateFields = {
   __typename?: 'service_aggregate_fields';
  avg?: Maybe<ServiceAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<ServiceMaxFields>;
  min?: Maybe<ServiceMinFields>;
  stddev?: Maybe<ServiceStddevFields>;
  stddev_pop?: Maybe<ServiceStddevPopFields>;
  stddev_samp?: Maybe<ServiceStddevSampFields>;
  sum?: Maybe<ServiceSumFields>;
  var_pop?: Maybe<ServiceVarPopFields>;
  var_samp?: Maybe<ServiceVarSampFields>;
  variance?: Maybe<ServiceVarianceFields>;
};


/** aggregate fields of "service" */
export type ServiceAggregateFieldsCountArgs = {
  columns?: Maybe<Array<ServiceSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "service" */
export type ServiceAggregateOrderBy = {
  avg?: Maybe<ServiceAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<ServiceMaxOrderBy>;
  min?: Maybe<ServiceMinOrderBy>;
  stddev?: Maybe<ServiceStddevOrderBy>;
  stddev_pop?: Maybe<ServiceStddevPopOrderBy>;
  stddev_samp?: Maybe<ServiceStddevSampOrderBy>;
  sum?: Maybe<ServiceSumOrderBy>;
  var_pop?: Maybe<ServiceVarPopOrderBy>;
  var_samp?: Maybe<ServiceVarSampOrderBy>;
  variance?: Maybe<ServiceVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "service" */
export type ServiceArrRelInsertInput = {
  data: Array<ServiceInsertInput>;
  on_conflict?: Maybe<ServiceOnConflict>;
};

/** aggregate avg on columns */
export type ServiceAvgFields = {
   __typename?: 'service_avg_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "service" */
export type ServiceAvgOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "service". All fields are combined with a logical 'AND'. */
export type ServiceBoolExp = {
  _and?: Maybe<Array<Maybe<ServiceBoolExp>>>;
  _not?: Maybe<ServiceBoolExp>;
  _or?: Maybe<Array<Maybe<ServiceBoolExp>>>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  name?: Maybe<StringComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
};

/** unique or primary key constraints on table "service" */
export enum ServiceConstraint {
  /** unique or primary key constraint */
  SERVICE_PKEY = 'service_pkey',
  /** unique or primary key constraint */
  SERVICE_PROJECT_ID_NAME_KEY = 'service_project_id_name_key'
}

/** input type for incrementing integer columne in table "service" */
export type ServiceIncInput = {
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "service" */
export type ServiceInsertInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate max on columns */
export type ServiceMaxFields = {
   __typename?: 'service_max_fields';
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by max() on columns of table "service" */
export type ServiceMaxOrderBy = {
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type ServiceMinFields = {
   __typename?: 'service_min_fields';
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by min() on columns of table "service" */
export type ServiceMinOrderBy = {
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** response of any mutation on the table "service" */
export type ServiceMutationResponse = {
   __typename?: 'service_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<Service>;
};

/** input type for inserting object relation for remote table "service" */
export type ServiceObjRelInsertInput = {
  data: ServiceInsertInput;
  on_conflict?: Maybe<ServiceOnConflict>;
};

/** on conflict condition type for table "service" */
export type ServiceOnConflict = {
  constraint: ServiceConstraint;
  update_columns: Array<ServiceUpdateColumn>;
  where?: Maybe<ServiceBoolExp>;
};

/** ordering options when selecting data from "service" */
export type ServiceOrderBy = {
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
};

/** primary key columns input for table: "service" */
export type ServicePkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "service" */
export enum ServiceSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** input type for updating data in table "service" */
export type ServiceSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
};

/** aggregate stddev on columns */
export type ServiceStddevFields = {
   __typename?: 'service_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "service" */
export type ServiceStddevOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type ServiceStddevPopFields = {
   __typename?: 'service_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "service" */
export type ServiceStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type ServiceStddevSampFields = {
   __typename?: 'service_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "service" */
export type ServiceStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type ServiceSumFields = {
   __typename?: 'service_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "service" */
export type ServiceSumOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "service" */
export enum ServiceUpdateColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  UPDATED_AT = 'updated_at'
}

/** aggregate var_pop on columns */
export type ServiceVarPopFields = {
   __typename?: 'service_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "service" */
export type ServiceVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type ServiceVarSampFields = {
   __typename?: 'service_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "service" */
export type ServiceVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type ServiceVarianceFields = {
   __typename?: 'service_variance_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "service" */
export type ServiceVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** columns and relationships of "slack_webhook" */
export type SlackWebhook = {
   __typename?: 'slack_webhook';
  channel: Scalars['String'];
  created_at: Scalars['timestamp'];
  id: Scalars['bigint'];
  name: Scalars['String'];
  /** An object relationship */
  project: Project;
  project_id: Scalars['bigint'];
  updated_at: Scalars['timestamp'];
  user_id: Scalars['String'];
  webhook_url: Scalars['String'];
};

/** aggregated selection of "slack_webhook" */
export type SlackWebhookAggregate = {
   __typename?: 'slack_webhook_aggregate';
  aggregate?: Maybe<SlackWebhookAggregateFields>;
  nodes: Array<SlackWebhook>;
};

/** aggregate fields of "slack_webhook" */
export type SlackWebhookAggregateFields = {
   __typename?: 'slack_webhook_aggregate_fields';
  avg?: Maybe<SlackWebhookAvgFields>;
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<SlackWebhookMaxFields>;
  min?: Maybe<SlackWebhookMinFields>;
  stddev?: Maybe<SlackWebhookStddevFields>;
  stddev_pop?: Maybe<SlackWebhookStddevPopFields>;
  stddev_samp?: Maybe<SlackWebhookStddevSampFields>;
  sum?: Maybe<SlackWebhookSumFields>;
  var_pop?: Maybe<SlackWebhookVarPopFields>;
  var_samp?: Maybe<SlackWebhookVarSampFields>;
  variance?: Maybe<SlackWebhookVarianceFields>;
};


/** aggregate fields of "slack_webhook" */
export type SlackWebhookAggregateFieldsCountArgs = {
  columns?: Maybe<Array<SlackWebhookSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "slack_webhook" */
export type SlackWebhookAggregateOrderBy = {
  avg?: Maybe<SlackWebhookAvgOrderBy>;
  count?: Maybe<OrderBy>;
  max?: Maybe<SlackWebhookMaxOrderBy>;
  min?: Maybe<SlackWebhookMinOrderBy>;
  stddev?: Maybe<SlackWebhookStddevOrderBy>;
  stddev_pop?: Maybe<SlackWebhookStddevPopOrderBy>;
  stddev_samp?: Maybe<SlackWebhookStddevSampOrderBy>;
  sum?: Maybe<SlackWebhookSumOrderBy>;
  var_pop?: Maybe<SlackWebhookVarPopOrderBy>;
  var_samp?: Maybe<SlackWebhookVarSampOrderBy>;
  variance?: Maybe<SlackWebhookVarianceOrderBy>;
};

/** input type for inserting array relation for remote table "slack_webhook" */
export type SlackWebhookArrRelInsertInput = {
  data: Array<SlackWebhookInsertInput>;
  on_conflict?: Maybe<SlackWebhookOnConflict>;
};

/** aggregate avg on columns */
export type SlackWebhookAvgFields = {
   __typename?: 'slack_webhook_avg_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by avg() on columns of table "slack_webhook" */
export type SlackWebhookAvgOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** Boolean expression to filter rows from the table "slack_webhook". All fields are combined with a logical 'AND'. */
export type SlackWebhookBoolExp = {
  _and?: Maybe<Array<Maybe<SlackWebhookBoolExp>>>;
  _not?: Maybe<SlackWebhookBoolExp>;
  _or?: Maybe<Array<Maybe<SlackWebhookBoolExp>>>;
  channel?: Maybe<StringComparisonExp>;
  created_at?: Maybe<TimestampComparisonExp>;
  id?: Maybe<BigintComparisonExp>;
  name?: Maybe<StringComparisonExp>;
  project?: Maybe<ProjectBoolExp>;
  project_id?: Maybe<BigintComparisonExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
  user_id?: Maybe<StringComparisonExp>;
  webhook_url?: Maybe<StringComparisonExp>;
};

/** unique or primary key constraints on table "slack_webhook" */
export enum SlackWebhookConstraint {
  /** unique or primary key constraint */
  SLACK_WEBHOOK_PKEY = 'slack_webhook_pkey',
  /** unique or primary key constraint */
  SLACK_WEBHOOK_PROJECT_ID_CHANNEL_KEY = 'slack_webhook_project_id_channel_key'
}

/** input type for incrementing integer columne in table "slack_webhook" */
export type SlackWebhookIncInput = {
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** input type for inserting data into table "slack_webhook" */
export type SlackWebhookInsertInput = {
  channel?: Maybe<Scalars['String']>;
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project?: Maybe<ProjectObjRelInsertInput>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
  webhook_url?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type SlackWebhookMaxFields = {
   __typename?: 'slack_webhook_max_fields';
  channel?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  user_id?: Maybe<Scalars['String']>;
  webhook_url?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "slack_webhook" */
export type SlackWebhookMaxOrderBy = {
  channel?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
  webhook_url?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type SlackWebhookMinFields = {
   __typename?: 'slack_webhook_min_fields';
  channel?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  user_id?: Maybe<Scalars['String']>;
  webhook_url?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "slack_webhook" */
export type SlackWebhookMinOrderBy = {
  channel?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
  webhook_url?: Maybe<OrderBy>;
};

/** response of any mutation on the table "slack_webhook" */
export type SlackWebhookMutationResponse = {
   __typename?: 'slack_webhook_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<SlackWebhook>;
};

/** input type for inserting object relation for remote table "slack_webhook" */
export type SlackWebhookObjRelInsertInput = {
  data: SlackWebhookInsertInput;
  on_conflict?: Maybe<SlackWebhookOnConflict>;
};

/** on conflict condition type for table "slack_webhook" */
export type SlackWebhookOnConflict = {
  constraint: SlackWebhookConstraint;
  update_columns: Array<SlackWebhookUpdateColumn>;
  where?: Maybe<SlackWebhookBoolExp>;
};

/** ordering options when selecting data from "slack_webhook" */
export type SlackWebhookOrderBy = {
  channel?: Maybe<OrderBy>;
  created_at?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  name?: Maybe<OrderBy>;
  project?: Maybe<ProjectOrderBy>;
  project_id?: Maybe<OrderBy>;
  updated_at?: Maybe<OrderBy>;
  user_id?: Maybe<OrderBy>;
  webhook_url?: Maybe<OrderBy>;
};

/** primary key columns input for table: "slack_webhook" */
export type SlackWebhookPkColumnsInput = {
  id: Scalars['bigint'];
};

/** select columns of table "slack_webhook" */
export enum SlackWebhookSelectColumn {
  /** column name */
  CHANNEL = 'channel',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id',
  /** column name */
  WEBHOOK_URL = 'webhook_url'
}

/** input type for updating data in table "slack_webhook" */
export type SlackWebhookSetInput = {
  channel?: Maybe<Scalars['String']>;
  created_at?: Maybe<Scalars['timestamp']>;
  id?: Maybe<Scalars['bigint']>;
  name?: Maybe<Scalars['String']>;
  project_id?: Maybe<Scalars['bigint']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  user_id?: Maybe<Scalars['String']>;
  webhook_url?: Maybe<Scalars['String']>;
};

/** aggregate stddev on columns */
export type SlackWebhookStddevFields = {
   __typename?: 'slack_webhook_stddev_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev() on columns of table "slack_webhook" */
export type SlackWebhookStddevOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_pop on columns */
export type SlackWebhookStddevPopFields = {
   __typename?: 'slack_webhook_stddev_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_pop() on columns of table "slack_webhook" */
export type SlackWebhookStddevPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate stddev_samp on columns */
export type SlackWebhookStddevSampFields = {
   __typename?: 'slack_webhook_stddev_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by stddev_samp() on columns of table "slack_webhook" */
export type SlackWebhookStddevSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate sum on columns */
export type SlackWebhookSumFields = {
   __typename?: 'slack_webhook_sum_fields';
  id?: Maybe<Scalars['bigint']>;
  project_id?: Maybe<Scalars['bigint']>;
};

/** order by sum() on columns of table "slack_webhook" */
export type SlackWebhookSumOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** update columns of table "slack_webhook" */
export enum SlackWebhookUpdateColumn {
  /** column name */
  CHANNEL = 'channel',
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  ID = 'id',
  /** column name */
  NAME = 'name',
  /** column name */
  PROJECT_ID = 'project_id',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USER_ID = 'user_id',
  /** column name */
  WEBHOOK_URL = 'webhook_url'
}

/** aggregate var_pop on columns */
export type SlackWebhookVarPopFields = {
   __typename?: 'slack_webhook_var_pop_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_pop() on columns of table "slack_webhook" */
export type SlackWebhookVarPopOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate var_samp on columns */
export type SlackWebhookVarSampFields = {
   __typename?: 'slack_webhook_var_samp_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by var_samp() on columns of table "slack_webhook" */
export type SlackWebhookVarSampOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** aggregate variance on columns */
export type SlackWebhookVarianceFields = {
   __typename?: 'slack_webhook_variance_fields';
  id?: Maybe<Scalars['Float']>;
  project_id?: Maybe<Scalars['Float']>;
};

/** order by variance() on columns of table "slack_webhook" */
export type SlackWebhookVarianceOrderBy = {
  id?: Maybe<OrderBy>;
  project_id?: Maybe<OrderBy>;
};

/** subscription root */
export type SubscriptionRoot = {
   __typename?: 'subscription_root';
  /** fetch data from the table: "action" */
  action: Array<Action>;
  /** fetch aggregated fields from the table: "action" */
  action_aggregate: ActionAggregate;
  /** fetch data from the table: "action" using primary key columns */
  action_by_pk?: Maybe<Action>;
  /** fetch data from the table: "aws_account" */
  aws_account: Array<AwsAccount>;
  /** fetch aggregated fields from the table: "aws_account" */
  aws_account_aggregate: AwsAccountAggregate;
  /** fetch data from the table: "aws_account" using primary key columns */
  aws_account_by_pk?: Maybe<AwsAccount>;
  /** fetch data from the table: "instance" */
  instance: Array<Instance>;
  /** fetch aggregated fields from the table: "instance" */
  instance_aggregate: InstanceAggregate;
  /** fetch data from the table: "instance_at_service" */
  instance_at_service: Array<InstanceAtService>;
  /** fetch aggregated fields from the table: "instance_at_service" */
  instance_at_service_aggregate: InstanceAtServiceAggregate;
  /** fetch data from the table: "instance_at_service" using primary key columns */
  instance_at_service_by_pk?: Maybe<InstanceAtService>;
  /** fetch data from the table: "instance" using primary key columns */
  instance_by_pk?: Maybe<Instance>;
  /** fetch data from the table: "notification_rule" */
  notification_rule: Array<NotificationRule>;
  /** fetch aggregated fields from the table: "notification_rule" */
  notification_rule_aggregate: NotificationRuleAggregate;
  /** fetch data from the table: "notification_rule" using primary key columns */
  notification_rule_by_pk?: Maybe<NotificationRule>;
  /** fetch data from the table: "organization" */
  organization: Array<Organization>;
  /** fetch aggregated fields from the table: "organization" */
  organization_aggregate: OrganizationAggregate;
  /** fetch data from the table: "organization" using primary key columns */
  organization_by_pk?: Maybe<Organization>;
  /** fetch data from the table: "project" */
  project: Array<Project>;
  /** fetch aggregated fields from the table: "project" */
  project_aggregate: ProjectAggregate;
  /** fetch data from the table: "project" using primary key columns */
  project_by_pk?: Maybe<Project>;
  /** fetch data from the table: "project_collaborator" */
  project_collaborator: Array<ProjectCollaborator>;
  /** fetch aggregated fields from the table: "project_collaborator" */
  project_collaborator_aggregate: ProjectCollaboratorAggregate;
  /** fetch data from the table: "project_in_organization" */
  project_in_organization: Array<ProjectInOrganization>;
  /** fetch aggregated fields from the table: "project_in_organization" */
  project_in_organization_aggregate: ProjectInOrganizationAggregate;
  /** fetch data from the table: "project_in_organization" using primary key columns */
  project_in_organization_by_pk?: Maybe<ProjectInOrganization>;
  /** fetch data from the table: "project_invitation" */
  project_invitation: Array<ProjectInvitation>;
  /** fetch aggregated fields from the table: "project_invitation" */
  project_invitation_aggregate: ProjectInvitationAggregate;
  /** fetch data from the table: "project_invitation" using primary key columns */
  project_invitation_by_pk?: Maybe<ProjectInvitation>;
  /** fetch data from the table: "service" */
  service: Array<Service>;
  /** fetch aggregated fields from the table: "service" */
  service_aggregate: ServiceAggregate;
  /** fetch data from the table: "service" using primary key columns */
  service_by_pk?: Maybe<Service>;
  /** fetch data from the table: "slack_webhook" */
  slack_webhook: Array<SlackWebhook>;
  /** fetch aggregated fields from the table: "slack_webhook" */
  slack_webhook_aggregate: SlackWebhookAggregate;
  /** fetch data from the table: "slack_webhook" using primary key columns */
  slack_webhook_by_pk?: Maybe<SlackWebhook>;
  /** fetch data from the table: "user_account" */
  user_account: Array<UserAccount>;
  /** fetch aggregated fields from the table: "user_account" */
  user_account_aggregate: UserAccountAggregate;
  /** fetch data from the table: "user_account" using primary key columns */
  user_account_by_pk?: Maybe<UserAccount>;
};


/** subscription root */
export type SubscriptionRootActionArgs = {
  distinct_on?: Maybe<Array<ActionSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ActionOrderBy>>;
  where?: Maybe<ActionBoolExp>;
};


/** subscription root */
export type SubscriptionRootActionAggregateArgs = {
  distinct_on?: Maybe<Array<ActionSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ActionOrderBy>>;
  where?: Maybe<ActionBoolExp>;
};


/** subscription root */
export type SubscriptionRootActionByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootAwsAccountArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** subscription root */
export type SubscriptionRootAwsAccountAggregateArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** subscription root */
export type SubscriptionRootAwsAccountByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootInstanceArgs = {
  distinct_on?: Maybe<Array<InstanceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceOrderBy>>;
  where?: Maybe<InstanceBoolExp>;
};


/** subscription root */
export type SubscriptionRootInstanceAggregateArgs = {
  distinct_on?: Maybe<Array<InstanceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceOrderBy>>;
  where?: Maybe<InstanceBoolExp>;
};


/** subscription root */
export type SubscriptionRootInstanceAtServiceArgs = {
  distinct_on?: Maybe<Array<InstanceAtServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceAtServiceOrderBy>>;
  where?: Maybe<InstanceAtServiceBoolExp>;
};


/** subscription root */
export type SubscriptionRootInstanceAtServiceAggregateArgs = {
  distinct_on?: Maybe<Array<InstanceAtServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<InstanceAtServiceOrderBy>>;
  where?: Maybe<InstanceAtServiceBoolExp>;
};


/** subscription root */
export type SubscriptionRootInstanceAtServiceByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootInstanceByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootNotificationRuleArgs = {
  distinct_on?: Maybe<Array<NotificationRuleSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<NotificationRuleOrderBy>>;
  where?: Maybe<NotificationRuleBoolExp>;
};


/** subscription root */
export type SubscriptionRootNotificationRuleAggregateArgs = {
  distinct_on?: Maybe<Array<NotificationRuleSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<NotificationRuleOrderBy>>;
  where?: Maybe<NotificationRuleBoolExp>;
};


/** subscription root */
export type SubscriptionRootNotificationRuleByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootOrganizationArgs = {
  distinct_on?: Maybe<Array<OrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<OrganizationOrderBy>>;
  where?: Maybe<OrganizationBoolExp>;
};


/** subscription root */
export type SubscriptionRootOrganizationAggregateArgs = {
  distinct_on?: Maybe<Array<OrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<OrganizationOrderBy>>;
  where?: Maybe<OrganizationBoolExp>;
};


/** subscription root */
export type SubscriptionRootOrganizationByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootProjectArgs = {
  distinct_on?: Maybe<Array<ProjectSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectOrderBy>>;
  where?: Maybe<ProjectBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectOrderBy>>;
  where?: Maybe<ProjectBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootProjectCollaboratorArgs = {
  distinct_on?: Maybe<Array<ProjectCollaboratorSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectCollaboratorOrderBy>>;
  where?: Maybe<ProjectCollaboratorBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectCollaboratorAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectCollaboratorSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectCollaboratorOrderBy>>;
  where?: Maybe<ProjectCollaboratorBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectInOrganizationArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectInOrganizationAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectInOrganizationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInOrganizationOrderBy>>;
  where?: Maybe<ProjectInOrganizationBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectInOrganizationByPkArgs = {
  organization_id: Scalars['bigint'];
  project_id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootProjectInvitationArgs = {
  distinct_on?: Maybe<Array<ProjectInvitationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInvitationOrderBy>>;
  where?: Maybe<ProjectInvitationBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectInvitationAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectInvitationSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectInvitationOrderBy>>;
  where?: Maybe<ProjectInvitationBoolExp>;
};


/** subscription root */
export type SubscriptionRootProjectInvitationByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootServiceArgs = {
  distinct_on?: Maybe<Array<ServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ServiceOrderBy>>;
  where?: Maybe<ServiceBoolExp>;
};


/** subscription root */
export type SubscriptionRootServiceAggregateArgs = {
  distinct_on?: Maybe<Array<ServiceSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ServiceOrderBy>>;
  where?: Maybe<ServiceBoolExp>;
};


/** subscription root */
export type SubscriptionRootServiceByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootSlackWebhookArgs = {
  distinct_on?: Maybe<Array<SlackWebhookSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<SlackWebhookOrderBy>>;
  where?: Maybe<SlackWebhookBoolExp>;
};


/** subscription root */
export type SubscriptionRootSlackWebhookAggregateArgs = {
  distinct_on?: Maybe<Array<SlackWebhookSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<SlackWebhookOrderBy>>;
  where?: Maybe<SlackWebhookBoolExp>;
};


/** subscription root */
export type SubscriptionRootSlackWebhookByPkArgs = {
  id: Scalars['bigint'];
};


/** subscription root */
export type SubscriptionRootUserAccountArgs = {
  distinct_on?: Maybe<Array<UserAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<UserAccountOrderBy>>;
  where?: Maybe<UserAccountBoolExp>;
};


/** subscription root */
export type SubscriptionRootUserAccountAggregateArgs = {
  distinct_on?: Maybe<Array<UserAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<UserAccountOrderBy>>;
  where?: Maybe<UserAccountBoolExp>;
};


/** subscription root */
export type SubscriptionRootUserAccountByPkArgs = {
  id: Scalars['String'];
};


/** expression to compare columns of type timestamp. All fields are combined with logical 'AND'. */
export type TimestampComparisonExp = {
  _eq?: Maybe<Scalars['timestamp']>;
  _gt?: Maybe<Scalars['timestamp']>;
  _gte?: Maybe<Scalars['timestamp']>;
  _in?: Maybe<Array<Scalars['timestamp']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['timestamp']>;
  _lte?: Maybe<Scalars['timestamp']>;
  _neq?: Maybe<Scalars['timestamp']>;
  _nin?: Maybe<Array<Scalars['timestamp']>>;
};


/** expression to compare columns of type timestamptz. All fields are combined with logical 'AND'. */
export type TimestamptzComparisonExp = {
  _eq?: Maybe<Scalars['timestamptz']>;
  _gt?: Maybe<Scalars['timestamptz']>;
  _gte?: Maybe<Scalars['timestamptz']>;
  _in?: Maybe<Array<Scalars['timestamptz']>>;
  _is_null?: Maybe<Scalars['Boolean']>;
  _lt?: Maybe<Scalars['timestamptz']>;
  _lte?: Maybe<Scalars['timestamptz']>;
  _neq?: Maybe<Scalars['timestamptz']>;
  _nin?: Maybe<Array<Scalars['timestamptz']>>;
};

/** columns and relationships of "user_account" */
export type UserAccount = {
   __typename?: 'user_account';
  /** An array relationship */
  aws_accounts: Array<AwsAccount>;
  /** An aggregated array relationship */
  aws_accounts_aggregate: AwsAccountAggregate;
  created_at: Scalars['timestamp'];
  email: Scalars['String'];
  id: Scalars['String'];
  /** An array relationship */
  projects: Array<Project>;
  /** An aggregated array relationship */
  projects_aggregate: ProjectAggregate;
  updated_at: Scalars['timestamp'];
  username: Scalars['String'];
};


/** columns and relationships of "user_account" */
export type UserAccountAwsAccountsArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** columns and relationships of "user_account" */
export type UserAccountAwsAccountsAggregateArgs = {
  distinct_on?: Maybe<Array<AwsAccountSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<AwsAccountOrderBy>>;
  where?: Maybe<AwsAccountBoolExp>;
};


/** columns and relationships of "user_account" */
export type UserAccountProjectsArgs = {
  distinct_on?: Maybe<Array<ProjectSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectOrderBy>>;
  where?: Maybe<ProjectBoolExp>;
};


/** columns and relationships of "user_account" */
export type UserAccountProjectsAggregateArgs = {
  distinct_on?: Maybe<Array<ProjectSelectColumn>>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  order_by?: Maybe<Array<ProjectOrderBy>>;
  where?: Maybe<ProjectBoolExp>;
};

/** aggregated selection of "user_account" */
export type UserAccountAggregate = {
   __typename?: 'user_account_aggregate';
  aggregate?: Maybe<UserAccountAggregateFields>;
  nodes: Array<UserAccount>;
};

/** aggregate fields of "user_account" */
export type UserAccountAggregateFields = {
   __typename?: 'user_account_aggregate_fields';
  count?: Maybe<Scalars['Int']>;
  max?: Maybe<UserAccountMaxFields>;
  min?: Maybe<UserAccountMinFields>;
};


/** aggregate fields of "user_account" */
export type UserAccountAggregateFieldsCountArgs = {
  columns?: Maybe<Array<UserAccountSelectColumn>>;
  distinct?: Maybe<Scalars['Boolean']>;
};

/** order by aggregate values of table "user_account" */
export type UserAccountAggregateOrderBy = {
  count?: Maybe<OrderBy>;
  max?: Maybe<UserAccountMaxOrderBy>;
  min?: Maybe<UserAccountMinOrderBy>;
};

/** input type for inserting array relation for remote table "user_account" */
export type UserAccountArrRelInsertInput = {
  data: Array<UserAccountInsertInput>;
  on_conflict?: Maybe<UserAccountOnConflict>;
};

/** Boolean expression to filter rows from the table "user_account". All fields are combined with a logical 'AND'. */
export type UserAccountBoolExp = {
  _and?: Maybe<Array<Maybe<UserAccountBoolExp>>>;
  _not?: Maybe<UserAccountBoolExp>;
  _or?: Maybe<Array<Maybe<UserAccountBoolExp>>>;
  aws_accounts?: Maybe<AwsAccountBoolExp>;
  created_at?: Maybe<TimestampComparisonExp>;
  email?: Maybe<StringComparisonExp>;
  id?: Maybe<StringComparisonExp>;
  projects?: Maybe<ProjectBoolExp>;
  updated_at?: Maybe<TimestampComparisonExp>;
  username?: Maybe<StringComparisonExp>;
};

/** unique or primary key constraints on table "user_account" */
export enum UserAccountConstraint {
  /** unique or primary key constraint */
  USER_ACCOUNT_EMAIL_KEY = 'user_account_email_key',
  /** unique or primary key constraint */
  USER_ACCOUNT_PKEY = 'user_account_pkey',
  /** unique or primary key constraint */
  USER_ACCOUNT_USERNAME_KEY = 'user_account_username_key'
}

/** input type for inserting data into table "user_account" */
export type UserAccountInsertInput = {
  aws_accounts?: Maybe<AwsAccountArrRelInsertInput>;
  created_at?: Maybe<Scalars['timestamp']>;
  email?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['String']>;
  projects?: Maybe<ProjectArrRelInsertInput>;
  updated_at?: Maybe<Scalars['timestamp']>;
  username?: Maybe<Scalars['String']>;
};

/** aggregate max on columns */
export type UserAccountMaxFields = {
   __typename?: 'user_account_max_fields';
  email?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['String']>;
  username?: Maybe<Scalars['String']>;
};

/** order by max() on columns of table "user_account" */
export type UserAccountMaxOrderBy = {
  email?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  username?: Maybe<OrderBy>;
};

/** aggregate min on columns */
export type UserAccountMinFields = {
   __typename?: 'user_account_min_fields';
  email?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['String']>;
  username?: Maybe<Scalars['String']>;
};

/** order by min() on columns of table "user_account" */
export type UserAccountMinOrderBy = {
  email?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  username?: Maybe<OrderBy>;
};

/** response of any mutation on the table "user_account" */
export type UserAccountMutationResponse = {
   __typename?: 'user_account_mutation_response';
  /** number of affected rows by the mutation */
  affected_rows: Scalars['Int'];
  /** data of the affected rows by the mutation */
  returning: Array<UserAccount>;
};

/** input type for inserting object relation for remote table "user_account" */
export type UserAccountObjRelInsertInput = {
  data: UserAccountInsertInput;
  on_conflict?: Maybe<UserAccountOnConflict>;
};

/** on conflict condition type for table "user_account" */
export type UserAccountOnConflict = {
  constraint: UserAccountConstraint;
  update_columns: Array<UserAccountUpdateColumn>;
  where?: Maybe<UserAccountBoolExp>;
};

/** ordering options when selecting data from "user_account" */
export type UserAccountOrderBy = {
  aws_accounts_aggregate?: Maybe<AwsAccountAggregateOrderBy>;
  created_at?: Maybe<OrderBy>;
  email?: Maybe<OrderBy>;
  id?: Maybe<OrderBy>;
  projects_aggregate?: Maybe<ProjectAggregateOrderBy>;
  updated_at?: Maybe<OrderBy>;
  username?: Maybe<OrderBy>;
};

/** primary key columns input for table: "user_account" */
export type UserAccountPkColumnsInput = {
  id: Scalars['String'];
};

/** select columns of table "user_account" */
export enum UserAccountSelectColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  EMAIL = 'email',
  /** column name */
  ID = 'id',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USERNAME = 'username'
}

/** input type for updating data in table "user_account" */
export type UserAccountSetInput = {
  created_at?: Maybe<Scalars['timestamp']>;
  email?: Maybe<Scalars['String']>;
  id?: Maybe<Scalars['String']>;
  updated_at?: Maybe<Scalars['timestamp']>;
  username?: Maybe<Scalars['String']>;
};

/** update columns of table "user_account" */
export enum UserAccountUpdateColumn {
  /** column name */
  CREATED_AT = 'created_at',
  /** column name */
  EMAIL = 'email',
  /** column name */
  ID = 'id',
  /** column name */
  UPDATED_AT = 'updated_at',
  /** column name */
  USERNAME = 'username'
}


export type ConfirmProjectInvitationMutationVariables = {
  input: ConfirmInvitationInput;
};


export type ConfirmProjectInvitationMutation = (
  { __typename?: 'mutation_root' }
  & { confirmInvitation: (
    { __typename?: 'ConfirmInvitationOutput' }
    & Pick<ConfirmInvitationOutput, 'result'>
  ) }
);

export type CreateActionMutationVariables = {
  input: ActionInsertInput;
};


export type CreateActionMutation = (
  { __typename?: 'mutation_root' }
  & { insert_action_one?: Maybe<(
    { __typename?: 'action' }
    & Pick<Action, 'id'>
  )> }
);

export type CreateAwsAccountMutationVariables = {
  input: AwsAccountInsertInput;
};


export type CreateAwsAccountMutation = (
  { __typename?: 'mutation_root' }
  & { insert_aws_account_one?: Maybe<(
    { __typename?: 'aws_account' }
    & Pick<AwsAccount, 'id'>
  )> }
);

export type CreateNotificationRuleMutationVariables = {
  input: NotificationRuleInsertInput;
};


export type CreateNotificationRuleMutation = (
  { __typename?: 'mutation_root' }
  & { insert_notification_rule_one?: Maybe<(
    { __typename?: 'notification_rule' }
    & Pick<NotificationRule, 'id'>
  )> }
);

export type CreateProjectMutationVariables = {
  input: ProjectInsertInput;
};


export type CreateProjectMutation = (
  { __typename?: 'mutation_root' }
  & { insert_project?: Maybe<(
    { __typename?: 'project_mutation_response' }
    & { returning: Array<(
      { __typename?: 'project' }
      & Pick<Project, 'id'>
    )> }
  )> }
);

export type CreateProjectInvitationMutationVariables = {
  input: ProjectInvitationInsertInput;
};


export type CreateProjectInvitationMutation = (
  { __typename?: 'mutation_root' }
  & { insert_project_invitation_one?: Maybe<(
    { __typename?: 'project_invitation' }
    & Pick<ProjectInvitation, 'project_id'>
  )> }
);

export type CreateSlackWebhookMutationVariables = {
  input: SlackWebhookInsertInput;
};


export type CreateSlackWebhookMutation = (
  { __typename?: 'mutation_root' }
  & { insert_slack_webhook_one?: Maybe<(
    { __typename?: 'slack_webhook' }
    & Pick<SlackWebhook, 'id'>
  )> }
);

export type DeleteActionMutationVariables = {
  id: Scalars['bigint'];
};


export type DeleteActionMutation = (
  { __typename?: 'mutation_root' }
  & { delete_action?: Maybe<(
    { __typename?: 'action_mutation_response' }
    & Pick<ActionMutationResponse, 'affected_rows'>
  )> }
);

export type DeleteNotificationRuleMutationVariables = {
  ruleId: Scalars['bigint'];
};


export type DeleteNotificationRuleMutation = (
  { __typename?: 'mutation_root' }
  & { delete_notification_rule?: Maybe<(
    { __typename?: 'notification_rule_mutation_response' }
    & Pick<NotificationRuleMutationResponse, 'affected_rows'>
  )> }
);

export type DeleteSlackWebHookMutationVariables = {
  id: Scalars['bigint'];
};


export type DeleteSlackWebHookMutation = (
  { __typename?: 'mutation_root' }
  & { delete_slack_webhook?: Maybe<(
    { __typename?: 'slack_webhook_mutation_response' }
    & Pick<SlackWebhookMutationResponse, 'affected_rows'>
  )> }
);

export type GetActionsQueryVariables = {
  projectId: Scalars['bigint'];
};


export type GetActionsQuery = (
  { __typename?: 'query_root' }
  & { action: Array<(
    { __typename?: 'action' }
    & Pick<Action, 'id' | 'name' | 'body'>
  )> }
);

export type AwsAccountByProjectIdQueryVariables = {
  projectId: Scalars['bigint'];
};


export type AwsAccountByProjectIdQuery = (
  { __typename?: 'query_root' }
  & { aws_account: Array<(
    { __typename?: 'aws_account' }
    & Pick<AwsAccount, 'id' | 'account_id' | 'role_name' | 'external_id'>
  )> }
);

export type InstanceQueryVariables = {
  projectId: Scalars['bigint'];
  instanceId: Scalars['String'];
};


export type InstanceQuery = (
  { __typename?: 'query_root' }
  & { instance: Array<(
    { __typename?: 'instance' }
    & Pick<Instance, 'instance_id' | 'name'>
  )> }
);

export type AwsInstancesQueryVariables = {
  projectId: Scalars['ID'];
};


export type AwsInstancesQuery = (
  { __typename?: 'query_root' }
  & { aws_instances: Array<(
    { __typename?: 'AwsInstance' }
    & Pick<AwsInstance, 'instanceId' | 'name' | 'status' | 'privateAddress' | 'publicAddress'>
    & { tags: (
      { __typename?: 'Tags' }
      & { tags: Array<Maybe<(
        { __typename?: 'Tag' }
        & Pick<Tag, 'key' | 'value'>
      )>> }
    ) }
  )> }
);

export type CountProjectQueryVariables = {
  name: Scalars['String'];
};


export type CountProjectQuery = (
  { __typename?: 'query_root' }
  & { project_aggregate: (
    { __typename?: 'project_aggregate' }
    & { aggregate?: Maybe<(
      { __typename?: 'project_aggregate_fields' }
      & Pick<ProjectAggregateFields, 'count'>
    )> }
  ) }
);

export type GetMetricQueryVariables = {
  instanceId: Scalars['String'];
  type: Scalars['String'];
  limit: Scalars['Int'];
};


export type GetMetricQuery = (
  { __typename?: 'query_root' }
  & { metrics: Array<(
    { __typename?: 'metrics' }
    & Pick<Metrics, 'time' | 'instance_id' | 'value'>
  )> }
);

export type InvitationQueryVariables = {
  projectId: Scalars['bigint'];
};


export type InvitationQuery = (
  { __typename?: 'query_root' }
  & { project_invitation: Array<(
    { __typename?: 'project_invitation' }
    & Pick<ProjectInvitation, 'mail_address'>
  )> }
);

export type NotificationRuleQueryVariables = {
  ruleId: Scalars['bigint'];
};


export type NotificationRuleQuery = (
  { __typename?: 'query_root' }
  & { notification_rule: Array<(
    { __typename?: 'notification_rule' }
    & Pick<NotificationRule, 'id' | 'rule_name' | 'rules'>
  )> }
);

export type NotificationRulesQueryVariables = {
  projectId: Scalars['bigint'];
};


export type NotificationRulesQuery = (
  { __typename?: 'query_root' }
  & { notification_rule: Array<(
    { __typename?: 'notification_rule' }
    & Pick<NotificationRule, 'id' | 'rule_name' | 'rules'>
  )> }
);

export type ProjectByNameQueryVariables = {
  name: Scalars['String'];
};


export type ProjectByNameQuery = (
  { __typename?: 'query_root' }
  & { project: Array<(
    { __typename?: 'project' }
    & Pick<Project, 'id' | 'name' | 'description' | 'created_at' | 'updated_at'>
  )> }
);

export type ProjectsQueryVariables = {};


export type ProjectsQuery = (
  { __typename?: 'query_root' }
  & { project: Array<(
    { __typename?: 'project' }
    & Pick<Project, 'id' | 'name' | 'description'>
  )> }
);

export type GetServicesQueryVariables = {};


export type GetServicesQuery = (
  { __typename?: 'query_root' }
  & { service: Array<(
    { __typename?: 'service' }
    & Pick<Service, 'id' | 'name'>
  )> }
);

export type GetSlackWebHookSettingsQueryVariables = {
  projectId: Scalars['bigint'];
};


export type GetSlackWebHookSettingsQuery = (
  { __typename?: 'query_root' }
  & { slack_webhook: Array<(
    { __typename?: 'slack_webhook' }
    & Pick<SlackWebhook, 'id' | 'name' | 'webhook_url' | 'channel'>
  )> }
);


export const ConfirmProjectInvitationDocument = gql`
    mutation ConfirmProjectInvitation($input: ConfirmInvitationInput!) {
  confirmInvitation(input: $input) {
    result
  }
}
    `;
export type ConfirmProjectInvitationMutationFn = ApolloReactCommon.MutationFunction<ConfirmProjectInvitationMutation, ConfirmProjectInvitationMutationVariables>;

/**
 * __useConfirmProjectInvitationMutation__
 *
 * To run a mutation, you first call `useConfirmProjectInvitationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useConfirmProjectInvitationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [confirmProjectInvitationMutation, { data, loading, error }] = useConfirmProjectInvitationMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useConfirmProjectInvitationMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<ConfirmProjectInvitationMutation, ConfirmProjectInvitationMutationVariables>) {
        return ApolloReactHooks.useMutation<ConfirmProjectInvitationMutation, ConfirmProjectInvitationMutationVariables>(ConfirmProjectInvitationDocument, baseOptions);
      }
export type ConfirmProjectInvitationMutationHookResult = ReturnType<typeof useConfirmProjectInvitationMutation>;
export type ConfirmProjectInvitationMutationResult = ApolloReactCommon.MutationResult<ConfirmProjectInvitationMutation>;
export type ConfirmProjectInvitationMutationOptions = ApolloReactCommon.BaseMutationOptions<ConfirmProjectInvitationMutation, ConfirmProjectInvitationMutationVariables>;
export const CreateActionDocument = gql`
    mutation CreateAction($input: action_insert_input!) {
  insert_action_one(object: $input, on_conflict: {constraint: action_pkey, update_columns: [name, body]}) {
    id
  }
}
    `;
export type CreateActionMutationFn = ApolloReactCommon.MutationFunction<CreateActionMutation, CreateActionMutationVariables>;

/**
 * __useCreateActionMutation__
 *
 * To run a mutation, you first call `useCreateActionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateActionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createActionMutation, { data, loading, error }] = useCreateActionMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateActionMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<CreateActionMutation, CreateActionMutationVariables>) {
        return ApolloReactHooks.useMutation<CreateActionMutation, CreateActionMutationVariables>(CreateActionDocument, baseOptions);
      }
export type CreateActionMutationHookResult = ReturnType<typeof useCreateActionMutation>;
export type CreateActionMutationResult = ApolloReactCommon.MutationResult<CreateActionMutation>;
export type CreateActionMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateActionMutation, CreateActionMutationVariables>;
export const CreateAwsAccountDocument = gql`
    mutation CreateAwsAccount($input: aws_account_insert_input!) {
  insert_aws_account_one(object: $input, on_conflict: {constraint: aws_account_pkey, update_columns: [account_id, role_name, external_id]}) {
    id
  }
}
    `;
export type CreateAwsAccountMutationFn = ApolloReactCommon.MutationFunction<CreateAwsAccountMutation, CreateAwsAccountMutationVariables>;

/**
 * __useCreateAwsAccountMutation__
 *
 * To run a mutation, you first call `useCreateAwsAccountMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateAwsAccountMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createAwsAccountMutation, { data, loading, error }] = useCreateAwsAccountMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateAwsAccountMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<CreateAwsAccountMutation, CreateAwsAccountMutationVariables>) {
        return ApolloReactHooks.useMutation<CreateAwsAccountMutation, CreateAwsAccountMutationVariables>(CreateAwsAccountDocument, baseOptions);
      }
export type CreateAwsAccountMutationHookResult = ReturnType<typeof useCreateAwsAccountMutation>;
export type CreateAwsAccountMutationResult = ApolloReactCommon.MutationResult<CreateAwsAccountMutation>;
export type CreateAwsAccountMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateAwsAccountMutation, CreateAwsAccountMutationVariables>;
export const CreateNotificationRuleDocument = gql`
    mutation CreateNotificationRule($input: notification_rule_insert_input!) {
  insert_notification_rule_one(object: $input, on_conflict: {constraint: notification_rule_pkey, update_columns: [rule_name, rules]}) {
    id
  }
}
    `;
export type CreateNotificationRuleMutationFn = ApolloReactCommon.MutationFunction<CreateNotificationRuleMutation, CreateNotificationRuleMutationVariables>;

/**
 * __useCreateNotificationRuleMutation__
 *
 * To run a mutation, you first call `useCreateNotificationRuleMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateNotificationRuleMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createNotificationRuleMutation, { data, loading, error }] = useCreateNotificationRuleMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateNotificationRuleMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<CreateNotificationRuleMutation, CreateNotificationRuleMutationVariables>) {
        return ApolloReactHooks.useMutation<CreateNotificationRuleMutation, CreateNotificationRuleMutationVariables>(CreateNotificationRuleDocument, baseOptions);
      }
export type CreateNotificationRuleMutationHookResult = ReturnType<typeof useCreateNotificationRuleMutation>;
export type CreateNotificationRuleMutationResult = ApolloReactCommon.MutationResult<CreateNotificationRuleMutation>;
export type CreateNotificationRuleMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateNotificationRuleMutation, CreateNotificationRuleMutationVariables>;
export const CreateProjectDocument = gql`
    mutation CreateProject($input: project_insert_input!) {
  insert_project(objects: [$input]) {
    returning {
      id
    }
  }
}
    `;
export type CreateProjectMutationFn = ApolloReactCommon.MutationFunction<CreateProjectMutation, CreateProjectMutationVariables>;

/**
 * __useCreateProjectMutation__
 *
 * To run a mutation, you first call `useCreateProjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateProjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createProjectMutation, { data, loading, error }] = useCreateProjectMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateProjectMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<CreateProjectMutation, CreateProjectMutationVariables>) {
        return ApolloReactHooks.useMutation<CreateProjectMutation, CreateProjectMutationVariables>(CreateProjectDocument, baseOptions);
      }
export type CreateProjectMutationHookResult = ReturnType<typeof useCreateProjectMutation>;
export type CreateProjectMutationResult = ApolloReactCommon.MutationResult<CreateProjectMutation>;
export type CreateProjectMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateProjectMutation, CreateProjectMutationVariables>;
export const CreateProjectInvitationDocument = gql`
    mutation CreateProjectInvitation($input: project_invitation_insert_input!) {
  insert_project_invitation_one(object: $input) {
    project_id
  }
}
    `;
export type CreateProjectInvitationMutationFn = ApolloReactCommon.MutationFunction<CreateProjectInvitationMutation, CreateProjectInvitationMutationVariables>;

/**
 * __useCreateProjectInvitationMutation__
 *
 * To run a mutation, you first call `useCreateProjectInvitationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateProjectInvitationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createProjectInvitationMutation, { data, loading, error }] = useCreateProjectInvitationMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateProjectInvitationMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<CreateProjectInvitationMutation, CreateProjectInvitationMutationVariables>) {
        return ApolloReactHooks.useMutation<CreateProjectInvitationMutation, CreateProjectInvitationMutationVariables>(CreateProjectInvitationDocument, baseOptions);
      }
export type CreateProjectInvitationMutationHookResult = ReturnType<typeof useCreateProjectInvitationMutation>;
export type CreateProjectInvitationMutationResult = ApolloReactCommon.MutationResult<CreateProjectInvitationMutation>;
export type CreateProjectInvitationMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateProjectInvitationMutation, CreateProjectInvitationMutationVariables>;
export const CreateSlackWebhookDocument = gql`
    mutation CreateSlackWebhook($input: slack_webhook_insert_input!) {
  insert_slack_webhook_one(object: $input, on_conflict: {constraint: slack_webhook_pkey, update_columns: [name, webhook_url, channel]}) {
    id
  }
}
    `;
export type CreateSlackWebhookMutationFn = ApolloReactCommon.MutationFunction<CreateSlackWebhookMutation, CreateSlackWebhookMutationVariables>;

/**
 * __useCreateSlackWebhookMutation__
 *
 * To run a mutation, you first call `useCreateSlackWebhookMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateSlackWebhookMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createSlackWebhookMutation, { data, loading, error }] = useCreateSlackWebhookMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateSlackWebhookMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<CreateSlackWebhookMutation, CreateSlackWebhookMutationVariables>) {
        return ApolloReactHooks.useMutation<CreateSlackWebhookMutation, CreateSlackWebhookMutationVariables>(CreateSlackWebhookDocument, baseOptions);
      }
export type CreateSlackWebhookMutationHookResult = ReturnType<typeof useCreateSlackWebhookMutation>;
export type CreateSlackWebhookMutationResult = ApolloReactCommon.MutationResult<CreateSlackWebhookMutation>;
export type CreateSlackWebhookMutationOptions = ApolloReactCommon.BaseMutationOptions<CreateSlackWebhookMutation, CreateSlackWebhookMutationVariables>;
export const DeleteActionDocument = gql`
    mutation DeleteAction($id: bigint!) {
  delete_action(where: {id: {_eq: $id}}) {
    affected_rows
  }
}
    `;
export type DeleteActionMutationFn = ApolloReactCommon.MutationFunction<DeleteActionMutation, DeleteActionMutationVariables>;

/**
 * __useDeleteActionMutation__
 *
 * To run a mutation, you first call `useDeleteActionMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteActionMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteActionMutation, { data, loading, error }] = useDeleteActionMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteActionMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<DeleteActionMutation, DeleteActionMutationVariables>) {
        return ApolloReactHooks.useMutation<DeleteActionMutation, DeleteActionMutationVariables>(DeleteActionDocument, baseOptions);
      }
export type DeleteActionMutationHookResult = ReturnType<typeof useDeleteActionMutation>;
export type DeleteActionMutationResult = ApolloReactCommon.MutationResult<DeleteActionMutation>;
export type DeleteActionMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteActionMutation, DeleteActionMutationVariables>;
export const DeleteNotificationRuleDocument = gql`
    mutation DeleteNotificationRule($ruleId: bigint!) {
  delete_notification_rule(where: {id: {_eq: $ruleId}}) {
    affected_rows
  }
}
    `;
export type DeleteNotificationRuleMutationFn = ApolloReactCommon.MutationFunction<DeleteNotificationRuleMutation, DeleteNotificationRuleMutationVariables>;

/**
 * __useDeleteNotificationRuleMutation__
 *
 * To run a mutation, you first call `useDeleteNotificationRuleMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteNotificationRuleMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteNotificationRuleMutation, { data, loading, error }] = useDeleteNotificationRuleMutation({
 *   variables: {
 *      ruleId: // value for 'ruleId'
 *   },
 * });
 */
export function useDeleteNotificationRuleMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<DeleteNotificationRuleMutation, DeleteNotificationRuleMutationVariables>) {
        return ApolloReactHooks.useMutation<DeleteNotificationRuleMutation, DeleteNotificationRuleMutationVariables>(DeleteNotificationRuleDocument, baseOptions);
      }
export type DeleteNotificationRuleMutationHookResult = ReturnType<typeof useDeleteNotificationRuleMutation>;
export type DeleteNotificationRuleMutationResult = ApolloReactCommon.MutationResult<DeleteNotificationRuleMutation>;
export type DeleteNotificationRuleMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteNotificationRuleMutation, DeleteNotificationRuleMutationVariables>;
export const DeleteSlackWebHookDocument = gql`
    mutation DeleteSlackWebHook($id: bigint!) {
  delete_slack_webhook(where: {id: {_eq: $id}}) {
    affected_rows
  }
}
    `;
export type DeleteSlackWebHookMutationFn = ApolloReactCommon.MutationFunction<DeleteSlackWebHookMutation, DeleteSlackWebHookMutationVariables>;

/**
 * __useDeleteSlackWebHookMutation__
 *
 * To run a mutation, you first call `useDeleteSlackWebHookMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteSlackWebHookMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteSlackWebHookMutation, { data, loading, error }] = useDeleteSlackWebHookMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteSlackWebHookMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<DeleteSlackWebHookMutation, DeleteSlackWebHookMutationVariables>) {
        return ApolloReactHooks.useMutation<DeleteSlackWebHookMutation, DeleteSlackWebHookMutationVariables>(DeleteSlackWebHookDocument, baseOptions);
      }
export type DeleteSlackWebHookMutationHookResult = ReturnType<typeof useDeleteSlackWebHookMutation>;
export type DeleteSlackWebHookMutationResult = ApolloReactCommon.MutationResult<DeleteSlackWebHookMutation>;
export type DeleteSlackWebHookMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteSlackWebHookMutation, DeleteSlackWebHookMutationVariables>;
export const GetActionsDocument = gql`
    query GetActions($projectId: bigint!) {
  action(where: {project_id: {_eq: $projectId}}) {
    id
    name
    body
  }
}
    `;

/**
 * __useGetActionsQuery__
 *
 * To run a query within a React component, call `useGetActionsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetActionsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetActionsQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useGetActionsQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<GetActionsQuery, GetActionsQueryVariables>) {
        return ApolloReactHooks.useQuery<GetActionsQuery, GetActionsQueryVariables>(GetActionsDocument, baseOptions);
      }
export function useGetActionsLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetActionsQuery, GetActionsQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<GetActionsQuery, GetActionsQueryVariables>(GetActionsDocument, baseOptions);
        }
export type GetActionsQueryHookResult = ReturnType<typeof useGetActionsQuery>;
export type GetActionsLazyQueryHookResult = ReturnType<typeof useGetActionsLazyQuery>;
export type GetActionsQueryResult = ApolloReactCommon.QueryResult<GetActionsQuery, GetActionsQueryVariables>;
export const AwsAccountByProjectIdDocument = gql`
    query awsAccountByProjectId($projectId: bigint!) {
  aws_account(where: {project_id: {_eq: $projectId}}) {
    id
    account_id
    role_name
    external_id
  }
}
    `;

/**
 * __useAwsAccountByProjectIdQuery__
 *
 * To run a query within a React component, call `useAwsAccountByProjectIdQuery` and pass it any options that fit your needs.
 * When your component renders, `useAwsAccountByProjectIdQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAwsAccountByProjectIdQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useAwsAccountByProjectIdQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<AwsAccountByProjectIdQuery, AwsAccountByProjectIdQueryVariables>) {
        return ApolloReactHooks.useQuery<AwsAccountByProjectIdQuery, AwsAccountByProjectIdQueryVariables>(AwsAccountByProjectIdDocument, baseOptions);
      }
export function useAwsAccountByProjectIdLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<AwsAccountByProjectIdQuery, AwsAccountByProjectIdQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<AwsAccountByProjectIdQuery, AwsAccountByProjectIdQueryVariables>(AwsAccountByProjectIdDocument, baseOptions);
        }
export type AwsAccountByProjectIdQueryHookResult = ReturnType<typeof useAwsAccountByProjectIdQuery>;
export type AwsAccountByProjectIdLazyQueryHookResult = ReturnType<typeof useAwsAccountByProjectIdLazyQuery>;
export type AwsAccountByProjectIdQueryResult = ApolloReactCommon.QueryResult<AwsAccountByProjectIdQuery, AwsAccountByProjectIdQueryVariables>;
export const InstanceDocument = gql`
    query Instance($projectId: bigint!, $instanceId: String!) {
  instance(where: {project_id: {_eq: $projectId}, instance_id: {_eq: $instanceId}}) {
    instance_id
    name
  }
}
    `;

/**
 * __useInstanceQuery__
 *
 * To run a query within a React component, call `useInstanceQuery` and pass it any options that fit your needs.
 * When your component renders, `useInstanceQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useInstanceQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      instanceId: // value for 'instanceId'
 *   },
 * });
 */
export function useInstanceQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<InstanceQuery, InstanceQueryVariables>) {
        return ApolloReactHooks.useQuery<InstanceQuery, InstanceQueryVariables>(InstanceDocument, baseOptions);
      }
export function useInstanceLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<InstanceQuery, InstanceQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<InstanceQuery, InstanceQueryVariables>(InstanceDocument, baseOptions);
        }
export type InstanceQueryHookResult = ReturnType<typeof useInstanceQuery>;
export type InstanceLazyQueryHookResult = ReturnType<typeof useInstanceLazyQuery>;
export type InstanceQueryResult = ApolloReactCommon.QueryResult<InstanceQuery, InstanceQueryVariables>;
export const AwsInstancesDocument = gql`
    query awsInstances($projectId: ID!) {
  aws_instances(projectId: $projectId) {
    instanceId
    name
    status
    privateAddress
    publicAddress
    tags {
      tags {
        key
        value
      }
    }
  }
}
    `;

/**
 * __useAwsInstancesQuery__
 *
 * To run a query within a React component, call `useAwsInstancesQuery` and pass it any options that fit your needs.
 * When your component renders, `useAwsInstancesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAwsInstancesQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useAwsInstancesQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<AwsInstancesQuery, AwsInstancesQueryVariables>) {
        return ApolloReactHooks.useQuery<AwsInstancesQuery, AwsInstancesQueryVariables>(AwsInstancesDocument, baseOptions);
      }
export function useAwsInstancesLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<AwsInstancesQuery, AwsInstancesQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<AwsInstancesQuery, AwsInstancesQueryVariables>(AwsInstancesDocument, baseOptions);
        }
export type AwsInstancesQueryHookResult = ReturnType<typeof useAwsInstancesQuery>;
export type AwsInstancesLazyQueryHookResult = ReturnType<typeof useAwsInstancesLazyQuery>;
export type AwsInstancesQueryResult = ApolloReactCommon.QueryResult<AwsInstancesQuery, AwsInstancesQueryVariables>;
export const CountProjectDocument = gql`
    query countProject($name: String!) {
  project_aggregate(where: {name: {_eq: $name}}) {
    aggregate {
      count
    }
  }
}
    `;

/**
 * __useCountProjectQuery__
 *
 * To run a query within a React component, call `useCountProjectQuery` and pass it any options that fit your needs.
 * When your component renders, `useCountProjectQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCountProjectQuery({
 *   variables: {
 *      name: // value for 'name'
 *   },
 * });
 */
export function useCountProjectQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<CountProjectQuery, CountProjectQueryVariables>) {
        return ApolloReactHooks.useQuery<CountProjectQuery, CountProjectQueryVariables>(CountProjectDocument, baseOptions);
      }
export function useCountProjectLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<CountProjectQuery, CountProjectQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<CountProjectQuery, CountProjectQueryVariables>(CountProjectDocument, baseOptions);
        }
export type CountProjectQueryHookResult = ReturnType<typeof useCountProjectQuery>;
export type CountProjectLazyQueryHookResult = ReturnType<typeof useCountProjectLazyQuery>;
export type CountProjectQueryResult = ApolloReactCommon.QueryResult<CountProjectQuery, CountProjectQueryVariables>;
export const GetMetricDocument = gql`
    query getMetric($instanceId: String!, $type: String!, $limit: Int!) {
  metrics(limit: $limit, order_by: {time: desc}, where: {type: {_eq: $type}, instance_id: {_eq: $instanceId}}) {
    time
    instance_id
    value
  }
}
    `;

/**
 * __useGetMetricQuery__
 *
 * To run a query within a React component, call `useGetMetricQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetMetricQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetMetricQuery({
 *   variables: {
 *      instanceId: // value for 'instanceId'
 *      type: // value for 'type'
 *      limit: // value for 'limit'
 *   },
 * });
 */
export function useGetMetricQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<GetMetricQuery, GetMetricQueryVariables>) {
        return ApolloReactHooks.useQuery<GetMetricQuery, GetMetricQueryVariables>(GetMetricDocument, baseOptions);
      }
export function useGetMetricLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetMetricQuery, GetMetricQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<GetMetricQuery, GetMetricQueryVariables>(GetMetricDocument, baseOptions);
        }
export type GetMetricQueryHookResult = ReturnType<typeof useGetMetricQuery>;
export type GetMetricLazyQueryHookResult = ReturnType<typeof useGetMetricLazyQuery>;
export type GetMetricQueryResult = ApolloReactCommon.QueryResult<GetMetricQuery, GetMetricQueryVariables>;
export const InvitationDocument = gql`
    query Invitation($projectId: bigint!) {
  project_invitation(where: {project_id: {_eq: $projectId}}) {
    mail_address
  }
}
    `;

/**
 * __useInvitationQuery__
 *
 * To run a query within a React component, call `useInvitationQuery` and pass it any options that fit your needs.
 * When your component renders, `useInvitationQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useInvitationQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useInvitationQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<InvitationQuery, InvitationQueryVariables>) {
        return ApolloReactHooks.useQuery<InvitationQuery, InvitationQueryVariables>(InvitationDocument, baseOptions);
      }
export function useInvitationLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<InvitationQuery, InvitationQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<InvitationQuery, InvitationQueryVariables>(InvitationDocument, baseOptions);
        }
export type InvitationQueryHookResult = ReturnType<typeof useInvitationQuery>;
export type InvitationLazyQueryHookResult = ReturnType<typeof useInvitationLazyQuery>;
export type InvitationQueryResult = ApolloReactCommon.QueryResult<InvitationQuery, InvitationQueryVariables>;
export const NotificationRuleDocument = gql`
    query NotificationRule($ruleId: bigint!) {
  notification_rule(where: {id: {_eq: $ruleId}}) {
    id
    rule_name
    rules
  }
}
    `;

/**
 * __useNotificationRuleQuery__
 *
 * To run a query within a React component, call `useNotificationRuleQuery` and pass it any options that fit your needs.
 * When your component renders, `useNotificationRuleQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useNotificationRuleQuery({
 *   variables: {
 *      ruleId: // value for 'ruleId'
 *   },
 * });
 */
export function useNotificationRuleQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<NotificationRuleQuery, NotificationRuleQueryVariables>) {
        return ApolloReactHooks.useQuery<NotificationRuleQuery, NotificationRuleQueryVariables>(NotificationRuleDocument, baseOptions);
      }
export function useNotificationRuleLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<NotificationRuleQuery, NotificationRuleQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<NotificationRuleQuery, NotificationRuleQueryVariables>(NotificationRuleDocument, baseOptions);
        }
export type NotificationRuleQueryHookResult = ReturnType<typeof useNotificationRuleQuery>;
export type NotificationRuleLazyQueryHookResult = ReturnType<typeof useNotificationRuleLazyQuery>;
export type NotificationRuleQueryResult = ApolloReactCommon.QueryResult<NotificationRuleQuery, NotificationRuleQueryVariables>;
export const NotificationRulesDocument = gql`
    query NotificationRules($projectId: bigint!) {
  notification_rule(where: {project_id: {_eq: $projectId}}) {
    id
    rule_name
    rules
  }
}
    `;

/**
 * __useNotificationRulesQuery__
 *
 * To run a query within a React component, call `useNotificationRulesQuery` and pass it any options that fit your needs.
 * When your component renders, `useNotificationRulesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useNotificationRulesQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useNotificationRulesQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<NotificationRulesQuery, NotificationRulesQueryVariables>) {
        return ApolloReactHooks.useQuery<NotificationRulesQuery, NotificationRulesQueryVariables>(NotificationRulesDocument, baseOptions);
      }
export function useNotificationRulesLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<NotificationRulesQuery, NotificationRulesQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<NotificationRulesQuery, NotificationRulesQueryVariables>(NotificationRulesDocument, baseOptions);
        }
export type NotificationRulesQueryHookResult = ReturnType<typeof useNotificationRulesQuery>;
export type NotificationRulesLazyQueryHookResult = ReturnType<typeof useNotificationRulesLazyQuery>;
export type NotificationRulesQueryResult = ApolloReactCommon.QueryResult<NotificationRulesQuery, NotificationRulesQueryVariables>;
export const ProjectByNameDocument = gql`
    query projectByName($name: String!) {
  project(where: {name: {_eq: $name}}) {
    id
    name
    description
    created_at
    updated_at
  }
}
    `;

/**
 * __useProjectByNameQuery__
 *
 * To run a query within a React component, call `useProjectByNameQuery` and pass it any options that fit your needs.
 * When your component renders, `useProjectByNameQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProjectByNameQuery({
 *   variables: {
 *      name: // value for 'name'
 *   },
 * });
 */
export function useProjectByNameQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ProjectByNameQuery, ProjectByNameQueryVariables>) {
        return ApolloReactHooks.useQuery<ProjectByNameQuery, ProjectByNameQueryVariables>(ProjectByNameDocument, baseOptions);
      }
export function useProjectByNameLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ProjectByNameQuery, ProjectByNameQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ProjectByNameQuery, ProjectByNameQueryVariables>(ProjectByNameDocument, baseOptions);
        }
export type ProjectByNameQueryHookResult = ReturnType<typeof useProjectByNameQuery>;
export type ProjectByNameLazyQueryHookResult = ReturnType<typeof useProjectByNameLazyQuery>;
export type ProjectByNameQueryResult = ApolloReactCommon.QueryResult<ProjectByNameQuery, ProjectByNameQueryVariables>;
export const ProjectsDocument = gql`
    query Projects {
  project {
    id
    name
    description
  }
}
    `;

/**
 * __useProjectsQuery__
 *
 * To run a query within a React component, call `useProjectsQuery` and pass it any options that fit your needs.
 * When your component renders, `useProjectsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProjectsQuery({
 *   variables: {
 *   },
 * });
 */
export function useProjectsQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<ProjectsQuery, ProjectsQueryVariables>) {
        return ApolloReactHooks.useQuery<ProjectsQuery, ProjectsQueryVariables>(ProjectsDocument, baseOptions);
      }
export function useProjectsLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<ProjectsQuery, ProjectsQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<ProjectsQuery, ProjectsQueryVariables>(ProjectsDocument, baseOptions);
        }
export type ProjectsQueryHookResult = ReturnType<typeof useProjectsQuery>;
export type ProjectsLazyQueryHookResult = ReturnType<typeof useProjectsLazyQuery>;
export type ProjectsQueryResult = ApolloReactCommon.QueryResult<ProjectsQuery, ProjectsQueryVariables>;
export const GetServicesDocument = gql`
    query GetServices {
  service {
    id
    name
  }
}
    `;

/**
 * __useGetServicesQuery__
 *
 * To run a query within a React component, call `useGetServicesQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetServicesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetServicesQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetServicesQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<GetServicesQuery, GetServicesQueryVariables>) {
        return ApolloReactHooks.useQuery<GetServicesQuery, GetServicesQueryVariables>(GetServicesDocument, baseOptions);
      }
export function useGetServicesLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetServicesQuery, GetServicesQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<GetServicesQuery, GetServicesQueryVariables>(GetServicesDocument, baseOptions);
        }
export type GetServicesQueryHookResult = ReturnType<typeof useGetServicesQuery>;
export type GetServicesLazyQueryHookResult = ReturnType<typeof useGetServicesLazyQuery>;
export type GetServicesQueryResult = ApolloReactCommon.QueryResult<GetServicesQuery, GetServicesQueryVariables>;
export const GetSlackWebHookSettingsDocument = gql`
    query GetSlackWebHookSettings($projectId: bigint!) {
  slack_webhook(where: {project_id: {_eq: $projectId}}) {
    id
    name
    webhook_url
    channel
  }
}
    `;

/**
 * __useGetSlackWebHookSettingsQuery__
 *
 * To run a query within a React component, call `useGetSlackWebHookSettingsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetSlackWebHookSettingsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetSlackWebHookSettingsQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useGetSlackWebHookSettingsQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<GetSlackWebHookSettingsQuery, GetSlackWebHookSettingsQueryVariables>) {
        return ApolloReactHooks.useQuery<GetSlackWebHookSettingsQuery, GetSlackWebHookSettingsQueryVariables>(GetSlackWebHookSettingsDocument, baseOptions);
      }
export function useGetSlackWebHookSettingsLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetSlackWebHookSettingsQuery, GetSlackWebHookSettingsQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<GetSlackWebHookSettingsQuery, GetSlackWebHookSettingsQueryVariables>(GetSlackWebHookSettingsDocument, baseOptions);
        }
export type GetSlackWebHookSettingsQueryHookResult = ReturnType<typeof useGetSlackWebHookSettingsQuery>;
export type GetSlackWebHookSettingsLazyQueryHookResult = ReturnType<typeof useGetSlackWebHookSettingsLazyQuery>;
export type GetSlackWebHookSettingsQueryResult = ApolloReactCommon.QueryResult<GetSlackWebHookSettingsQuery, GetSlackWebHookSettingsQueryVariables>;