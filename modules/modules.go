package modules

type Module struct {
	Tag       string
	Patterns  []string
	Responses []string
	Replacer  func(string, string) string
}

var modules []Module

func RegisterModule(module Module) {
	modules = append(modules, module)
}

func GetModulesIntents() []Module {
	return modules
}
