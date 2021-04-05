### Note
`server.debug`, `server.cov` and `server-coverage.html` are included into repo for demo purposes, remove these files first if you want to generate them on your own.

# Getting coverage in acceptance tests

1) Create [special test file](src/server/server_test.go) that brings up your application with timeout long enough to run test suite.
2) Compile instrumenter binary
```shell
GOPATH=$HOME/simple-http-server/ go test -c -covermode=set -coverpkg=server -o server.debug server
```
3) Run compiled binary and test suite
```shell
./server.debug -test.coverprofile=server.cov
# Do some testing
```
4) Open coverage report in default browser
```shell
GOPATH=$HOME/simple-http-server/ go tool cover -html=server.cov
```
Or generate HTML report file
```shell
GOPATH=$HOME/simple-http-server/ go tool cover -html=server.cov -o server-coverage.html
```
