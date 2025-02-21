
FROM golang:1.23.5-alpine AS builder

RUN apk add --no-cache postgresql-client
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN ls -la

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/api/main.go

FROM alpine:latest
RUN apk add --no-cache postgresql-client

WORKDIR /app

COPY --from=builder /app/todo-app .


COPY migrations ./migrations


RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    rm migrate.linux-amd64.tar.gz


EXPOSE 3003

CMD ["./todo-app", "--config-file", "/config.yml", "--listener", "1"]