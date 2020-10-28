#!/bin/bash
#Should probably change executor if git is going to be used

url="https://github.com/ThoughtWorks-DPS/${PROJECT}/releases/download/v${CURRENT_VERSION}/${PROJECT}_${CURRENT_VERSION}_${IMAGE}"
echo $CURRENT_VERSION
echo $url
curl -L $url --output ${PROJECT}_${IMAGE}
tar -xvzf ${PROJECT}_${IMAGE}
./poc-va-cli get