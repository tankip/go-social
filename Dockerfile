# base image for the golang app
FROM golang:1.18
WORKDIR /go/src/github.com/go-social
COPY . .
# Download necessary Go modules
RUN go mod download

RUN go build -o main .
CMD ["./main"]
