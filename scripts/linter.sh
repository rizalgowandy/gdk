#!/usr/bin/env bash

echo -e "\e[32mRunning: \e[33mlinter\e[0m"

golangci-lint run -c "$PWD"/.github/.golangci.yaml || exit 1
echo -e "\e[32m - golangci-lint: \e[33mpass\e[0m"
