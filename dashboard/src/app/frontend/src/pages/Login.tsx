import React, {useState, useEffect} from 'react';
import Button from '@material-ui/core/Button';
import { useTranslation } from 'react-i18next';
import * as rm from 'typed-rest-client/RestClient';
import useCookie from 'react-use-cookie';
import {useInputChange} from '../utils/useInputChange';

const restc: rm.RestClient = new rm.RestClient('rest-samples', 'http://localhost:3000/');

function httpBinOptions(): rm.IRequestOptions {
  const options: rm.IRequestOptions = {};
  options.responseProcessor = (obj: any) => {
    return obj;
  };
  return options;
}

interface Login {
  username?: string;
  password?: string;
  jweToken?: string;
  errors?: string[];
}

interface LoginForm {
  username: string;
  password: string;
}

export const Login: React.FC = () => {
  const { t } = useTranslation();
  const [jweToken, setJweToken] = useCookie('inflion-jwe-token', '');
  const [input, handleInputChange] = useInputChange<LoginForm>()

  const submit = async () => {
    const options: rm.IRequestOptions = httpBinOptions();
    const loginRequest: Login = { username: input.username, password: input.password };

    try {
      const response: rm.IRestResponse<Login> = await restc.create<Login>(
        '/api/v1/login',
        loginRequest,
        options,
      );
      if (response.result !== undefined && response.result !== null) {
        setJweToken(response.result.jweToken || '')
        window.location.assign('/projects');
      }
    } catch(error) {
      console.log(error)
    }
  };

  return (
    <>
      <div>
        <input type="text" name="username" onChange={handleInputChange}/>
        <input type="password" name="password" onChange={handleInputChange}/>
        <Button variant="contained" color="inherit" onClick={() => submit()}>
          {t('Login')}
        </Button>
      </div>
    </>
  );
};
