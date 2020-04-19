#!/bin/bash

SECRET="inflion"

cd main && hasura migrate apply --admin-secret $SECRET
cd timescale && hasura migrate apply --admin-secret $SECRET
