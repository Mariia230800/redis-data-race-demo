FROM golang:1.24-alpine

WORKDIR /app

# Копируем go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Сборка бинаря
RUN go build -o cron-service ./cmd/cron-service

# Открываем порты, если нужны (для cron скорее не нужен)
# EXPOSE 50051

# Команда запуска
CMD ["./cron-service"]