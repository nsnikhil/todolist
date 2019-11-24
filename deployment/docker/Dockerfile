FROM golang:alpine

RUN apk add make

WORKDIR /todolist

COPY go.mod go.sum ./

RUN go mod download

COPY .. .

RUN make build

EXPOSE 8080

CMD ["./out/main", "serve"]