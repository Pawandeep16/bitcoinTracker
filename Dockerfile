FROM golang:1.21 as builder
WORKDIR /app
COPY go.mod go.mod ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk add --no-cache tzdata

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 9090
CMD ["./main"]