package application

import (
	distorage "github.com/MaximZayats/godi-fiber-example/pkg/di/generated"
	"github.com/MaximZayats/godi/di"
	"github.com/MaximZayats/godi/injection"
	"github.com/gofiber/fiber/v2"
)

type H = func(*fiber.Ctx) error

func setupDependencies(container ...*di.Container) {
	c := di.GetContainer(container...)
	di.AddInstance[string]("I'm string from DI!!!", c)
}

func handler(c *fiber.Ctx, stringFromDI string) error {
	return c.SendString("Hello from di: " + stringFromDI)
}

func GetApp() *fiber.App {
	injection.Configure(distorage.Config)
	c := di.NewContainer()
	setupDependencies(c)

	app := fiber.New()

	app.Get("/", injection.Inject[H](handler, c))

	ok := injection.VerifyInjections()
	if !ok {
		panic("Restart required")
	}

	return app

}
