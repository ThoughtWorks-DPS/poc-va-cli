#!/bin/bash
#Should probably change executor if git is going to be used
CURRENT_VERSION=$(git describe --abbrev=0 | tr -d 'v')

url="https://github.com/ThoughtWorks-DPS/${PROJECT}/releases/download/v${CURRENT_VERSION}/${PROJECT}_${CURRENT_VERSION}_${IMAGE}"
curl -L $url --output ${PROJECT}_${IMAGE}
tar -xvzf ${PROJECT}_${IMAGE}


./voltron help

