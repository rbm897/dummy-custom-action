FROM golang:1.20-alpine

WORKDIR /app

COPY main.go go.mod /app/

# CMD ['/bin/go', 'run', 'main.go']

RUN go build -o /app/b



CMD ["/app/b"]