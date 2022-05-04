// Code generated by DI. DO NOT EDIT.

package decorators

import (
	"reflect"

	"github.com/MaximZayats/godi/codegen"
	"github.com/MaximZayats/godi/di"
	"github.com/gofiber/fiber/v2"
)

var Config = codegen.Config{
	PackageName:         "decorators",
	PathToStorageFolder: `./pkg/storage/decorators`,
	StorageFileName:     "di_injection_decorators_gen.go",
	GetterFunction:      getDecorator,
}

var functions = map[string]any{

	"func(*fiber.Ctx, string) error": func(f func(*fiber.Ctx, string) error, c *di.Container) any {
		return func(a *fiber.Ctx) error {
			return f(
				a,
				di.MustGetFromContainer[string](c),
			)
		}
	},
}

func getDecorator(f any) (any, bool) {
	v, ok := functions[reflect.TypeOf(f).String()]
	return v, ok
}
