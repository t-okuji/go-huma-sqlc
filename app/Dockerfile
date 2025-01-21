FROM golang:1.23.4-alpine AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]

FROM golang:1.23.4-alpine AS builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd/app

FROM scratch AS prod
COPY --from=builder /bin/app /app
CMD [ "/app" ]