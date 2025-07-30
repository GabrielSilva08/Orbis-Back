# Usa uma imagem oficial do Golang
FROM golang:1.24

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend .

RUN go build -v -o main ./cmd

EXPOSE 8080

CMD ["./main"]