.PHONY:	bench
bench:
	@go test -bench=. -benchmem  ./...

.PHONY:	ut
ut:
	@go test -tags=goexperiment.arenas -race ./...

.PHONY:	setup
setup:
	@sh ./.script/setup.sh

.PHONY:	fmt
fmt:
	@sh ./.script/goimports.sh

.PHONY:	lint
lint:
	@golangci-lint run -c .golangci.yml

.PHONY: tidy
tidy:
	@go mod tidy -v

.PHONY: check
check:
	@$(MAKE) fmt
	@$(MAKE) tidy

.PHONY: mock
mock:
	@mockgen -source=./ratelimiter/types.go -package=mocks -destination=./ratelimiter/mocks/types.mock.go
	@mockgen -package=redismocks -destination=./ratelimiter/redismocks/cmdable.mock.go github.com/redis/go-redis/v9 Cmdable
	@go mod tidy