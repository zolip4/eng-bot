FROM golang:1.22-alpine AS builder

WORKDIR /engbot/app

RUN apk add --no-cache nodejs npm

RUN go install github.com/air-verse/air@latest

COPY ./app/ /engbot/app

RUN go mod tidy

WORKDIR /engbot/resources

COPY ./resources/ /engbot/resources/

RUN npm install

RUN echo $PATH
RUN npm install --yes
RUN npm run build

WORKDIR /engbot/app

EXPOSE 80

CMD ["air"]

# CMD ["./main"]
