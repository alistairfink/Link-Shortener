FROM golang:1.12.6 as builder
LABEL maintainer="AlistairFink <alistairfink@gmail.com>"

WORKDIR /go/src/github.com/alistairfink/Link-Shortener
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/Link-Shortener .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/Link-Shortener .
COPY --from=builder /go/src/github.com/alistairfink/Link-Shortener/Config.json .

EXPOSE 41692

CMD ["./Link-Shortener"] 