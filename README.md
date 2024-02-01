# be-service-sample

#### Table of contents:

- [be-service-sample](#be-service-sample) - [Table of contents:](#table-of-contents)
  - [- Helpful links for developers](#--helpful-links-for-developers)
    - [Why lynqtech wants to have a common service design?](#why-lynqtech-wants-to-have-a-common-service-design)
  - [Swagger](#swagger)
  - [Static Code Analysis](#static-code-analysis)
  - [How to use a common lynqtech golangci-lint configuration](#how-to-use-a-common-lynqtech-golangci-lint-configuration)
  - [Run with docker-compose](#run-with-docker-compose)
  - [Formatting](#formatting)
  - [Detailed information](#detailed-information)
  - [Helpful links for developers](#helpful-links-for-developers)

---

This is a sample service. This services demonstrates how lynqtech would like to see new services structured.

#### Why lynqtech wants to have a common service design?

- A common service design helps people from other squads to quickly orient themselves in the service.
- This service design has a clear separation between 3 layers (presentation, domain and data).

#### Does it determine which packages have to used?

It does not specify which packages to use. Example: HTTP handlers must reside in `pkg/server/handler`, but whether you use raw [net/http](https://pkg.go.dev/net/http), [Echo](https://echo.labstack.com/) or [Gin](https://gin-gonic.com/) to implement them is up to you. Exceptions, i.e. packages that **must** be used, are the ones marked `standard` in [backend go libs](https://lynqtech.atlassian.net/wiki/x/6wc2G). Note: The backend sample service repository is marked as standard, but this relates to the structure only.

## Swagger

- to install Swagger Documentation 2.0 generator:

  ```
  # 1.16 or newer
  $ go install github.com/swaggo/swag/cmd/swag@latest

  # 1.15 or older go version:
  go get -u github.com/swaggo/swag/cmd/swag
  ```

- to generate docs, from the project's root folder run:
  ```
  swag init -g pkg/server/echo.go -output api/openapi/docs
  ```

## Static Code Analysis

This project will be scanned for potential code issues by golangci-lint during CI. The settings will be read from [lqt-go-linter/.golangci.yml](lqt-go-linter/.golangci.yml).

At the moment, the build will continue even if the linter finds new issues. In the future, the build will break on new findings.

To run it locally, call `make lint`.

The linting step can also be integrated into your IDE. Please check the docs for your IDE.

Please review new findings. If you are pretty sure your code is correct, you may accept it by adding `/nolint:<rule>` to the line affected (e.g. `/nolint:deadcode`). Please add a meaningful commit message. TIA.

## How to use a common lynqtech golangci-lint configuration
[How to use lgt-go-linter in other repositories](https://lynqtech.atlassian.net/l/cp/0n1T1uQu)

## Run with docker-compose

This service provides a docker-compose file to simplify local development. The docker-compose command starts a postrgesql, a db migration and service container.

```
docker-compose up
```

## Formatting

Non-Go code is formatted using [Prettier](https://prettier.io/).

To check files whether they are formatted correctly:

```
prettier --check .
```

Actually prettify the files:

```
prettier --write .
```

If you don't want files to be checked or written to by Prettier, put them into `.prettierignore`.

## Detailed information

- [Domain Models](./pkg/model/domain/domain.md)
- [DTO Models](./pkg/model/dto/dto.md)
- [HTTP Handler - Presentation Layer](./pkg/server/handler/handler.md)
- [Use Cases - Domain Layer](./pkg/usecase/usecase.md)
- [Repository - Data Source Layer](./pkg/repository/repository.md)

## Helpful links for developers

- [Conventions](https://lynqtech.atlassian.net/wiki/x/tCC9Fw)
- [Logging Conventions](https://lynqtech.atlassian.net/wiki/x/XQEqHQ)
- [Error Handling](https://lynqtech.atlassian.net/wiki/x/bgCHI)
- [Eventing](https://lynqtech.atlassian.net/wiki/x/vAj3G)
- [Static Code Analysis](https://lynqtech.atlassian.net/wiki/x/GQDYIQ)

ToDo:
A list of ToDos
- Validate the range of all DTO inputs
- Make "NotAllowed" to an expected (non 500) error