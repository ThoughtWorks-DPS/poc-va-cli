$CURRENT_VERSION = git describe --abbrev=0 | tr -d 'v'

url="https://github.com/ThoughtWorks-DPS/${PROJECT}/releases/download/v${CURRENT_VERSION}/${PROJECT}_${CURRENT_VERSION}_${IMAGE}"

Invoke-RestMethod -Uri $url -Method Get -ContentType "application/zip"  -OutFile ${PROJECT}_${IMAGE}
Expand-Archive -LiteralPath ./${PROJECT}_${IMAGE} -DestinationPath ./

./poc-va-cli get
