FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go-proxy-server

FROM alpine:latest  

COPY --from=builder /go-proxy-server /go-proxy-server

EXPOSE 8080

ENTRYPOINT ["/go-proxy-server"]
