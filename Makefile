compile:
	@echo Enter Version Number:; \
	read VERSION;

	go build -o ./bin/$$(VERSION)/0xBlade/ ./src/;
	cd ../bin/$$(VERSION)/;
	zip 0xBlade.zip 0xBlade;
