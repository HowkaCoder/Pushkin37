# Dockerfile для news_service
FROM golang:latest

WORKDIR /go/Pushkin37/news_service

COPY . .

RUN go build -o news_service

CMD ["./news_service"]

