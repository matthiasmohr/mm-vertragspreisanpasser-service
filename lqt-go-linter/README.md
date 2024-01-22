# lqt-go-linter

Contains a common `golangci.yml` to be used in our backend projects.

# How to use lgt-go-linter in other repositories


In short:

```
git subtree add --prefix lqt-go-linter git@github.com:enercity/lqt-go-linter.git tags/v1.3.0 --squash
```

Then point all executions of `golangci-lint` to the newly created folder:

```
golangci-lint --config=lqt-go-linter/.golangci.yml run
```

For details see [How to use lqt-go-linter in go services](https://lynqtech.atlassian.net/l/cp/0n1T1uQu).

## pin golangci-lint version with docker

Use Docker to pin the linter to the same version.

### Linux

This command mounts your ssh-agent to allow access to private git repos.

```
	docker run --rm -it \
		-w /app \
		-v ${PWD}:/app:ro \
		-e SSH_AUTH_SOCK=/ssh-agent \
		-v ${SSH_AUTH_SOCK}:/ssh-agent \
		golangci/golangci-lint:v1.46.2 \
  		golangci-lint run --allow-parallel-runners --timeout 30m -v --config=lqt-go-linter/.golangci.yml
```


### Mac
Sadly mounting ssh-agent is not possible on Mac OS. So instead we need to mount the go mod cache as a volume.

This requires you to run
`go mod download` before you can run the linting.

```
	docker run --rm -it \
		-w /app \
		-v ${PWD}:/app:ro \
		-v ~/go/pkg/mod/cache:/go/pkg/mod/cache:ro \
		golangci/golangci-lint:v1.46.2 \
  		golangci-lint run --allow-parallel-runners --timeout 30m -v --config=lqt-go-linter/.golangci.yml
```
