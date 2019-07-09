FROM golang:latest as builder
WORKDIR /go/src/socks5
COPY server.go .
RUN go get && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./socks5

FROM curve25519xsalsa20poly1305/openvpn-tunnel:latest

COPY socks5-entrypoint.sh /usr/local/bin/
COPY --from=builder /go/src/socks5/socks5 /usr/local/bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN chmod +x /usr/local/bin/socks5-entrypoint.sh

ENV SOCKS5_UP   ""
ENV OPENVPN_UP  /usr/local/bin/socks5-entrypoint.sh
