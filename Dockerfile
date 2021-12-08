FROM golang:1.17-alpine3.14 AS builder

COPY . /github.com/siteddv/pocketel_bot/
WORKDIR /github.com/siteddv/pocketel_bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/siteddv/pocketel_bot/bin/bot .
COPY --from=0 /github.com/siteddv/pocketel_bot/configs configs/

EXPOSE 80

CMD ["./bot"]