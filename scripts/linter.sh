#!/usr/bin/env bash

echo -e "\e[32mRunning: \e[33mlinter.\e[0m"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  command time -f %E golangci-lint run -c "$PWD"/.golangci.yaml || exit 1
else
  golangci-lint run -c "$PWD"/.golangci.yaml || exit 1
fi

echo -e "\e[32mLinter: \e[33msuccess.\e[0m"
