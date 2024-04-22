.PHONY: test build random generate-key start

test:
	cd tests && go test -v 

build:
	go build -o ronin-random-beacon

random: build
	./ronin-random-beacon random $(ARGS)

generate-key: build
	@echo "Generating a key..."
	./ronin-random-beacon generate-key

start: build
	./ronin-random-beacon start $(ARGS)

