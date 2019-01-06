#!/usr/bin/env bash

GOOS=linux go build -o definitionGame
zip deployment.zip definitionGame
mv definitionGameDeployment.zip ~/Desktop
rm definitionGame
