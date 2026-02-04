FROM golang:1.25-alpine AS builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /go-cv-app cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache ca-certificates

# Copy the binary
COPY --from=builder /go-cv-app .

EXPOSE 8080
CMD ["./go-cv-app"]