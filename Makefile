build:
	go build

run:
	go build && ./pye serve

dev:
	go build && ./pye serve --debug