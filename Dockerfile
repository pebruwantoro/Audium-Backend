FROM golang:1.22-alpine AS build

RUN apk update && apk add bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/backgroundsound ./backgroundsound
COPY --from=build /app/audio ./audio
COPY --from=build /app/server .

EXPOSE 7070

CMD ["./server"]