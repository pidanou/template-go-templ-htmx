FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go



FROM alpine:3.19

COPY --from=builder /app/server /app/server

CMD ["/app/server"]


