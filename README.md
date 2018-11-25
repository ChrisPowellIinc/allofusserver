# All of Us

## Start the database
run: `mongod` in your terminal.

Open another terminal and run: `mongo` to start the mongo shell. In this shell create the database using this command: `use allofus` where `allofus` is the name of the database we will be using for this application

## Start the app.
you can start the app by running this command on your terminal. Make sure this command is run on the root folder of the app.
`$ go run main.go`

To run tests, run this command:
`$ go test ./...`

# about the app

This app is the server for the allofus social network. The mobile app will be authenticated by this app through the api provided by this server.