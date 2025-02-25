VERSION = v1.0

compile:
	go build -o ./bin/$(VERSION)/0xBlade ./src/;
	cd ./bin/$(VERSION)/ && tar -czvf 0xBlade_$(VERSION).tar.gz *
