FROM golang:1.23-alpine AS builder

COPY . /github.com/erikqwerty/auth/
WORKDIR /github.com/erikqwerty/auth/

RUN go mod download
RUN go build -o ./bin/auth_server cmd/auth/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/erikqwerty/auth/bin/auth_server .

CMD ["./auth_server"]