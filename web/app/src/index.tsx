import { ApolloProvider } from '@apollo/react-hooks';
import { ThemeProvider as MaterialThemeProvider } from '@material-ui/core/styles';
import { StylesProvider } from '@material-ui/styles';
import React from 'react';
import ReactDOM from 'react-dom';
import { ThemeProvider as StyledThemeProvider } from 'styled-components';
import App from './App';

import authConfig from './config/auth.json';
import { Auth0Provider } from './utils/react-auth0-spa';
import * as serviceWorker from './utils/serviceWorker';
import theme from './themes/theme';
import history from './utils/history';
import { apolloClient } from './graphql';

const onRedirectCallback = (appState: any) => {
  history.push(appState && appState.targetUrl ? appState.targetUrl : window.location.pathname);
};

ReactDOM.render(
  <StylesProvider injectFirst>
    <MaterialThemeProvider theme={theme}>
      <StyledThemeProvider theme={theme}>
        <Auth0Provider
          domain={authConfig.domain}
          client_id={authConfig.clientId}
          audience={authConfig.audience}
          redirect_uri={window.location.origin}
          onRedirectCallback={onRedirectCallback}
        >
          <ApolloProvider client={apolloClient}>
            <App />
          </ApolloProvider>
        </Auth0Provider>
      </StyledThemeProvider>
    </MaterialThemeProvider>
  </StylesProvider>,
  document.getElementById('root'),
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
