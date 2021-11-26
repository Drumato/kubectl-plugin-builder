build: format test
	go build -o bin/kubectl-plugin-builder ./cmd/kubectl-plugin-builder

format:
	go fmt ./cmd/... \
		./internal/cmd/... ./internal/cli/... ./internal/kpbtemplate \
			./internal/kpbtemplate/cli/... ./internal/kpbtemplate/command/... \
			./internal/kpbtemplate/gitignore/... ./internal/kpbtemplate/license/... \
			./internal/kpbtemplate/makefile/... ./internal/kpbtemplate/module/... \
			./internal/kpbtemplate/node/... ./internal/kpbtemplate/module/... 

test:
	go test ./cmd/... \
		./internal/cmd/... ./internal/cli/... ./internal/kpbtemplate \
			./internal/kpbtemplate/cli/... ./internal/kpbtemplate/command/... \
			./internal/kpbtemplate/gitignore/... ./internal/kpbtemplate/license/... \
			./internal/kpbtemplate/makefile/... ./internal/kpbtemplate/module/... \
			./internal/kpbtemplate/node/... ./internal/kpbtemplate/module/... 
