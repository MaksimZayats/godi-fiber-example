package application

import (
	"github.com/MaximZayats/godi-fiber-example/pkg/storage/decorators"
	"github.com/MaximZayats/godi/di"
	"github.com/MaximZayats/godi/injection"
	"github.com/gofiber/fiber/v2"
)

type H = func(*fiber.Ctx) error

func setupDependencies(container ...*di.Container) {
	c := di.GetContainer(container...)
	di.AddInstance[string]("I'm string from DI!!!", c)
	di.AddInstance[int](123123, c)
}

func handler(c *fiber.Ctx, stringFromDI string) error {
	return c.SendString("Hello from di: " + stringFromDI)
}

func GetApp() *fiber.App {
	c := di.NewContainer()
	setupDependencies(c)

	injection.Configure(decorators.Config)

	app := fiber.New()

	app.Get("/", injection.Inject[H](handler, c))

	injection.MustVerifyInjections()

	return app

}
