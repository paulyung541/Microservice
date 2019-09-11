GO    					:= go

PROJECT_NAME            ?= backend

SHELL_CMD               ?= zsh
TESTARGS                ?= -v -race
COVERARGS               ?= -coverprofile=profile.out -covermode=atomic
TEST                    ?= $($(SHELL_CMD) go list ./... | grep -v '/vendor/')
GOFMT_FILES             ?= $($(SHELL_CMD) find . -name '*.go' | grep -v vendor | xargs)
GOCOV                   := $(GOPATH)/bin/gocov

RUN_DIR                 ?= $(GOPATH)/src/cmd/$(PROJECT_NAME)/main.go

_comma := ,
_space :=
_space +=


.PHONY: all
all: run

.PHONY: run
run:
	@echo ">>>running $(PROJECT_NAME)"
	@$(GO) run cmd/$(PROJECT_NAME)/main.go

.PHONY: setup
setup:
	@echo ">> installing dependencies"
	@$(GO) get -u "github.com/alecthomas/gometalinter"
	@$(GO) get -u "github.com/golang/tools/cmd/goimports"
	@gometalinter --install --update

.PHONY: test
test:
	@echo ">> running tests"
	@$(GO) test $(TEST) $(TESTARGS)

.PHONY: fmt
fmt:
	@gofmt -w cmd/backend/main.go
	@goimports -w cmd/backend/main.go

	@gofmt -w api/auth.go
	@goimports -w api/auth.go

	@gofmt -w api/init.go
	@goimports -w api/init.go

.PHONY: lint
lint:
	@echo ">> linting code"
	@golint ./...


.PHONY: fmtcheck
fmtcheck:
	@echo ">> checking code style"
	@gometalinter --vendor --disable-all \
		--enable=gofmt \
		--enable=goimports \
		./...

.PHONY: build
build:
	@echo ">> building binaries"
	$(GO) build --prefix $(PREFIX)

.PHONY: docker
docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .