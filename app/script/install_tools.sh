go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@v2.1.2
go install github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949
go install github.com/cweill/gotests/gotests@v1.6.0
go install github.com/fatih/gomodifytags@v1.16.0
go install github.com/josharian/impl@v1.1.0
go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
go install github.com/go-delve/delve/cmd/dlv@v1.7.2
GOBIN=/tmp/ go install github.com/go-delve/delve/cmd/dlv@v1.7.2 && mv /tmp/dlv $GOPATH/bin/dlv-dap
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
go install golang.org/x/tools/gopls@v0.7.4
go install entgo.io/ent/cmd/ent@v0.8.0
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.3
go install github.com/golang/mock/mockgen@v1.6.0
