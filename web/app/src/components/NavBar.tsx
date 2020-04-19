import React, { useState } from 'react';
import { makeStyles, useTheme, Theme } from '@material-ui/core/styles';
import clsx from 'clsx';

import { useHistory } from 'react-router-dom';
import Drawer from '@material-ui/core/Drawer';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import CloudQueueOutlinedIcon from '@material-ui/icons/CloudQueueOutlined';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import AccountCircle from '@material-ui/icons/AccountCircle';
import AddIcon from '@material-ui/icons/Add';
import DashboardIcon from '@material-ui/icons/Dashboard';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';

import { Link as RouterLink, LinkProps as RouterLinkProps } from 'react-router-dom';

import { useAuth0 } from '../utils/react-auth0-spa';

interface ListItemLinkProps {
  primary: string;
  to: string;
  icon?: React.ReactElement;
}

function ListItemLink(props: ListItemLinkProps) {
  const { icon, primary, to } = props;

  const renderLink = React.useMemo(
    () =>
      React.forwardRef<HTMLAnchorElement, Omit<RouterLinkProps, 'innerRef' | 'to'>>(
        (itemProps, ref) => (
          // With react-router-dom@^6.0.0 use `ref` instead of `innerRef`
          // See https://github.com/ReactTraining/react-router/issues/6056
          <RouterLink to={to} {...itemProps} innerRef={ref} />
        ),
      ),
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

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) => ({
  title: {
    flexGrow: 1,
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
  hide: {
    display: 'none',
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
    whiteSpace: 'nowrap',
  },
  drawerOpen: {
    width: drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
  drawerClose: {
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
    overflowX: 'hidden',
    width: theme.spacing(7) + 1,
    [theme.breakpoints.up('sm')]: {
      width: theme.spacing(9) + 1,
    },
  },
  toolbar: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: theme.spacing(0, 1),
    ...theme.mixins.toolbar,
  },
}));

const NavBar: React.FC = () => {
  const classes = useStyles();

  const theme = useTheme();

  const [open, setOpen] = useState(false);

  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();

  const [accountAnchorEl, setAccountAnchorEl] = React.useState<null | EventTarget>(null);

  const [addAnchorEl, setAddAnchorEl] = React.useState<null | EventTarget>(null);

  const accountMenuOpen = Boolean(accountAnchorEl);
  const addMenuOpen = Boolean(addAnchorEl);

  const history = useHistory();

  const logoutWithRedirect = () =>
    logout({
      returnTo: window.location.origin,
    });

  const handleAccountMenu = (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setAccountAnchorEl(event.currentTarget);
  };

  const handleAddMenu = (event: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    setAddAnchorEl(event.currentTarget);
  };

  const handleAccountMenuClose = () => {
    setAccountAnchorEl(null);
  };

  const handleAddMenuClose = () => {
    setAddAnchorEl(null);
  };

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  const clickNewProjectLink = () => {
    history.push('/new');
    setAddAnchorEl(null);
  };

  return (
    <>
      <AppBar
        position="fixed"
        className={clsx(classes.appBar, {
          [classes.appBarShift]: open,
        })}
      >
        <Toolbar>
          {isAuthenticated && (
            <IconButton
              color="inherit"
              aria-label="open drawer"
              onClick={handleDrawerOpen}
              edge="start"
              className={clsx(classes.menuButton, {
                [classes.hide]: open,
              })}
            >
              <MenuIcon />
            </IconButton>
          )}

          <Typography variant="h6" className={classes.title}>
            Inflion
          </Typography>

          {!isAuthenticated && (
            <Button color="inherit" onClick={() => loginWithRedirect({})}>
              Login
            </Button>
          )}

          {isAuthenticated && (
            <>
              <IconButton
                aria-label="add"
                aria-controls="menu-add"
                aria-haspopup="true"
                onClick={handleAddMenu}
                color="inherit"
              >
                <AddIcon />
              </IconButton>

              <IconButton
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleAccountMenu}
                color="inherit"
              >
                <AccountCircle />
              </IconButton>

              <Menu
                id="menu-appbar"
                anchorEl={accountAnchorEl as HTMLAnchorElement}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={accountMenuOpen}
                onClose={handleAccountMenuClose}
              >
                <MenuItem onClick={logoutWithRedirect}>Logout</MenuItem>
              </Menu>

              <Menu
                id="menu-add"
                anchorEl={addAnchorEl as HTMLAnchorElement}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={addMenuOpen}
                onClose={handleAddMenuClose}
              >
                <MenuItem onClick={clickNewProjectLink}>New Project</MenuItem>
              </Menu>
            </>
          )}
        </Toolbar>
      </AppBar>

      {isAuthenticated && (
        <Drawer
          variant="permanent"
          className={clsx(classes.drawer, {
            [classes.drawerOpen]: open,
            [classes.drawerClose]: !open,
          })}
          classes={{
            paper: clsx({
              [classes.drawerOpen]: open,
              [classes.drawerClose]: !open,
            }),
          }}
          open={open}
        >
          <div className={classes.toolbar}>
            <IconButton onClick={handleDrawerClose}>
              {theme.direction === 'rtl' ? <ChevronRightIcon /> : <ChevronLeftIcon />}
            </IconButton>
          </div>

          <Divider />

          <List>
            <ListItemLink to="/" primary="Dashboard" icon={<DashboardIcon />} />
            <ListItemLink to="/projects" primary="Projects" icon={<CloudQueueOutlinedIcon />} />
          </List>
        </Drawer>
      )}
    </>
  );
};

export default NavBar;
