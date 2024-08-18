LOCAL_BIN:=${CURDIR}/bin
export GOBIN:=${LOCAL_BIN}

# Setup

lint:
	${LOCAL_BIN}/golangci-lint run --config=.golangci.yaml ./...

bin-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Test

export TEST_PG_PORT:=5433
export TEST_PG_DSN:=postgres://postgres@localhost:${TEST_PG_PORT}/postgres?sslmode=disable

test:
	make test-env-up
	go test -race ./...
	make test-env-down
	
test-env-up:
	make test-env-down
	docker-compose -f ./docker-compose.test.yaml up -d --wait
	${LOCAL_BIN}/goose -dir db/migrations postgres "${TEST_PG_DSN}" up

test-env-down:
	docker-compose -f ./docker-compose.test.yaml down --volumes

# Local

export LOCAL_PG_PORT:=5434
export LOCAL_PG_DSN:=postgres://postgres@localhost:${LOCAL_PG_PORT}/postgres?sslmode=disable

local:
	make local-env-up
	go run ./cmd/reddneck PG_DSN=${LOCAL_PG_DSN}
	make local-env-down

local-env-up:
	make local-env-down
	docker-compose -f ./docker-compose.local.yaml up -d --wait
	${LOCAL_BIN}/goose -dir db/migrations postgres "${LOCAL_PG_DSN}" up

local-env-down:
	docker-compose -f ./docker-compose.local.yaml down --volumes

