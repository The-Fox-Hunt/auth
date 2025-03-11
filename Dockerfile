# Используем минимальный образ Go
FROM golang:1.24-alpine AS builder

# Добавляем сертификаты для HTTPS (например, для PostgreSQL)
RUN apk add --no-cache ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник
RUN go build -o /app/auth cmd/service/main.go

# Финальный минимальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарник из билдера
COPY --from=builder /app/auth /app/auth

# Делаем бинарник исполняемым (на всякий случай)
RUN chmod +x /app/auth

# Запускаем приложение
CMD ["/app/auth"]
