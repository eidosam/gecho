FROM golang:1.11-alpine as goalpine
WORKDIR /go/src/github.com/eidosam/gecho
ADD . .
RUN CGO_ENABLED=0 go build -o /gecho ./cmd/server.go

FROM alpine
RUN apk --no-cache add ca-certificates
EXPOSE 8080
COPY --from=goalpine /gecho /bin/gecho
ENTRYPOINT /bin/gecho
