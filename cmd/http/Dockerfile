FROM golang:1.17 as builder

WORKDIR /src
COPY . .

RUN go mod download
ENV CGO_ENABLED=0

RUN go test ./...
RUN mkdir -p /app
RUN go build -ldflags="-w -s" -o /app ./cmd/http

FROM scratch
COPY --from=builder /app /app
CMD ["/app/http"]
ENV GIN_MODE=release
