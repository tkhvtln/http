FROM golang:1.22.5
WORKDIR /app
COPY server/server.go /app/
RUN go build -o server server.go
CMD ["./server"]