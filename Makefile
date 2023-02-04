tidy:
	go mod tidy

run:
	go run main.go

test:
	go test ./... --cover
	
mock: 
	~/bin/mockery --all --keeptree