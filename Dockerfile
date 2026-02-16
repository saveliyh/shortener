FROM golang:1.25.7 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.env ./
COPY ./cmd ./cmd
COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/.

EXPOSE 9000


CMD ["/api"]