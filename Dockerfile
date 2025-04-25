FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o devlogger ./cmd/server

EXPOSE 50051

CMD ["./devlogger"]
