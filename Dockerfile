# ---- Build stage ----
FROM golang:1.24-bookworm AS builder
WORKDIR /src

ENV CGO_ENABLED=1 GO111MODULE=on

# For go-sqlite3 build
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential ca-certificates && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/app ./...

# ---- Runtime stage ----
FROM debian:bookworm-slim
WORKDIR /app

# minimal runtime + wget for healthcheck
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates wget && rm -rf /var/lib/apt/lists/*

COPY --from=builder /bin/app /app/app

# Data dir for SQLite (we'll mount a volume here)
RUN mkdir -p /app/data

ENV BAKERY_TOKEN=change-me
ENV DB_FILE=/app/data/bakes.db

EXPOSE 8080
CMD ["/app/app"]