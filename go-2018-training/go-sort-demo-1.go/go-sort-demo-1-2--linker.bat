@echo off

set strCmd=go build -ldflags="-w -s" go-sort-demo-1-2.go

echo #%strCmd%
%strCmd%

pause
