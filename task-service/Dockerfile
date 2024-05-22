# Dockerfile for task-service
FROM golang:1.16

WORKDIR /app

RUN rm -f go.mod rm go.sum
RUN go mod init task-service
RUN go mod tidy


COPY . ./

RUN go build -o /task-service

EXPOSE 50051

CMD [ "/task-service" ]
