# Dockerfile for user-service
FROM golang:1.16

WORKDIR /app

RUN rm -f go.mod rm go.sum
RUN go mod init api-gateway
RUN go mod tidy


COPY . ./

RUN go build -o /api-gateway

EXPOSE 50051

CMD [ "/api-gateway" ]

