FROM golang:1.22 AS BUILDER

WORKDIR /app

COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build -o /app/goApiJwt .

FROM golang:1.22-alpine as runner

COPY --from=BUILDER /app/goApiJwt /

EXPOSE 8080

ENTRYPOINT ["/goApiJwt"]