
# Fux

An interface for works with [gorilla mux](https://github.com/gorilla/mux)

## Installation

Install with

```bash
    go get -u github.com/kamandlou/fux
```

## Usage/Examples

```go
func main() {
    f := fux.New()
    f.Get("/", HomeHandler)
    f.Post("/products", ProductsHandler)
    f.Get("/articles", ArticlesHandler)
    f.Run(":8080")
}
```

