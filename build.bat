cd ../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=windows
cd src/demo/gopherlua_test
go build -v -ldflags="-s -w"