@echo off

set strCmd=go build -ldflags="-w -s" go-http-download-demo1.go

echo #%strCmd%
%strCmd%

pause