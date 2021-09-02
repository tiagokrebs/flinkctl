# Flinkctl

## Get started with spf13/cobra
```
go get -u github.com/spf13/cobra
```

Next, start your application structure with cobra init

```
# install
export GOMODULE111=on
go install github.com/spf13/cobra/cobra@latest

# create app folder
mkdir -p flinkctl
cd flinkctl

# create cobra structure
cobra init --pkg-name github.com/tiagokrebs/flinkctl

# add cobra commands
cobra add serve
cobra add config
cobra add create -p 'configCmd'

go mod init github.com/tiagokrebs/flinkctl
go mod tidy
go mod vendor

go run main.go

go build

./flinkctl
```