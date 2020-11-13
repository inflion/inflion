import React from 'react';
import Typography from '@material-ui/core/Typography';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import Link from '@material-ui/core/Link';

import { useParams } from 'react-router-dom';

export const DefaultBreadcrumbs = () => {
  const { orgId, projectId } = useParams();

  return (
    <Breadcrumbs aria-label="breadcrumb">
      <Link color="inherit" href="/">
        inflion/inflion
      </Link>

      <Typography color="textPrimary">{orgId}</Typography>

      <Typography color="textPrimary">{projectId}</Typography>
    </Breadcrumbs>
  );
};
