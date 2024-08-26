#!/bin/bash

goimports -l -w $(find . -type f -name '*.go' -not -path "./.idea/*")
