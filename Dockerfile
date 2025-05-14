# Stage 1: Build
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект и собираем
COPY . .
RUN go build -o server ./cmd

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# Скопируем бинарник из стадии билда
COPY --from=builder /app/server .

EXPOSE 8080

# Запуск
CMD ["./server"]
