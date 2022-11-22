@echo off
go build -ldflags "-w -s" -o release/En.exe main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/En32.exe main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/En main.go

@echo off
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/En_darwin main.go
