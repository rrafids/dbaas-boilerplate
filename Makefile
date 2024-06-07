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


git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:rrafids/dbaas-boilerplate.git
git push -u origin main