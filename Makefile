migrate:
	liquibase update --changeLogFile sql/changelog.sql --url "${DATABASE_URL}"

dev:
	wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run cmd/main.go

tailwind-watch:
	cd static/js && npx tailwindcss -i ../css/input.css -o ../css/style.css --watch

tailwind-build:
	cd static/js && npx tailwindcss -i ../css/input.css -o ../css/style.css --minify

build:
	templ generate
	make tailwind-build
	go build cmd/main.go
