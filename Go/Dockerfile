FROM golang:alpine

WORKDIR /build


COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . . 

RUN go build -o main .

RUN go test -v ./...


WORKDIR /dist

RUN cp /build/main .

EXPOSE 8080

CMD ["/dist/main"]
