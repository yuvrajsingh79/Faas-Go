FROM golang:latest

RUN mkdir /app

RUN mkdir $GOPATH/src/grpc-server
RUN go get google.golang.org/grpc
ADD . $GOPATH/src/grpc-server
WORKDIR $GOPATH/src/grpc-server

RUN go build -o /app/main .

CMD ["/app/main"]
EXPOSE 50052