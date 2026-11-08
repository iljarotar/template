VERSION := `git describe --tags`

build:
    go build -o bin/template -ldflags="-X main.version={{ VERSION }}"

docker-build:
    docker build . --tag template:local

example:
    go run main.go -t examples/document.html -t examples/layout.html -d examples/data.json -o examples/index.html
