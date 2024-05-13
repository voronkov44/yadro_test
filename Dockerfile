FROM golang:1.22-alpine as builder
WORKDIR /opt
COPY . .
RUN go build -o /main main.go
# Финальный этап, копируем собранное приложение
FROM alpine:3.17
WORKDIR /opt
COPY --from=builder main /opt/main
COPY test /opt/test/

ENTRYPOINT ["/opt/main"]
