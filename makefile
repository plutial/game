all:
	go build -o bin/game .

run:
	go build -o bin/game . && ./bin/game

clean:
	rm -r bin
