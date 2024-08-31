#!/bin/bash

GIT_COMMIT=$(git rev-parse --short HEAD)
GIT_LAST_COMMIT_INFO=$(git log -1 --format="%h %cd")

go build -ldflags="-X 'main.LastCommitInfo=$GIT_LAST_COMMIT_INFO'" -o jsonmd