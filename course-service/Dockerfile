FROM golang:1.18-alpine

WORKDIR /app
COPY . /app/
RUN go mod init && go mod tidy

CMD ["go", "run", "main.go"]
