VERSION := `git describe --tags`

build:
    go build -o bin/template -ldflags="-X main.version={{ VERSION }}"
