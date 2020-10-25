#!/usr/bin/zsh

echo -e "\e[32mSetup: \e[33mstarting.\e[0m\n"

echo -e "\e[32mInstalling: \e[33mgin for live reload.\e[0m"
command -v gin 2>/dev/null || go get -v github.com/codegangsta/gin
echo ""

echo -e "\e[32mInstalling: \e[33mair for live reload.\e[0m"
command -v air 2>/dev/null || go get -v github.com/cosmtrek/air@v1.12.1
echo ""

echo -e "\e[32mInstalling: \e[33mmockgen for mock generator.\e[0m"
command -v mockgen 2>/dev/null || go get -v github.com/golang/mock/mockgen@v1.4.3
echo ""

echo -e "\e[32mInstalling: \e[33mgolangci-lint for linter.\e[0m"
command -v golangci-lint 2>/dev/null || go get -v github.com/golangci/golangci-lint/cmd/golangci-lint@1.30.0
echo ""

echo -e "\e[32mInstalling: \e[33mwire for compile time dependency injection.\e[0m"
command -v wire 2>/dev/null || go get -v github.com/google/wire/cmd/wire@v0.4.0
echo ""

echo -e "\e[32mInstalling: \e[33mswag for open api documentation.\e[0m"
command -v swag 2>/dev/null || go get -v github.com/swaggo/swag/cmd/swag@v1.6.7
echo ""

echo -e "\e[32mInstalling: \e[33mgolines for formatting long lines code.\e[0m"
command -v golines 2>/dev/null || go get -v github.com/segmentio/golines
echo ""

echo -e "\e[32mSetup: \e[33mpre-commit hook.\e[0m"
file=.git/hooks/pre-commit
cp scripts/pre-commit.sh $file
chmod +x $file
test -f $file && echo "$file exists."
echo ""

echo -e "\e[32mSetup: \e[33msuccess.\e[0m"
