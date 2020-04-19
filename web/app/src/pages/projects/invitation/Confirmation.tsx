import React from 'react';
import { useParams } from 'react-router-dom';

export const ProjectInvitationConfirm: React.FC = () => {
  const { token } = useParams();

  if (token !== undefined) {
    sessionStorage.setItem('confirmation', token);
  }

  return <></>;
};
