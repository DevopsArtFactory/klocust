export CGO_ENABLED=0

define echo_target
	@echo ">>> $@"
endef

fmt:
	@$(echo_target)
	@go fmt ./...

build: fmt
	@$(echo_target)
	@go build -o build/klocust ./cmd/main.go
	@ls -lh build/klocust

clean: fmt
	@$(echo_target)
	@rm -rf build/klocust

run: clean build
	@$(echo_target)
	@build/klocust

.PHONY: fmt build clean run
