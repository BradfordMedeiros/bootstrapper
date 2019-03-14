all: bootstrapper

bootstrapper:
	mkdir -p ./build
	go build -o ./build/bootstrapper ./src/main.go
	cp -r ./data ./build
clean:
	rm -rf ./build

publish-docker: docker
	docker push bradfordmedeiros/bootstrapper:0.3
docker:
	docker build -t bradfordmedeiros/bootstrapper:0.3 .
clean-docker:
	docker rmi -f bootstrapper

install: bootstrapper
	@echo install placeholder
	mkdir -p /etc/blacksmith/
	mv ./build /etc/blacksmith/bootstrapper

uninstall:
	@echo uninstall placeholder	
	rm -rf /etc/blacksmith/bootstrapper
	# should delete blacksmith if that is empty probably todo
