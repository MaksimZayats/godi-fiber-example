# godi-fiber-example

## A real example of using [GoDI](https://github.com/MaximZayats/godi)

## Snippet
```go
type H = func(*fiber.Ctx) error

// `stringFromDI` will be injected into the handler
func handler(c *fiber.Ctx, stringFromDI string) error {
    return c.SendString("Hello from di: " + stringFromDI)
}

func main() {
    di.AddInstance[string]("I'm string from DI!!!", c)
    ...
    app.Get("/", injection.Inject[H](handler))
}
```
