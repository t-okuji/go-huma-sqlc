FROM golang:1.23.4-alpine as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]
