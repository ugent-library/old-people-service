# build stage
FROM golang:alpine AS build

WORKDIR /opt/person-service

COPY . .

RUN go get -d -v ./...
RUN go build -o person-service -v

# final stage
FROM alpine:latest

WORKDIR /opt/person-service

COPY --from=build /opt/person-service/person-service .
