build:
	go build -o bin/ ./cmd/...
test:
	go test -v ./...
run:
	go run ./cmd/...
clean:
	rm -rf bin/
generate:
	go get github.com/99designs/gqlgen@v0.17.24
	go run github.com/99designs/gqlgen generate ./...
compose-local:
	docker-compose -f docker-compose.local.yml up -d --remove-orphans
compose-local-down:
	docker-compose -f docker-compose.local.yml down
compose-local-restart:
	docker-compose -f docker-compose.local.yml restart
compose-local-logs:
	docker-compose -f docker-compose.local.yml logs -f
compose-local-logs-api:
	docker-compose -f docker-compose.local.yml logs -f api

compose-prod: 
	docker-compose -f compose.yml up -d --remove-orphans
	${MAKE} run
	
compose-prod-down:
	docker-compose -f compose.yml down
compose-prod-restart:
	docker-compose -f compose.yml restart
compose-prod-logs:
	docker-compose -f compose.yml logs -f

.PHONY: build test clean run compose-local compose-local-down compose-local-restart compose-local-logs compose-local-logs-api compose-prod compose-prod-down compose-prod-restart compose-prod-logs
