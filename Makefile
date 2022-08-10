# Start server
start: 
	air

migrate-schema:
	go run ./cmd/migration/main.go

run:
	go run ./cmd/app/main.go