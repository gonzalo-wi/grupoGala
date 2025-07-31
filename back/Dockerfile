FROM golang:1.21 AS build
WORKDIR /app
COPY . .
RUN go build -o server

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]