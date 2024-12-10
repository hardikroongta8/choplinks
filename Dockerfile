FROM golang:1.23-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 7500
EXPOSE 3306
CMD ["go", "run", "main.go"]