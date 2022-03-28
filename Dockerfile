FROM golang:1.16

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/ .

CMD ["./out/go-sigv4-proxy"]