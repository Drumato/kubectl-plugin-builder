build: format test
	go build -o bin/kubectl-plugin-builder ./cmd/kubectl-plugin-builder

format:
	go fmt ./cmd/... ./internal/cmd/... ./internal/kpbtemplate/cli/... \
		./internal/kpbtemplate/gitignore/... ./internal/kpbtemplate/license/... \
		./internal/kpbtemplate/makefile/... ./internal/kpbtemplate/module/... \
		./internal/kpbtemplate

test:
	go test ./cmd/... ./internal/cmd/... ./internal/kpbtemplate/cli/... \
		./internal/kpbtemplate/gitignore/... ./internal/kpbtemplate/license/... \
		./internal/kpbtemplate/makefile/... ./internal/kpbtemplate/module/... \
		./internal/kpbtemplate
