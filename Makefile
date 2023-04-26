build-contracts:
	"$(CURDIR)/scripts/build-contracts.sh"

build-mocks:
	make clean
	go generate ./...

clean:
	-rm */**/*_mock.go

test:
	make build-mocks
	go test -coverprofile=cover.out ./...

test-tool:
	make test
	go tool cover -html=cover.out -o=cover.html

run-consensus:
	go run cmd/consensus/main.go

install:
	bash scripts/install.sh
