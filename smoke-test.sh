#!/bin/bash
#Should probably change executor if git is going to be used
CURRENT_VERSION=$(git describe --abbrev=0 | tr -d 'v')

url="https://github.com/ThoughtWorks-DPS/${PROJECT}/releases/download/v${CURRENT_VERSION}/${PROJECT}_${CURRENT_VERSION}_${IMAGE}"
echo $CURRENT_VERSION
echo $url

curl -L $url --output ${PROJECT}_${IMAGE}

if [ $IMAGE = "Windows_x86_64.zip" ]; then
  echo "im in windows"
  unzip -o ${PROJECT}_${IMAGE}
else
  tar -xvzf ${PROJECT}_${IMAGE}
fi

./poc-va-cli get