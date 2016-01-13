FROM golang:1.5

WORKDIR /go/src/github.com/hairqles/queue-proxy

COPY . /go/src/github.com/hairqles/queue-proxy

RUN go build -o queue-proxy

EXPOSE 80

CMD ["./queue-proxy"]
