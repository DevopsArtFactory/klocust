CGO_ENABLED?=0

# echo current target name
define echo_target
	@echo ">>> $@"
endef

# for make run with arguments
ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

.PHONY: fmt
fmt:
	@$(echo_target)
	@go fmt ./...

.PHONY: test
test:
	@$(echo_target)
	@go test ./...

.PHONY: build
build: fmt
	@$(echo_target)
	@go build -ldflags "-w -s" -o build/klocust ./cmd/main.go

.PHONY: clean
clean:
	@$(echo_target)
	@rm -rf build/klocust

.PHONY: run
run:
	@$(echo_target)
	@go run ./cmd/main.go $(RUN_ARGS)
