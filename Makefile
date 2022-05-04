SHELL=/bin/sh

run:
	nodemon --exec "go run" --signal SIGTERM main.go