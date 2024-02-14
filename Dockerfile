FROM golang:1.22-alpine3.19
WORKDIR /app

COPY . .

RUN go build -o ./src/app

EXPOSE 8080

CMD ["./src/app"]