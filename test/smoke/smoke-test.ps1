$CURRENT_VERSION = git describe --tag --abbrev=0
$IMAGE = "Windows_x86_64.zip"
$PROJECT = "poc-va-cli"

$url='https://github.com/ThoughtWorks-DPS/' + $PROJECT + '/releases/download/v' + $CURRENT_VERSION + '/' + $PROJECT + '_' + $CURRENT_VERSION + '_' + $IMAGE

curl -L $url --output cli.zip
unzip cli.zip -d cli/

./cli/poc-va-cli.exe get
