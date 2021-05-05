#!/bin/bash

go mod download
grnc-yaml-bind
go test ./... -coverprofile=coverage.out