#!/usr/bin/env bash

echo -e "\e[32mRunning:\e[33m setup.\e[0m\n"

echo -e "\e[32mInstalling:\e[33m air for live reload.\e[0m"
command -v air 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin # v1.42.0
echo ""

echo -e "\e[32mInstalling:\e[33m mockgen for mock generator.\e[0m"
command -v mockgen 2>/dev/null || GO111MODULE=off go get -v github.com/golang/mock/mockgen # v1.6.0
echo ""

echo -e "\e[32mInstalling:\e[33m golangci-lint for linter.\e[0m"
command -v golangci-lint 2>/dev/null || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin 1.52.2
echo ""

echo -e "\e[32mInstalling:\e[33m gomodifytags for generating tags.\e[0m"
command -v gomodifytags 2>/dev/null || go install -v github.com/fatih/gomodifytags@v1.16.0
echo ""

echo -e "\e[32mSetup:\e[33m pre-commit hook.\e[0m"
file=.git/hooks/pre-commit
cp scripts/pre-commit.sh $file
chmod +x $file
test -f $file && echo "$file exists."
echo ""

echo -e "\e[32mSetup:\e[33m success.\e[0m"
