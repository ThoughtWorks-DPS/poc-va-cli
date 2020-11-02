$CURRENT_VERSION = git describe --abbrev=0 | tr -d 'v'

url="https://github.com/ThoughtWorks-DPS/${PROJECT}/releases/download/v${CURRENT_VERSION}/${PROJECT}_${CURRENT_VERSION}_${IMAGE}"

curl -L $url --output ${PROJECT}_${IMAGE}Expand-Archive -LiteralPath ./${PROJECT}_${IMAGE} -DestinationPath ./
unzip ${PROJECT}_${IMAGE}

./poc-va-cli get
