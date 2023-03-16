FROM golang:1.20

WORKDIR /go/src/github.com/ahmetoguzhanengin/RestAPI

COPY . .

RUN go get -v

RUN go build -o main .

CMD ["./main"]