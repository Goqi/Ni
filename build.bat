@echo off
go build -ldflags "-w -s" -o release/Cell.exe main.go
upx.exe -9 release/Cell.exe


@echo off
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/Cell main.go
upx.exe -9 release/Cell

@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=386
go build -ldflags "-w -s" -o release/Cell32.exe main.go
upx.exe -9 release/Cell32.exe

@echo off
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags "-w -s" -o release/Cell_darwin main.go
upx.exe -9 release/Cell_darwin
