FROM golang:1.17.3-alpine3.14 AS builder

RUN go version

COPY . /github.com/Aserose/FerrisWheel/
WORKDIR /github.com/Aserose/FerrisWheel/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/Aserose/FerrisWheel/.bin/bot .

EXPOSE 80

CMD ["./bot"]