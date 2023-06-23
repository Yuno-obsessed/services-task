# syntax=docker/dockerfile:1

FROM golang:1.20-alpine as builder

WORKDIR /cmd
COPY go.mod go.sum ./

COPY ./receiver/cmd/server/*.go ./server/
RUN CGO_ENABLED=0 GOOS=linux go build ./server/main.go

FROM builder AS run-test-stage
RUN go test -v ./...

FROM alpine:latest as image

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /cmd/ .

ENTRYPOINT [ "./receiver" ]