# Устанавливаем базовый образ с Go
FROM golang:1.23.2-alpine AS builder
# Устанавливаем рабочую директорию в корне контейнера
WORKDIR /go/src/app

# Копируем все файлы в контейнер
COPY . .

# Загружаем зависимости и собираем приложение
RUN go mod download && go build -o app cmd/main.go

# Используем минимальный образ для запуска
FROM alpine:latest
WORKDIR /root

# Копируем скомпилированное приложение и все необходимые файлы
COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/docs ./docs
COPY .env .

# Запускаем приложение
CMD ["./app"]