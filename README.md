# POC-VA-CLI for delivery infrastructure
[![Maintainability](https://api.codeclimate.com/v1/badges/043b7771faf510da9b7a/maintainability)](https://codeclimate.com/github/ThoughtWorks-DPS/poc-va-cli/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/043b7771faf510da9b7a/test_coverage)](https://codeclimate.com/github/ThoughtWorks-DPS/poc-va-cli/test_coverage)
## Local Environment Setup
1. Install pre-commit
>  brew install pre-commit
2. Install git secrets
> brew install git-secrets
>
> git-secrets --add-provider -- curl -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/ThoughtWorks-DPS/poc-resources/main/git-secrets-pattern.txt
>
> pre-commit install -f 
3. Build voltron
> go build
4. Run voltron
> ./voltron <cmd>
5. Run test files in current directory
> go test
6. Run all test files recursively
> go test ./...
7. Publishing builds
Install goreleaser
> brew install goreleaser

## Available Voltron commands

```
Usage:
  voltron [command]

Available Commands:
  hello       Call hello endpoints in API
  help        Help about any command
  version     Current version of Voltron and API

Flags:
      --config string   config file (default is voltron/.api_config.yaml)
  -h, --help            help for voltron
  -t, --toggle          Help message for toggle

Use "voltron [command] --help" for more information about a command.
```


