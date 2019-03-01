all: bootstrapper

bootstrapper:
	mkdir -p ./build
	go build -o ./build/bootstrapper ./src/main.go
	cp -r ./data ./build
clean:
	rm -rf ./build

docker:
	docker build -t bootstrapper .
clean-docker:
	docker rmi -f bootstrapper
