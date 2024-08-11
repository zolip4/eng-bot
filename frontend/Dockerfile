# Используем официальный образ Node.js
FROM node:16-alpine

# Устанавливаем tini для управления процессами
RUN apk add --no-cache tini

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем package.json и package-lock.json
COPY package*.json ./

# Устанавливаем зависимости
RUN npm install

# Копируем весь проект
COPY . .

# Экспонируем порт
EXPOSE 8080

# Используем tini и запускаем приложение
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["npm", "run", "serve"]
