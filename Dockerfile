# ── 建置階段 ─────────────────────────────────────────────
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./

ARG COMMIT=dev

RUN go build \
    -ldflags "-X main.commit=${COMMIT}" \
    -o server .

# ── 執行階段 ─────────────────────────────────────────────
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
