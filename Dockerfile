FROM golang:1.22-alpine as builder
WORKDIR /Users/drop-/GolandProjects/awesomeProject/yadro_test
COPY . .
RUN go mod download
RUN go build -o /main main.go
# Финальный этап, копируем собранное приложение
FROM alpine:3.17
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]