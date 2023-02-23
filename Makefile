target: api 
.PHONY: build test start run generate  compose-prod compose-prod-down compose-prod-logs compose-prod-restart compose-local compose-local-down compose-local-logs compose-local-restart

api: build compose-local test start
frontend:
	cd app && pnpm run dev
build:
	go build -o bin/ ./cmd/...
test:
	go test -v ./...
start:
	./bin/api
run:
	go run ./cmd/...
clean:	compose-local-down
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

compose-prod: 
	docker-compose -f compose.yml up -d --remove-orphans
	${MAKE} run
	
compose-prod-down:
	docker-compose -f compose.yml down
compose-prod-restart:
	docker-compose -f compose.yml restart
compose-prod-logs:
	docker-compose -f compose.yml logs -f

