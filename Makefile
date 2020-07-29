.PHONY: default

default:
	rm -r templates/*
	go run main.go -y yaml -g -o templates -s -f README.md
