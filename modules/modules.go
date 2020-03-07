package modules

// Module is a structure for dynamic intents with a Tag, some Patterns and Responses and
// a Replacer function to execute the dynamic changes.
type Module struct {
	Tag       string
	Patterns  []string
	Responses []string
	Replacer  func(string, string, string) (string, string)
}

var modules []Module

// RegisterModule registers a module in the list
func RegisterModule(module Module) {
	modules = append(modules, module)
}

// GetModules returns all the registered modules
func GetModules() []Module {
	return modules
}

// ReplaceContent apply the Replacer of the matching module to the response and returns it
func ReplaceContent(tag, entry, response, token string) (string, string) {
	for _, module := range modules {
		if module.Tag != tag {
			continue
		}

		return module.Replacer(entry, response, token)
	}

	return tag, response
}
