# ---------- BUILD STAGE ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build backend
RUN go build -o server ./cmd

# install migrate cli
RUN wget -O migrate.tar.gz \
https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz \
 && tar -xzf migrate.tar.gz \
 && mv migrate /app/migrate

# ---------- RUNTIME ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/migrate .
COPY cmd/migrate/migrations ./migrations

EXPOSE 8080

CMD ["./server"]