run-http:
	go run cmd/http/main.go
run-cli:
	gor run cmd/cli/main.go
migrate-up:
	migrate -database 'postgres://rrafids:rrafids@localhost:5432/cuspay?sslmode=disable' -path=./migration up
migrate-down:
	migrate -database 'postgres://rrafids:rrafids@localhost:5432/cuspay?sslmode=disable' -path=./migration down
test:
	go test ./... -coverprofile coverage.out
	rm -rf coverage.out