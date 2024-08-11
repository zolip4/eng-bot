# Используем официальный образ Golang
FROM golang:1.20-alpine

# Устанавливаем необходимые утилиты
RUN apk add --no-cache git screen

# Создаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Устанавливаем air для автоматической перекомпиляции кода при изменении
RUN go install github.com/cosmtrek/air@v1.42.0

# Создаем директорию для хранения временных файлов
RUN mkdir -p /app/tmp

# Добавляем .air.toml
COPY .air.toml .

# Команда запуска air через screen
CMD screen -dmS air_session air && tail -f /dev/null
