FROM golang:1.22.2-alpine3.19 AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o main .

FROM alpine:3.19

WORKDIR /root

COPY --from=build /app/main .

CMD ["./main"]