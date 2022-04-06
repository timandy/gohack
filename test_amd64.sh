export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64
go test -v -race
