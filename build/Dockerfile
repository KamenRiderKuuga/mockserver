FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
ENV GOPROXY=https://goproxy.io
RUN go mod tidy && go mod vendor && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags release -a -installsuffix cgo -o mockserver ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/mockserver .
ENTRYPOINT ["./mockserver"]