FROM golang:1.20-alpine

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3002

CMD [ "go","run","cmd/main.go" ]