run:
	go run ./cmd/forselect/main.go -- ./test/invalid.go

example:
	go build -o forselect cmd/forselect/main.go
	./forselect ./testdata
