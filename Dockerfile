FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o my-app

EXPOSE 3300

CMD ./my-app
