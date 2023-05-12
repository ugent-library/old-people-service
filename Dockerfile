# build stage
FROM golang:alpine AS build

WORKDIR /opt/people

COPY . .

RUN go get -d -v ./...
RUN go build -o people -v

# final stage
FROM alpine:latest

WORKDIR /opt/people

COPY --from=build /opt/people/people .
