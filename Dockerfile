FROM golang:1.16-alpine AS builder
WORKDIR /go/src/github.com/vigo/lsvirtualenvs
COPY . .

RUN apk add --no-cache git=2.34.2-r0 \
    ca-certificates=20211220-r0 \
    && CGO_ENABLED=0 GOOS=linux \
    go build -a -installsuffix cgo -o lsvirtualenvs .

FROM alpine:3.15
COPY --from=builder /go/src/github.com/vigo/lsvirtualenvs/lsvirtualenvs /bin/lsvirtualenvs
ENTRYPOINT ["/bin/lsvirtualenvs"]
