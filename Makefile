.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t pocketel_bot:v0.1 .

start-container:
	docker run --name pocketel_bot -p 80:80 --env-file .env pocketel_bot:v0.1