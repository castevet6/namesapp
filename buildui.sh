#!/bin/bash

if [ ! -d "./ui" ]; then
    mkdir ui
fi

cd ui-client/
npm run build
cd ..

