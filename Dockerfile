# Используем официальный образ Go для сборки
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /build

# Копируем файлы go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все файлы исходного кода в рабочую директорию
COPY . .

# Сборка приложения из директории cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./cmd/server

# Создаем минимальный образ для продакшена на базе scratch
FROM scratch

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /go/bin/app /app

# Указываем точку входа для контейнера
ENTRYPOINT ["/app"]

# Указываем порт, на котором работает приложение
EXPOSE 80
