FROM golang
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go test ./...
RUN go build -o data-sync-manager ./orchestrator
CMD ["./data-sync-manager", "--config=config.yaml"]
