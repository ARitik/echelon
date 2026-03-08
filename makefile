.PHONY: dev dev-client dev-server test test-client test-server

dev:
	@echo "Starting both servers..."
	$(MAKE) -j2 dev-client dev-server

dev-client:
	cd client && npm run dev

dev-server:
	cd server && go run ./cmd/server

test:
	$(MAKE) test-client
	$(MAKE) test-server

test-client:
	cd client && npm run test:run

test-server:
	cd server && go test ./...