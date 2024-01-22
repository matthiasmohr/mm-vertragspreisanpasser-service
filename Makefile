GOLANGCI_LINT_VERSION=v1.42
GOLANGCI_LINT_COMMAND=golangci-lint run --allow-parallel-runners --timeout 30m -v --config lqt-go-linter/.golangci.yml
PRETTIER_NODE_DOCKER_IMAGE_TAG=16.13.1-alpine3.15
PRETTIER_VERSION=2.5.1

# Run unit tests.
test:
	go test -coverpkg ./... -v ./...

# Convert coverage report to HTML.
dist/coverage.html: dist/coverage.out
	go tool cover -html=$< -o $@

# Generate coverage report.
dist/coverage.out: dist
	go test -coverprofile $@ -coverpkg ./... ./...

# Create ./dist directory.
dist:
	mkdir -p dist

# Lint Go files, see https://golangci-lint.run/.
lint:
	# Use Docker to pin the linter to the same version as the one run via Drone.
	docker run --rm -it \
		-w /app \
		-v ${PWD}:/app:ro \
		-e SSH_AUTH_SOCK=/ssh-agent \
		-v ${SSH_AUTH_SOCK}:/ssh-agent \
		-e GOPROXY=${GOPROXY} \
		golangci/golangci-lint:${GOLANGCI_LINT_VERSION} \
		/bin/sh -c \
			"mkdir ~/.ssh \
			&& ssh-keyscan -T 60 github.com >> ~/.ssh/known_hosts \
			&& git config --global url.'git@github.com:enercity'.insteadOf "https://github.com/enercity" \
			&& ${GOLANGCI_LINT_COMMAND}"

# Lint Go files. This make target is only used in the Drone pipeline.
lint-drone:
	${GOLANGCI_LINT_COMMAND}

# Check whether files are formatted correctly with Prettier (https://prettier.io/).
prettier-check:
	# Use Docker to pin Prettier to the same version as the one run via Drone.
	docker run --rm -it \
		-w /app \
		-v ${PWD}:/app:ro \
		node:${PRETTIER_NODE_DOCKER_IMAGE_TAG} \
		/bin/sh -c \
			"npm install --global --save-dev --save-exact prettier@${PRETTIER_VERSION} \
			&& prettier -c ."

# Format files with Prettier (https://prettier.io/).
prettier-write:
	# Use Docker to pin Prettier to the same version as the one run via Drone.
	docker run --rm -it \
		-w /app \
		-v ${PWD}:/app \
		node:${PRETTIER_NODE_DOCKER_IMAGE_TAG} \
		/bin/sh -c \
			"npm install --global --save-dev --save-exact prettier@${PRETTIER_VERSION} \
			&& prettier -w ."

# Clean files generated from Makefile.
clean:
	rm -rf dist

# These are all paths that should be built every time they are called as targets directly or indirectly.
.PHONY: \
	clean \
	dist/coverage.html \
	dist/coverage.out \
	lint \
	lint-drone \
	prettier-check \
	prettier-write \
	test \
