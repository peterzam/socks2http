ARG GOLANG_VERSION="1.18"

FROM golang:$GOLANG_VERSION-alpine as builder
RUN apk --no-cache add git
RUN git clone --branch dev https://codeberg.org/peterzam/socks2http.git
WORKDIR /go/socks2http
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./s2h

FROM scratch
COPY --from=builder /go/socks2http/s2h /
ENTRYPOINT ["/s2h"]