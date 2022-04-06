export CGO_ENABLED=0
export GOOS=linux
export GOARCH=arm
export GOARM=6
go test -v
