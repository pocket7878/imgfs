.PHONY: imgfs

imgfs: main.go Makefile
	go build

run: imgfs
	./imgfs
