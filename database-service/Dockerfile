
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /database-service cmd/main.go

EXPOSE 8084

CMD [ "/database-service" ]
