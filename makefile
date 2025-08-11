all:
	mkdir -p bin
	go build -o bin/game .

run:
	mkdir -p bin
	go build -o bin/game . && ./bin/game

clean:
	rm -r bin
