CGO_ENABLED="0"

GOOS="darwin"
GOARCH="amd64"
go build -o tskssh_${GOOS}_${GOARCH}

GOOS="linux"
GOARCH="amd64"
go build -o tskssh_${GOOS}_${GOARCH}
