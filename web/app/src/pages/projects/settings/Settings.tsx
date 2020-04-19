import React from 'react';
import NotificationsIcon from '@material-ui/icons/Notifications';
import { Grid, Paper, List, ListItem, ListItemIcon, ListItemText } from '@material-ui/core';
import { Route, MemoryRouter } from 'react-router';
import { Link as RouterLink, LinkProps as RouterLinkProps, Switch } from 'react-router-dom';
import { Omit } from '@material-ui/types';
import { Slack } from './slack/Slack';

interface ListItemLinkProps {
  icon?: React.ReactElement;
  primary: string;
  to: string;
}

function ListItemLink(props: ListItemLinkProps) {
  const { icon, primary, to } = props;

  const renderLink = React.useMemo(
    () =>
      React.forwardRef<any, Omit<RouterLinkProps, 'to'>>((itemProps, ref) => (
        <RouterLink to={to} ref={ref} {...itemProps} />
      )),
    [to],
  );

  return (
    <li>
      <ListItem button component={renderLink}>
        {icon ? <ListItemIcon>{icon}</ListItemIcon> : null}
        <ListItemText primary={primary} />
      </ListItem>
    </li>
  );
}

export const Settings = () => {
  return (
    <>
      <MemoryRouter initialEntries={['/slack']} initialIndex={0}>
        <Grid container alignItems="flex-start" spacing={2}>
          <Grid item xs={3}>
            <Paper variant="outlined">
              <List component="nav" aria-label="main mailbox folders">
                <ListItemLink to="/slack" primary="Slack" icon={<NotificationsIcon />} />
              </List>
            </Paper>
          </Grid>

          <Grid item xs={9}>
            <Switch>
              <Route path="/slack">
                <Slack></Slack>
              </Route>
            </Switch>
          </Grid>
        </Grid>
      </MemoryRouter>
    </>
  );
};
