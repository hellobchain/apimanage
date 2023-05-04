#!/bin/bash

build() {
  os=$1
  echo "build ${os} being"
  CGO_ENABLED=0 GOOS=${os} GOARCH=amd64 go build -o ApiManager-${os} cmd/main.go &&
    mkdir -p releases/ApiManager-${os}-amd64/html &&
    mv ApiManager-${os} releases/ApiManager-${os}-amd64/ApiManager &&
    cp -rp deployments/sql/db.sql deployments/config releases/ApiManager-${os}-amd64/ &&
    cp -rp html/static html/views releases/ApiManager-${os}-amd64/html

  if [ $os = linux ] || [ $os = darwin ] || [ $os = mac ]; then
    cp -rp run.sh deployments/config releases/ApiManager-${os}-amd64/deployments/config
  fi

  echo "build ${os} end"
}
case $1 in
build)
  if [[ $2 = "linux" ]]; then
    build linux
  elif [[ $2 = "windows" ]]; then
    build windows
  elif [[ $2 = "darwin" ]]; then
    build darwin
  else
      build linux
      build windows
      build darwin
  fi
  ;;
clear)
  rm -rf releases
  ;;
*)
  echo "$0 build [linux|windows|darwin|all] | clear"
  exit 4
  ;;
esac
