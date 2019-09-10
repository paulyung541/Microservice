GO    					:= go

PROJECT_NAME            ?= auth

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
all: format test

.PHONY: run
run:
	@echo ">>>running $(PROJECT_NAME)"
	@$(GO) run cmd/$(PROJECT_NAME)/main.go

.PHONY: setup
setup:
	@echo ">> installing dependencies"
	@$(GO) get -u "github.com/alecthomas/gometalinter"
	@gometalinter --install --update

.PHONY: test
test:
	@echo ">> running tests"
	@$(GO) test $(TEST) $(TESTARGS)

.PHONY: lint
lint:
	@echo ">> linting code"
	@gometalinter --vendor --disable-all \
		--enable=varcheck \
		--enable=misspell \
		--enable=vet \
		--enable=vetshadow \
		--enable=golint \
		--deadline=10m \
		./...


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