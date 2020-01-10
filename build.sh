#!/bin/sh
echo "WARNING! THIS SCRIPT ARE SUPPOSED TO BE RUN FROM A LINUX OPERATING SYSTEM."

# Linux
go build -o build/wstalker-linux cmd/wstalker/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o build/wstalker-win.exe cmd/wstalker/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o build/wstalker-macos cmd/wstalker/main.go
