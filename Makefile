.PHONY: default

default:
	go run main.go -y yaml -g -o templates -s -f README.md
