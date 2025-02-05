# PARI-Test

> Pari Backend Test.

Implementation. The project using [typical-go](https://github.com/typical-go/typical-go) as its build-tool.

- Application
  - [x] [Go-Standards](https://github.com/golang-standards/project-layout) Project Layout
  - [x] Environment Variable Configuration
  - [x] Health-Check
  - [ ] Debug API
  - [x] Graceful Shutdown
- Layered architecture
  - [x] [SOLID Principle](https://en.wikipedia.org/wiki/SOLID)
  - [x] Dependency Injection (using `@ctor` annotation)
  - [x] Database Transaction
- HTTP Server
  - [x] [Echo framework](https://echo.labstack.com/)
  - [ ] Server Side Caching
    - [ ] Cache but revalidate (Header `Cache-Control: no-cache`)
    - [ ] Set Expiration Time (Header `Cache-Control: max-age=120`)
- RESTful
  - [x] Create Resource (`POST` verb)
  - [x] Update Resource (`PUT` verb)
  - [ ] Partially Update Resource (`PATCH` verb)
  - [x] Find Resource (`GET` verb)
    - [x] Offset Pagination (Query param `?limit=100&offset=0`)
    - [ ] Sorting (Query param `?sort=-title,created_at`)
  - [x] Delete resource (`DELETE` verb)
- Testing
  - [ ] Table Driven Test
- Others
  - [ ] Database migration and seed tool
  - [ ] Generate code for repository layer
  - [ ] Releaser
  
  ## Project Layout

Typical-Rest encourage [standard go project layout](https://github.com/golang-standards/project-layout)

Source codes:
- [`internal`](internal): private codes for the project
  - [`internal/`](internal/)
    - [`internal/infra`](internal/app/infra): infrastructure for the project e.g. config and connection object
    - [`internal/controllers`](internal/app/controller): presentation layer
    - [`internal/service`](internal/app/service): logic layer
    - [`internal/repo`](internal/app/repo): data-access layer for database repo or domain model
  - [`internal/generated`](internal/generated): code generated e.g. mock, etc.
- [`pkg`](pkg): shareable codes e.g. helper/utility Library
- [`cmd`](cmd): the main package

Others directory:
- [`static`](static) for static file, like template .html

## Dependency Injection

Typical-Rest encourage [dependency injection](https://en.wikipedia.org/wiki/Dependency_injection) using [uber-dig](https://github.com/uber-go/dig) and annotations (`@ctor`).

## Application Config

Typical-Rest encourage [application config with environment variables](https://12factor.net/config) using [envconfig](https://github.com/kelseyhightower/envconfig) and annotation (`@envconfig`).

```go
type (
  // AppCfg application configuration
  // @envconfig (prefix:"APP")
  AppCfg struct {
    Address string `envconfig:"ADDRESS" default:":8089" required:"true"`
    Debug   bool   `envconfig:"DEBUG" default:"true"`
  }
)
```

Add import side-effect to make it work
```go
import(
  _ "github.com/typical-go/typical-rest-server/internal/generated/envcfg"
)
```

## How to run

1. Create Database with name pari and schema pari
2. Create table with ddl
3. Create .env file (example can be seen)
4. run go (go run cmd/main.go)

## References

Golang:
- [Go Documentation](https://golang.org/doc/)
- [Go For Industrial Programming](https://peter.bourgon.org/go-for-industrial-programming/)
- [Uber Go Style Guide](https://github.com/uber-go/guide)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

RESTful API:
- [Best Practices for Designing a Pragmatic RESTful API](https://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api)
- [Everything You Need to know About API Pagination](https://nordicapis.com/everything-you-need-to-know-about-api-pagination/)