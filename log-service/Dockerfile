FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM golang:1.23
WORKDIR /app
COPY --from=build /app/main /app/main
COPY .env /app/.env
EXPOSE 8080
CMD ["/app/main"]