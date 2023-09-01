# build stage
FROM golang:1.20-alpine AS build

WORKDIR /opt/people-service

COPY . .

RUN go get -d -v ./...
RUN go build -o people-service -v

# final stage
FROM alpine:latest

WORKDIR /opt/people-service

COPY --from=build /opt/people-service/people-service .
