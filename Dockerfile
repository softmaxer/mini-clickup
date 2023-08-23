FROM golang:1.21.0-bookworm as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 go build -o miniclickup ./cmd/main.go

CMD ["./miniclickup"]
