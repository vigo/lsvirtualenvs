FROM golang:alpine AS builder
WORKDIR /go/src/github.com/vigo/lsvirtualenvs
COPY . .
RUN apk add --no-cache git
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/vigo/lsvirtualenvs/main /app
