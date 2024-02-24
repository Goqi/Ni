@echo off
go build -ldflags "-w -s" -o release/Ni.exe main.go
upx.exe -9 release/Ni.exe


@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/Ni main.go
upx.exe -9 release/Ni

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/Ni32.exe main.go
upx.exe -9 release/Ni32.exe

@echo off
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/Ni_darwin main.go
upx.exe -9 release/Ni_darwin
