import Button from '@material-ui/core/Button';
import React, { useState } from 'react';
import Highlight from '../../components/Highlight';
import { useAuth0 } from '../../utils/react-auth0-spa';

export const CallGraphQL = () => {
  const { getTokenSilently } = useAuth0();

  const [showResult, setShowResult] = useState(false);
  const [apiMessage, setApiMessage] = useState('');

  const callApi = async () => {
    try {
      const token = await getTokenSilently();

      // const response = await fetch('http://localhost:3002/graphql?query={ user(id: "4") { name } }', {
      const response = await fetch(
        //'http://localhost:8080/graphql?query={ instance(id: "1") { id instanceId } }',
        'http://localhost:8080/graphql?query={ instances { id instanceId } }',
        {
          method: 'GET',
          // body: JSON.stringify({ query: '{ user { name } }'}),
          headers: {
            Authorization: `Bearer ${token}`,
          },
        },
      );

      const responseData = await response.json();

      setShowResult(true);
      setApiMessage(responseData);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      {/* <Button color="primary" className="mt-5" onClick={callApi}> */}
      <Button variant="contained" color="primary" onClick={callApi}>
        Ping API
      </Button>

      <div className="result-block-container">
        <div className={`result-block ${showResult && 'show'}`}>
          <h6 className="muted">Result</h6>
          <Highlight>{JSON.stringify(apiMessage, null, 2)}</Highlight>
        </div>
      </div>
    </>
  );
};
