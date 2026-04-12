.PHONY: dev build test test-go test-frontend test-e2e clean

dev:
	wails dev

build:
	wails build -clean -platform linux/amd64

build-all:
	wails build -clean -platform linux/amd64 windows/amd64,darwin/amd64,darwin/arm64

test: test-go test-frontend

test-go:
	go test -v -cover ./backend/...

test-frontend:
	cd frontend && npm run test

test-e2e:
	cd e2e && npx playwright test

clean:
	rm -rf build/bin
