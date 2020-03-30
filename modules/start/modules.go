package start

import (
	"fmt"

	"github.com/gookit/color"
)

// A start Module is a module that will be executed when a connection is opened by a user
type Module struct {
	Action func(string)
}

var modules []Module

// RegisterModule registers the given module in the array
func RegisterModule(module Module) {
	modules = append(modules, module)
}

// ExecuteModules will execute all the registered start modules with the user token
func ExecuteModules(token string) {
	fmt.Println(color.FgGreen.Render("Executing start modules.."))

	for _, module := range modules {
		module.Action(token)
	}
}
