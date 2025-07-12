FROM golang:1.22.4

WORKDIR /oauth
COPY . .

# Ensure logs dir exists with write permission
RUN mkdir -p /oauth/logs && chmod -R 777 /oauth/logs

RUN go build -o oauth cmd/main.go

CMD ["./oauth"]
