FROM golang:1.23-alpine as build
WORKDIR /opt
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./anny_viz

FROM alpine:latest as run
WORKDIR /opt
COPY --from=build /opt/anny_viz .
CMD ["/opt/anny_viz"]