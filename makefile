all:
	go build -o bin/game **/*.go

run:
	go build -o bin/game **/*.go; ./bin/game
