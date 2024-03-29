FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]