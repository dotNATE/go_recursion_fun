build:
	go build -o bin/main

run:
	go run main.go fund.go holding.go

test:
	go test .