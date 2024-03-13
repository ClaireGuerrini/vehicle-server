FROM golang:1.22
COPY ./ /app
WORKDIR /app
RUN go build ./cmd/server
ENTRYPOINT ["/app/server"]
