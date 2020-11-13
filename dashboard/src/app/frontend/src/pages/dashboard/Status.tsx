import React from 'react';
import { Typography, Button, Card, CardContent, CardActions } from '@material-ui/core';
import styled from 'styled-components';

const StatusCard = styled(Card)`
  width: 260px;
  height: 200px;
  background-color: ${(props) => props.color};
  position: relative;
`;

const BottomCardActions = styled(CardActions)`
  position: absolute;
  bottom: 0;
  right: 0;
`;

type StatusProps = {
  type: 'ok' | 'warn' | 'ng';
  name: string;
};

const colors = { ok: '#C3FFB9', warn: '#FEFFBF', ng: '#FFBFBF' };

export const Status: React.FC<StatusProps> = ({ name, type }) => {
  return (
    <>
      <StatusCard color={colors[type]}>
        <CardContent>
          <Typography variant="h6" component="h2">
            {name}
          </Typography>
        </CardContent>
        <BottomCardActions>
          <Button size="small">Details</Button>
        </BottomCardActions>
      </StatusCard>
    </>
  );
};
