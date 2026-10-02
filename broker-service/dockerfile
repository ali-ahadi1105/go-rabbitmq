#builder
FROM golang:1.26-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o broker ./cmd/api

#run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/broker .
EXPOSE 80
CMD [ "./broker" ]