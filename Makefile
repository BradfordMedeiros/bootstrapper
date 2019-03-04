all: bootstrapper

bootstrapper:
	mkdir -p ./build
	go build -o ./build/bootstrapper ./src/main.go
	cp -r ./data ./build
clean:
	rm -rf ./build

publish-docker: docker
	docker push bradfordmedeiros/bootstrapper:0.2
docker:
	docker build -t bradfordmedeiros/bootstrapper:0.2 .
clean-docker:
	docker rmi -f bootstrapper
