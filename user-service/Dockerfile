# Dockerfile for user-service
FROM golang:1.16

WORKDIR /app

RUN rm -f go.mod rm go.sum
RUN go mod init user-service
RUN go mod tidy


COPY . ./

RUN go build -o /user-service

EXPOSE 50051

CMD [ "/user-service" ]
