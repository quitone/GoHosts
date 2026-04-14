.PHONY: dev build test test-go test-frontend test-e2e clean fmt lint lint-go lint-frontend check

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

# 格式化所有代码
fmt:
	@echo "==> Formatting Go code..."
	goimports -w -local gohosts ./backend ./.
	@echo "==> Formatting Frontend code..."
	cd frontend && npm run format

# 运行所有 lint 检查
lint: lint-go lint-frontend

lint-go:
	@echo "==> Linting Go code..."
	golangci-lint run ./...

lint-frontend:
	@echo "==> Linting Frontend code..."
	cd frontend && npm run lint

# 运行类型检查（前端）
type-check:
	cd frontend && npm run type-check

# 全面检查（格式化 + lint + 测试）
check: fmt lint test
	@echo "All checks passed!"