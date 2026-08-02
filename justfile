VERSION := `git describe --tags`

build:
    go build -o bin/template -ldflags="-X github.com/iljarotar/template/main.version={{ VERSION }}"
