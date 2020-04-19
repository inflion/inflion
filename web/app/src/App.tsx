import React, { useEffect } from 'react';
import styled from 'styled-components';
import { Route, BrowserRouter as Router, Switch } from 'react-router-dom';

import { makeStyles, Theme } from '@material-ui/core/styles';

import Container from '@material-ui/core/Container';
import CssBaseline from '@material-ui/core/CssBaseline';
import Button from '@material-ui/core/Button';

import { About } from './pages/About';
import { Organizations } from './pages/organizations/Organizations';
import { InstanceDetail } from './pages/projects/instance/InstanceDetail';
import { OrganizationTop } from './pages/organizations/OrganizationTop';
import { ProjectPage } from './pages/projects/ProjectPage';
import { useAuth0 } from './utils/react-auth0-spa';

import NavBar from './components/NavBar';
import CircularProgress from '@material-ui/core/CircularProgress';

import i18n from 'i18next';
import { useTranslation, initReactI18next } from 'react-i18next';

import { resources } from './i18n/resources';
import { UserRegistration } from './pages/UserRegistration';
import { NewProjectPage } from './pages/projects/NewProjectPage';
import { ProjectInvitationConfirm } from './pages/projects/invitation/Confirmation';

import { useConfirmProjectInvitationMutation } from './graphql';
import { Dashboard } from './pages/dashboard/Dashboard';

i18n
  .use(initReactI18next) // passes i18n down to react-i18next
  .init({
    resources,
    lng: 'ja',
    fallbackLng: 'en',
    interpolation: {
      escapeValue: false,
    },
  })
  .catch((reason) => console.log(reason));

const Root = styled.div`
  display: flex;
  -webkit-font-smoothing: 'antialiased';
  -moz-osx-font-smoothing: 'grayscale';
`;

const Content = styled.main`
  flex-grow: 1;
  padding: ${(props) => props.theme.spacing(3, 2)};
`;

const useStyles = makeStyles((theme: Theme) => ({
  toolbar: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: theme.spacing(0, 1),
    ...theme.mixins.toolbar,
  },
}));

export const App = () => {
  const classes = useStyles();
  const { t } = useTranslation();
  const { loading, loginWithRedirect, isAuthenticated, getTokenSilently, user } = useAuth0();
  const [confirmProjectInvitation] = useConfirmProjectInvitationMutation();

  useEffect(() => {
    const fetchToken = async () => {
      const token: string | undefined = await getTokenSilently();
      if (typeof token == 'undefined') {
        return;
      }
      localStorage.setItem('token', token);
    };
    if (loading) {
      return;
    }
    if (!isAuthenticated) {
      return;
    }

    fetchToken().catch((reason) => console.log(reason));
  }, [loading, isAuthenticated, getTokenSilently, user]);

  const token = sessionStorage.getItem('confirmation');

  if (isAuthenticated && token) {
    confirmProjectInvitation({ variables: { input: { token: token } } });
    sessionStorage.removeItem('confirmation');
  }

  return (
    <>
      <CssBaseline />

      <Root>
        <Router>
          <NavBar />

          <Content>
            <div className={classes.toolbar} />

            {loading && <CircularProgress />}

            {!loading && isAuthenticated ? (
              <Container maxWidth="lg">
                <Switch>
                  <Route path="/about">
                    <About />
                  </Route>
                  <Route path="/new">
                    <NewProjectPage />
                  </Route>
                  <Route path="/orgs">
                    <Organizations />
                  </Route>
                  <Route path="/instance/:instanceId">
                    <InstanceDetail />
                  </Route>
                  <Route path="/projects/:projectName">
                    <ProjectPage />
                  </Route>
                  <Route path="/user_registration">
                    <UserRegistration />
                  </Route>
                  <Route path="/:orgId">
                    <OrganizationTop />
                  </Route>
                  <Route path="/">
                    <Dashboard />
                  </Route>
                </Switch>
              </Container>
            ) : (
              <>
                <Switch>
                  <Route path="/project/invitation/confirm/:token">
                    <ProjectInvitationConfirm />
                  </Route>
                </Switch>

                <div>
                  {t('You are not authorized')}
                  {t('Please login')}
                </div>
                <Button variant="contained" color="inherit" onClick={() => loginWithRedirect({})}>
                  {t('Login')}
                </Button>
              </>
            )}
          </Content>
        </Router>
      </Root>
    </>
  );
};
