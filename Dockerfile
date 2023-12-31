FROM golang:alpine
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY ./ ./
ENV DSN "host='10.10.0.136' port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

# Соберём приложение
RUN go build ./cmd/crudApp
EXPOSE 8090
# Запустим приложение
CMD ["./crudApp"]

