
FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/api/main.go

FROM golang:1.23.5-alpine3.21

WORKDIR /app

COPY --from=builder /app/todo-app .
COPY config.yml ./config.yml


COPY migrations ./migrations


COPY wait-for-db.sh /app/wait-for-db.sh
RUN apk add --no-cache postgresql-client

RUN chmod +x /app/wait-for-db.sh


RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz && \
    tar -xvzf migrate.linux-amd64.tar.gz -C /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate && \
    rm migrate.linux-amd64.tar.gz


EXPOSE 3003

CMD ["./todo-app", "--config-file", "./config.yml", "--listener", "1"]