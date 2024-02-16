FROM golang:1.22-alpine3.19
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./src/app

EXPOSE 8080

CMD ["./src/app"]