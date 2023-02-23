FROM golang:1.19-alpine3.17 as builder
RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go test ./... -v

RUN go build -o ./out/app .



FROM alpine:3.17 
RUN apk add ca-certificates

COPY --from=builder /out/app /app

EXPOSE 8080

CMD ["/app"]

