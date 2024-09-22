run:
	go run ./cmd/forselect/main.go -- ./test/invalid.go

example:
	go build cmd/forselect/main.go
	./main ./test
