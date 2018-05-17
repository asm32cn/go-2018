@echo off

set strCmd=go build -ldflags="-w -s" CopyFile.go

echo #%strCmd%
%strCmd%

pause
