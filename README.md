
# Fux

An interface for works with [Gorilla mux](https://github.com/gorilla/mux)

## Installation

Install with

```bash
    go get -u github.com/kamandlou/fux
```

## Examples

```go
func main() {
    f := fux.New()
    f.Get("/", HomeHandler)
    f.Post("/products", ProductsHandler)
    f.Get("/articles", ArticlesHandler)
    f.Run(":8080")
}
```

Paths can have variables. They are defined using the format {name} or {name:pattern}. If a regular expression pattern is not defined, the matched variable will be anything until the next slash. For example:

```go
f := fux.New()
f.Get("/products/{key}", ProductHandler)
f.Get("/articles/{category}/", ArticlesCategoryHandler)
f.Get("/articles/{category}/{id:[0-9]+}", ArticleHandler)
f.Run(":8080")
```
