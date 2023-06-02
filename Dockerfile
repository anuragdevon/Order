FROM golang:alpine

WORKDIR /order

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./order ./main.go

EXPOSE 8030

CMD ./order
