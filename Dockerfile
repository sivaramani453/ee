FROM golang:1.26.1-alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o server main.go

FROM gcr.io/distroless/static-debian12
WORKDIR /
COPY --from=builder /app/server /server
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/server"]