# Build stage
FROM golang:1.21 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/ghc ./main.go

# Run stage
FROM alpine:3.14.10
EXPOSE 8080

COPY --from=builder /app/bin/ghc .

# Increase GC percentage and limit the number of OS threads
ENV GOGC 1000
ENV GOMAXPROCS 3

CMD ["/ghc"]

