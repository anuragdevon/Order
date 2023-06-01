FROM golang:alpine

WORKDIR /order_svc

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./order_svc ./main.go

EXPOSE 8030

CMD ./order_svc
