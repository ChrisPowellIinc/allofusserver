test:
	go test ./...

start:
	go run main.go --debug

deploy: push
	git push heroku master

push:
	git push origin master
