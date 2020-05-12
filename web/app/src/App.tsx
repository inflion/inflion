import Button from '@material-ui/core/Button';
import CircularProgress from '@material-ui/core/CircularProgress';
import CssBaseline from '@material-ui/core/CssBaseline';
import { makeStyles } from '@material-ui/core/styles';
import clsx from 'clsx';
import i18n from 'i18next';
import React, { useEffect } from 'react';
import { initReactI18next, useTranslation } from 'react-i18next';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import NavBar from './components/NavBar';
import { useConfirmProjectInvitationMutation } from './graphql';
import { resources } from './i18n/resources';
import { About } from './pages/About';
import { Dashboard } from './pages/dashboard/Dashboard';
import { Organizations } from './pages/organizations/Organizations';
import { OrganizationTop } from './pages/organizations/OrganizationTop';
import { ProjectInvitationConfirm } from './pages/projects/invitation/Confirmation';
import { NewProjectPage } from './pages/projects/NewProjectPage';
import { ProjectPage } from './pages/projects/ProjectPage';
import { Projects } from './pages/projects/Projects';
import { UserRegistration } from './pages/UserRegistration';
import { useAuth0 } from './utils/react-auth0-spa';
const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
  },
  toolbar: {
    paddingRight: 24, // keep right padding when drawer closed
  },
  toolbarIcon: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: '0 8px',
    ...theme.mixins.toolbar,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  appBarShift: {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  menuButton: {
    marginRight: 36,
  },
  menuButtonHidden: {
    display: 'none',
  },
  title: {
    flexGrow: 1,
  },
  drawerPaper: {
    position: 'relative',
    whiteSpace: 'nowrap',
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerPaperClose: {
    overflowX: 'hidden',
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    width: theme.spacing(7),
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing(9),
    },
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: '100vh',
    overflow: 'auto',
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  paper: {
    padding: theme.spacing(2),
    display: 'flex',
    overflow: 'auto',
    flexDirection: 'column',
  },
  fixedHeight: {
    height: 240,
  },
}));

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

export default function App() {
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
  const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
  return (
    <>
      <div className={classes.root}>
        <CssBaseline />
        <Router>
          <NavBar />
          <main className={classes.content}>
            <div className={classes.appBarSpacer} />

            <div className={classes.toolbar} />

            {loading && <CircularProgress />}

            {!loading && isAuthenticated ? (
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
                <Route path="/instance/:instanceId"></Route>
                <Route path="/projects/:projectName">
                  <ProjectPage />
                </Route>
                <Route path="/projects">
                  <Projects />
                </Route>
                >
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
          </main>
        </Router>
      </div>
    </>
  );
}
