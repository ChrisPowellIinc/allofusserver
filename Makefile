test:
	go test ./...

start:
	go run main.go --debug

push:
	git push origin master

deploy:
	git push heroku master