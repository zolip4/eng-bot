# Указание базового образа с Go 1.22
FROM golang:1.22-alpine AS builder

# Установка рабочего каталога
WORKDIR /app

# Установка необходимых пакетов, включая Node.js и npm
RUN apk add --no-cache nodejs npm

# Установка Air с правильным импортным путем
RUN go install github.com/air-verse/air@latest

# Копирование файлов бэкенда
COPY ./app/ .

# Установка зависимостей Go
RUN go mod tidy

# Копирование файлов фронтенда
COPY ./resources/ ./resources/

# Сборка фронтенда
WORKDIR /app/resources

# Установка зависимостей фронтенда
RUN npm install

# Сборка React-приложения
RUN npm run build

# Возврат к рабочей директории для Go
WORKDIR /app

# Экспозиция порта 80
EXPOSE 80

# Использование Air для разработки
CMD ["air"]

# В продакшене можно использовать:
# CMD ["./main"]
