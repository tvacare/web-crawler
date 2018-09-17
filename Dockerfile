FROM golang:1.8 as builder

RUN mkdir -p $GOPATH/src/github.com/tvacare/web-crawler
WORKDIR $GOPATH/src/github.com/tvacare/web-crawler

COPY . .

RUN go get github.com/go-sql-driver/mysql
RUN go build -o main .

FROM golang:1.8

RUN mkdir -p /app

COPY --from=builder /go/src/github.com/tvacare/web-crawler .

# RUN go test ./...
CMD ["./main"]