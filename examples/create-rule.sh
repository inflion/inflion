#!/bin/bash

pushd .

if [ ! -f ../bin/lionctl ]; then
  cd ../lionctl && go build -o ../bin/lionctl
fi

popd

../bin/lionctl rule create -f ./rule.json