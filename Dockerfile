#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/fetch main.go

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/fetch /fetch
ENTRYPOINT ["./fetch"]
CMD ["https://www.google.com", "https://www.autify.com"]
