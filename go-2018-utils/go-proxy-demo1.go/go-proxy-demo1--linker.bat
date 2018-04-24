@echo off

set strCmd=go build -ldflags="-w -s" go-proxy-demo1.go

echo #%strCmd%
%strCmd%

pause