#!/usr/bin/env bash

echo -e "\e[32mRunning: \e[33mtest.\e[0m"

echo -e "\e[32mType: \e[33munit test.\e[0m"
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  command time -f %E go test -v -failfast -race -count=1 ./... >/dev/null || exit 1
else
  go test -v -failfast -race -count=1 ./... >/dev/null || exit 1
fi
echo -e "\e[32mUnit test: \e[33msuccess.\e[0m"

echo -e "\e[32mTest: \e[33msuccess.\e[0m"
