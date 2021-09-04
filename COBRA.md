# [Cobra](https://github.com/spf13/cobra)
Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files.

Cobra is used in many Go projects such as [Kubernetes](http://kubernetes.io/),
[Hugo](https://gohugo.io), and [Github CLI](https://github.com/cli/cli) to
name a few. [This list](./projects_using_cobra.md) contains a more extensive list of projects using Cobra.


## Get started with cobra
[User Guide](https://github.com/spf13/cobra/blob/master/user_guide.md)

```
go get -u github.com/spf13/cobra
```

Next, start your application structure with cobra init.
```
# install
export GOMODULE111=on
go install github.com/spf13/cobra/cobra@latest

# create app folder
mkdir -p myapp
cd myapp

# create cobra structure
cobra init --pkg-name github.com/myuser/myapp

# add cobra commands
cobra add serve
cobra add config
cobra add create -p 'configCmd'

```

Then, do a simple test.
```
go mod init github.com/myuser/myapp
go mod tidy
go mod vendor

go run main.go config
go run main.go config create
```

Build!
```
go build

./myapp config
./myapp config create
```