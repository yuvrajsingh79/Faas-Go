FROM golang:latest

RUN mkdir /app

RUN mkdir $GOPATH/src/grpc-client
RUN go get google.golang.org/grpc
ADD . $GOPATH/src/grpc-client
WORKDIR $GOPATH/src/grpc-client

RUN go build -o /app/main .

CMD ["/app/main"]
EXPOSE 50053