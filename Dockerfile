FROM golang:1.26-alpine AS builder
WORKDIR /work
COPY . .
RUN apk add just git
RUN just build

FROM scratch
COPY --from=builder /work/bin/template /

ENTRYPOINT ["/template"]
