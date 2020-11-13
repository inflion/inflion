import React from 'react';
import { NewOrgForm } from './NewOrgForm';
import { OrgsList } from './OrgsList';

export const Organizations: React.FC = () => {
  let notify: () => void = () => {};

  const handleCreated = (created: { name: string; displayName: string }) => {
    notify();
  };

  const observer = (callback: () => void) => {
    notify = callback;
  };

  return (
    <>
      <OrgsList update={observer} />
      <NewOrgForm created={handleCreated} />
    </>
  );
};
