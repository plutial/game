all:
	export GDK_BACKEND=wayland
	mkdir -p bin
	go build -o bin/game .

run:
	export GDK_BACKEND=wayland
	mkdir -p bin
	go build -o bin/game . && ./bin/game

clean:
	rm -r bin
