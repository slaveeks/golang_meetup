FROM golang:1.18-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./*.go ./

RUN go build -o /meetup

FROM alpine:3.14.1

WORKDIR /usr/app

COPY --from=build /meetup /meetup

ENTRYPOINT ["/meetup"]