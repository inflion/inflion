import React, { useState, useEffect, useContext } from 'react';
import useCookie from 'react-use-cookie';

interface RedirectLoginOptions {
  /**
   * The URL where Auth0 will redirect your browser to with
   * the authentication result. It must be whitelisted in
   * the "Allowed Callback URLs" field in your Auth0 Application's
   * settings.
   */
  redirect_uri?: string;
}

interface LogoutOptions {
  /**
   * The URL where Auth0 will redirect your browser to after the logout.
   *
   * > Note that if the `client_id` parameter is included, the
   * `returnTo` URL that is provided must be listed in the
   * Application's "Allowed Logout URLs" in the Auth0 dashboard.
   * However, if the `client_id` parameter is not included, the
   * `returnTo` URL must be listed in the "Allowed Logout URLs" at
   * the account level in the Auth0 dashboard.
   */
  returnTo?: string;
}

async function loginWithRedirect(options: RedirectLoginOptions = {}) {
  window.location.assign(options.redirect_uri || '/login');
}

interface AuthContext {
  isAuthenticated: boolean;
  user: any;
  loading: boolean;
  loginWithRedirect(o: RedirectLoginOptions): Promise<void>;
  logout(o?: LogoutOptions): void;
}
interface AuthProviderOptions {
  children: React.ReactElement;
}

export const AuthContext = React.createContext<AuthContext | null>(null);

export const useAuth = () => useContext(AuthContext)!;

export const AuthProvider = ({ children, ...initOptions }: AuthProviderOptions) => {
  const [jweToken, setJweToken] = useCookie('inflion-jwe-token', '');
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState();
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const init = async () => {
      if (jweToken === '') {
        setIsAuthenticated(false);
      } else {
        setIsAuthenticated(true);
      }
    };
    init();
    // eslint-disable-next-line
  }, []);

  const handleRedirectCallback = async () => {
    setIsAuthenticated(true);
    setUser(user);
  };
  return (
    <AuthContext.Provider
      value={{
        isAuthenticated,
        user,
        loading,
        loginWithRedirect: (o: RedirectLoginOptions) => loginWithRedirect(o),
        logout: (o: LogoutOptions | undefined) => {
          console.log('logout')
          setJweToken('');
          window.location.assign(o?.returnTo || '/')
          return;
        },
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
