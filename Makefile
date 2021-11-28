.PHONY:
.SILENT:
.DEFAULT_GOAL := run


build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t ferris:0.09 .

start-container:
	docker run --env-file .env -p 80:80 ferris:0.09

test:
	go test github.com/Aserose/ferrisWheel/internal/web
	go test github.com/Aserose/ferrisWheel/internal/storage/boltDB
	go test github.com/Aserose/ferrisWheel/internal/web/dataSource/geocode
	go test github.com/Aserose/ferrisWheel/internal/web/dataSource/vk
	go test github.com/Aserose/ferrisWheel/internal/web/dataSource
	go test github.com/Aserose/ferrisWheel/internal/web/tg