package argparse

import (
	"fmt"
	"slices"
	"strings"
)

type InvalidFormatError struct {
	Argument string
	Reason string
}
type UnknownArgumentError struct {
	Argument string
}

func (e InvalidFormatError) Error() string {
	return fmt.Sprintf("invalid argument format	%s: %s", e.Argument, e.Reason)
}
func (e UnknownArgumentError) Error() string {
	return fmt.Sprintf("unknown argument %s", e.Argument)
}

const ANY = "any"

type ArgumentConfig struct {
	Name string
	DefaultValue string
	CurrentValue string
	AllowedValues []string
}

type Arguments []*ArgumentConfig

// ValidateAndExtractArgument parses and validates a command-line argument with the specified format.
// The function checks argument length, prefix, and separator position to ensure proper structure.
// Parameters:
//   - arg: the full command-line argument to parse
//   - prefix: expected argument prefix (e.g., "--")
//   - sign: separator between argument and value (e.g., "=")
//   - size: length of the argument name
// Returns the argument name, its value, and nil if valid; empty strings and error otherwise.
func ValidateAndExtractArgument(arg, prefix, sign string, size int) (string, string, error) {
	expectedSize := len(prefix) + size
	if len(arg) < expectedSize {
		return "", "", InvalidFormatError{ Argument: arg, Reason: "insufficient argument length" }
	}
	if !strings.HasPrefix(arg, prefix) {
		return "", "", InvalidFormatError{ Argument: arg, Reason: "missing required prefix" }
	}
	signIndex := strings.Index(arg, sign)
	if signIndex != expectedSize {
		return "", "", InvalidFormatError{ Argument: arg, Reason: "separator in incorrect position" }
	}
	return arg[:signIndex], arg[signIndex + 1:], nil
}

// ValidateArgumentValue validates if the received value matches the allowed values
// in the provided argument configuration.
// If the value is invalid or not allowed, the function returns the default value automatically.
// Parameters:
//   - value: value to be validated
//   - argConfig: argument configuration containing allowed values and defaults
// Returns the validated value if valid, or the default value otherwise.
func ValidateArgumentValue(value string, argConfig ArgumentConfig) string {
	if valid := slices.Contains(argConfig.AllowedValues, ANY); valid { return value }
	if valid := slices.Contains(argConfig.AllowedValues, value); valid { return value }
	return argConfig.DefaultValue
}

// ProcessArguments parses a slice of command-line arguments and validates them 
// against expected options, assigning values to matching configurations.
// All configurations are initialized with their default values. Invalid or unknown 
// arguments are silently ignored, ensuring all configurations have valid values.
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
		arg, val, errArg := ValidateAndExtractArgument(fullArg, prefix, sign, size)
		if errArg != nil { continue }
		var config *ArgumentConfig
		for _, argConfig := range configs {
			if argConfig.Name == arg {
				config = argConfig
				break
			}
		}
		if config == nil { continue }
		config.CurrentValue = ValidateArgumentValue(val, *config)
	}
}
