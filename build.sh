CGO_ENABLED="0"

GOOS="darwin"
GOARCH="amd64"
go build -o bin/tskssh_${GOOS}_${GOARCH} cmd/tskssh.go

GOOS="linux"
GOARCH="amd64"
go build -o bin/tskssh_${GOOS}_${GOARCH} cmd/tskssh.go
