FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags "-s -w" -o avocadoro

FROM alpine:3.11
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/avocadoro .

RUN addgroup -S avocadoro --gid 1000 && adduser -S avocadoro -G avocadoro --uid 1000
RUN chown -R avocadoro /app

USER avocadoro
CMD ["./avocadoro"]
