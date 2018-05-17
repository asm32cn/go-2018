@echo off

set strCmd=go build -ldflags="-w -s -H windowsgui" go-dialog.go

echo #%strCmd%
%strCmd%

pause
