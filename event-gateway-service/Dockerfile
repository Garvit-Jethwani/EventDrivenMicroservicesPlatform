
FROM golang:1.18-alpine
WORKDIR /app
COPY . .
RUN go build -o /event-gateway
EXPOSE 50051
CMD ["/event-gateway"]
