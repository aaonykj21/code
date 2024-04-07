FROM golang:1.22.2-alpine3.19

WORKDIR /app

COPY . .

RUN go build -o api

EXPOSE 8081

CMD [ "./api" ]