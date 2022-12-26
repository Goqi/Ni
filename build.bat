@echo off
go build -ldflags "-w -s" -o release/Ernuclei.exe main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/Ernuclei32.exe main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/Ernuclei main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/Ernuclei_darwin main.go
