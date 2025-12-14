.DEFAULT_GOAL := run


build:
	go build -o dev.exe
run:build
	dev.exe
