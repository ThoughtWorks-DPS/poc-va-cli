#!/bin/bash
#Should probably change executor if git is going to be used
CURRENT_VERSION=$(git describe --abbrev=0 | tr -d 'v')
IMAGE="Linux_i386.tar.gz"
PROJECT="poc-va-cli"

url="https://github.com/ThoughtWorks-DPS/${PROJECT}/releases/download/v${CURRENT_VERSION}/${PROJECT}_${CURRENT_VERSION}_${IMAGE}"
echo $CURRENT_VERSION
echo $url
curl -L $url --output ${PROJECT}_${IMAGE}
tar -xvzf ${PROJECT}_${IMAGE}
./poc-va-cli get