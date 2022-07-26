FROM golang:1.16-alpine

WORKDIR /app

COPY ./ .

RUN go mod download

RUN go build -o /parties-service

EXPOSE 8080

CMD [ "/parties-service" ]
