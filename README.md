# Flinkctl

## Design Guideline
[Package Oriented Design - Ardan Labs](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html)

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

```

Then, do a simple test
```
go mod init github.com/tiagokrebs/flinkctl
go mod tidy
go mod vendor

go run main.go

go build

./flinkctl
```