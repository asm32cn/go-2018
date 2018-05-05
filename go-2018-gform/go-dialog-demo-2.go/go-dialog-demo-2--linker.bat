@echo off

set strCmd=go build -ldflags="-w -s -H windowsgui" go-dialog-demo-2.go

echo #%strCmd%
%strCmd%

pause
