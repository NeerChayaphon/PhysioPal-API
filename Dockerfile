FROM golang:1.19 as build

ENV GO111MODULE=on

WORKDIR /app/server

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build main.go

FROM alpine:latest as server

WORKDIR /app/server

COPY --from=build /app/server/main .

ARG MONGODB_STAGING_URI 
ARG REDIS_LOCAL_URI
ARG REDIS_LOCAL_PASSWORD

ENV MONGODB_STAGING_URI=$MONGODB_STAGING_URI
ENV REDIS_LOCAL_URI=$REDIS_LOCAL_URI
ENV REDIS_LOCAL_PASSWORD=$REDIS_LOCAL_PASSWORD

RUN chmod +x ./main

CMD ["./main"]