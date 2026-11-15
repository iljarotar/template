VERSION := `git describe --tags`
TEMPLATE_TAG := "latest"

build:
    go build -o bin/template -ldflags="-X main.version={{ VERSION }}"

docker-build:
    docker build . --tag template:local

example:
    go run main.go -t examples/document.html -t examples/layout.html -d examples/data.json -o examples/index.html

example-docker:
    docker run --rm -v ./examples:/examples ghcr.io/iljarotar/template:{{ TEMPLATE_TAG }} -t examples/document.html -t examples/layout.html -d examples/data.json -o examples/index.html
