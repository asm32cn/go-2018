@echo off

set strCmd=go build -ldflags="-w -s -H windowsgui" go-gform-demo-1.go
::set strCmd=go build -ldflags="-w -s" go-gform-demo-1.go

echo #%strCmd%
%strCmd%

pause
