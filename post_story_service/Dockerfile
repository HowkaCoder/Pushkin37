# Dockerfile для post_story_service
FROM golang:latest

WORKDIR /go/Pushkin37/post_story_service

COPY . .

RUN go build -o post_story_service

CMD ["./post_story_service"]

