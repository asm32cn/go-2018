@echo off

set strCmd=go build -ldflags="-w -s" asm32-buff-sqlite3-go.go

echo #%strCmd%

%strCmd%

pause
