FROM golang:latest

LABEL maintainer="Hugo Lageneste <hugo@olivia-ai.org>"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8080

CMD ["./main"]