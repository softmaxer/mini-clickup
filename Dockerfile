from golang:1.21-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o miniclickup ./cmd/main.go

CMD ["./miniclickup"]
