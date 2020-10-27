## Local Environment Setup
1. Install pre-commit
>  brew install pre-commit
2. Install git secrets
> brew install git-secrets
> git-secrets --add-provider -- curl https://raw.githubusercontent.com/ThoughtWorks-DPS/poc-resources/main/git-secrets-pattern.txt
3. Build poc-va-cli
> go build
4. Run poc-va-cli
> ./poc-va-cli <cmd>
5. Run test files in current directory
> go test
6. Run all test files recursively
> go test ./...
7. Publishing builds
Install goreleaser
> brew install goreleaser
