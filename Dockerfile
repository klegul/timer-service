FROM golang:1.23-alpine as build
WORKDIR /opt
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./timer-service

FROM alpine:latest as run
WORKDIR /opt
COPY --from=build /opt/timer-service .
CMD ["/opt/timer-service"]