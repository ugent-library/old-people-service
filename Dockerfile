# build stage
FROM golang:1.20-alpine AS build

WORKDIR /opt/people-service

COPY . .

RUN go get -d -v ./...
RUN go build -o people-service -v

# final stage
FROM alpine:latest

ARG SOURCE_BRANCH
ARG SOURCE_COMMIT
ARG IMAGE_NAME
ENV SOURCE_BRANCH $SOURCE_BRANCH
ENV SOURCE_COMMIT $SOURCE_COMMIT
ENV IMAGE_NAME $IMAGE_NAME

WORKDIR /opt/people-service

COPY --from=build /opt/people-service/people-service .
COPY --from=build /opt/people-service/public public
CMD mkdir -p api/v1
COPY --from=build /opt/people-service/api/v1/openapi.yaml /opt/people-service/api/v1/
