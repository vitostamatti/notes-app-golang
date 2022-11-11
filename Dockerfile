# syntax=docker/dockerfile:1

FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY app ./app

RUN ls

RUN go build ./app/main.go

EXPOSE 8000

CMD [ "./main" ]