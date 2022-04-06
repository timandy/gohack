SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64

del gohack.s /F /Q
go build -gcflags="-N -l" -o gohack.exe
go tool objdump -S gohack.exe > gohack.s
del gohack.exe /F /Q
