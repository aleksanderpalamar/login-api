FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o login-api .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/login-api .

EXPOSE 3000

CMD ["./login-api"]