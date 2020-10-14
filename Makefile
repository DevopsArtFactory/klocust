export CGO_ENABLED=0

define echo_target
	@echo ">>> $@"
endef

ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

fmt:
	@$(echo_target)
	@go fmt ./...

build: fmt
	@$(echo_target)
	@go build -o build/klocust ./cmd/main.go
	@ls -lh build/klocust

clean:
	@$(echo_target)
	@rm -rf build/klocust

run:
	@$(echo_target)
	@go run ./cmd/main.go $(RUN_ARGS)

.PHONY: fmt build clean run
