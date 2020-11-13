import { setContext } from 'apollo-link-context';
import { ApolloClient } from 'apollo-client';
import { ApolloLink } from 'apollo-link';
import { onError } from 'apollo-link-error';
import { HttpLink } from 'apollo-link-http';
import graphqlConfig from '../config/graphql.json';
import { InMemoryCache } from 'apollo-cache-inmemory';

const authLink = setContext((_, { headers }) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem('token');
  // return the headers to the context so httpLink can read them
  console.log(`Bearer ${token}`);
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : '',
    },
  };
});

const httpLink = new HttpLink({
  uri: graphqlConfig.endpoint,
  credentials: 'same-origin',
});

const apolloLink = ApolloLink.from([
  onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors)
      graphQLErrors.forEach(({ message, locations, path }) =>
        console.log(`[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`),
      );
    if (networkError) console.log(`[Network error]: ${networkError}`);
  }),
  authLink.concat(httpLink),
]);

export const apolloClient = new ApolloClient({
  connectToDevTools: true,
  link: apolloLink,
  cache: new InMemoryCache(),
});
