FROM golang:latest as builder
WORKDIR /go/src/socks5
COPY server.go .
RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./socks5

FROM curve25519xsalsa20poly1305/openvpn-tunnel:latest

COPY --from=builder /go/src/socks5/socks5 /usr/local/bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV OPENVPN_UP socks5
