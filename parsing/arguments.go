package parsing

import (
	"strings"
)

type ArgumentConfig struct {
	Name string
	DefaultValue string
	CurrentValue string
	AllowedValues []string
}

type Arguments []*ArgumentConfig

// ParseFullArgument parses and validates a command-line argument with the specified format.
// Parameters:
//   - arg: the full command-line argument to parse
//   - prefix: expected argument prefix (e.g., "--")
//   - sign: separator between argument and value (e.g., "=")
//   - size: length of the argument name
// Returns the argument name and its value if valid, empty strings otherwise.
func ParseFullArgument(arg, prefix, sign string, size int) (string, string) {
	expectedSize := len(prefix) + size
	if len(arg) < expectedSize { return "", "" }
	if !strings.HasPrefix(arg, prefix) { return "", "" }
	signIndex := strings.Index(arg, sign)
	if signIndex != expectedSize { return "", "" }
	return arg[:signIndex], arg[signIndex + 1:]
}

func ValidateArgumentValue() bool {
	return true
}

// ProcessArguments parses a slice of command-line arguments and validates them 
// against expected options, assigning values to matching configurations.
// Parameters:
//   - args: slice of arguments to process
//   - configs: slice of argument configurations
//   - prefix: expected argument prefix (e.g., "--") 
//   - sign: separator between argument and value (e.g., "=")
//   - size: length of the argument name
func ProcessArguments(args []string, configs Arguments, prefix, sign string, size int) {
	if len(args) < 2 { return }
	for _, fullArg := range args {
		arg, val := ParseFullArgument(fullArg, prefix, sign, size)
		// Validate argument
		var config *ArgumentConfig
		for _, argConfig := range configs {
			if argConfig.Name == arg {
				config = argConfig
				break
			}
		}
		if config == nil { continue }
		if val == "" { 
			config.CurrentValue = config.DefaultValue
		} else {
			// TODO: Validate value using ValidateArgumentValue function.
			config.CurrentValue = val
		}
	}
}
