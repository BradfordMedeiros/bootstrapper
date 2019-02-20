
bootstrapper:
	mkdir -p ./build
	go build -o ./build/bootstrapper main.go

clean:
	rm -rf ./build