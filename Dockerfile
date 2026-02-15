FROM golang:1.25.7 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY ./database ./database

RUN CGO_ENABLED=0 GOOS=linux go build -o /api

EXPOSE 9000

CMD ["/api"]