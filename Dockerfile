FROM golang:1.22.2-alpine3.19 AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o main .

FROM alpine:3.19

WORKDIR /

COPY --from=build /app/main .


EXPOSE 8000

CMD ["./main"]