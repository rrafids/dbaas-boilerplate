#!/bin/bash

go build -o dspro
if [ $? -ne 0 ]; then
   exit 1
fi

./dspro generate elasticsearch app/int/test-app.yaml
