package argparse

import (
	"strings"
	"slices"
)

const ANY = "any"

type ArgumentConfig struct {
	Name string
	DefaultValue string
	CurrentValue string
	AllowedValues []string
}

type Arguments []*ArgumentConfig

// ValidateAndExtractArgument parses and validates a command-line argument with
// the specified format.
// Parameters:
//   - arg: the full command-line argument to parse
//   - prefix: expected argument prefix (e.g., "--")
//   - sign: separator between argument and value (e.g., "=")
//   - size: length of the argument name
// Returns the argument and its value if valid, empty strings otherwise.
func ValidateAndExtractArgument(arg, prefix, sign string, size int) (string, string) {
	expectedSize := len(prefix) + size
	if len(arg) < expectedSize { return "", "" }
	if !strings.HasPrefix(arg, prefix) { return "", "" }
	signIndex := strings.Index(arg, sign)
	if signIndex != expectedSize { return "", "" }
	return arg[:signIndex], arg[signIndex + 1:]
}

// ValidateArgumentValue validates if the received value matches the allowed values
// in the provided argument configuration.
// Parameters:
//   - value: value to be validated
//   - argConfig: argument configuration containing allowed values and defaults
// Returns the validated value and true if valid, default value and false otherwise.
func ValidateArgumentValue(value string, argConfig ArgumentConfig) (string, bool) {
	if valid := slices.Contains(argConfig.AllowedValues, ANY); valid {
		return value, valid
	}
	if valid := slices.Contains(argConfig.AllowedValues, value); valid {
		return value, valid
	}
	return argConfig.DefaultValue, false
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
	for _, config := range configs {
		config.CurrentValue = config.DefaultValue
	}
	if len(args) < 1 { return }
	for _, fullArg := range args {
		arg, val := ValidateAndExtractArgument(fullArg, prefix, sign, size)
		var config *ArgumentConfig
		for _, argConfig := range configs {
			if argConfig.Name == arg {
				config = argConfig
				break
			}
		}
		if config == nil { continue }
		val, ok := ValidateArgumentValue(val, *config)
		if !ok { continue }
		config.CurrentValue = val
	}
}
