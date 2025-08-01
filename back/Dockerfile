FROM golang:1.23 AS build
WORKDIR /app
RUN apt-get update && apt-get install -y gcc libc-dev libsqlite3-dev
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o server

FROM debian:bookworm-slim
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends \
    libsqlite3-0 \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*
COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]