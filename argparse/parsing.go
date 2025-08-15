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

type SingleValueArg struct {
	Name string
	DefaultValue string
	CurrentValue string
	AllowedValues []string
}
type MultiValueArg struct {
	Name string
	Separator string
	DefaultValues []string
	CurrentValues []string
	AllowedValues []string
}

type SingleArgs []*SingleValueArg
type MultiArgs []*MultiValueArg

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

// ExtractMultipleValues extracts single values from a string containing them separated by a sign.
// It will work with strings that do not cointain multiple seperated values. However, beware that such string could represent a single value or multiple but not seperated.
// Parameters:
//   - fullString: string containing the values to be sepatared
//   - separator: expected sign to be separating the values
// Returns a slice containing all the extracted values.
func ExtractMultipleValues(fullString, separator string) []string {
	length := len(fullString)
	if length < 1 { return nil }
	if separator == "" { separator = "," }
	if !strings.HasSuffix(fullString, separator) {
		fullString = fullString + separator
	}
	firstIndex, lastIndex := 0, strings.Index(fullString, separator)
	var result []string
	for {
		value := fullString[firstIndex:lastIndex]
		result = append(result, value)
		firstIndex = lastIndex + 1
		if firstIndex >= length { break }
		relativeIndex := strings.Index(fullString[firstIndex:], separator)
		if relativeIndex < 0 { break }
		lastIndex = firstIndex + relativeIndex
	}
	return result
}

// ValidateArgumentValue validates if the received value matches the allowed values
// in the provided argument configuration.
// If the value is invalid or not allowed, the function returns the default value automatically.
// Parameters:
//   - value: value to be validated
//   - argConfig: argument configuration containing allowed values and defaults
// Returns the validated value if valid, or the default value otherwise.
func ValidateArgumentValue(value string, argConfig SingleValueArg) string {
	if valid := slices.Contains(argConfig.AllowedValues, ANY); valid { return value }
	if valid := slices.Contains(argConfig.AllowedValues, value); valid { return value }
	return argConfig.DefaultValue
}

// ValidateArgumentValues validates if all the received values match the allowed values
// in the provided argument configuration.
// If a value is invalid or not allowed, it will not be considered in the final result,
// repeated values are considered once.
// If the final result is empty, the function returns the default values.
// Parameters:
//   - values: slice of values to be validated
//   - argConfig: argument configuration containing allowed values and defaults
// Returns the validated values if not empty, default values otherwise.
func ValidateArgumentValues(values []string, argConfig MultiValueArg) []string {
	if slices.Contains(argConfig.AllowedValues, ANY) { return values }
	valuesKeys := make(map[string]string)
	for _, value := range values {
		if !slices.Contains(argConfig.AllowedValues, value) { continue }
		valuesKeys[value] = ""
	}
	if len(valuesKeys) == 0 { return argConfig.DefaultValues }
	var result []string
	for key := range valuesKeys {
		result = append(result, key)
	}
	return result
}

// ProcessSingleValueArgs parses a slice of command-line arguments and validates them 
// against expected options, assigning values to matching configurations.
// All configurations are initialized with their default values. Invalid or unknown 
// arguments are silently ignored, ensuring all configurations have valid values.
// Parameters:
//   - args: slice of arguments to process
//   - configs: slice of argument configurations
//   - prefix: expected argument prefix (e.g., "--") 
//   - sign: separator between argument and value (e.g., "=")
//   - size: length of the argument name
func ProcessSingleValueArgs(args []string, configs SingleArgs, prefix, sign string, size int) {
	for _, config := range configs { config.CurrentValue = config.DefaultValue }
	if len(args) < 1 { return }
	for _, fullArg := range args {
		arg, val, errArg := ValidateAndExtractArgument(fullArg, prefix, sign, size)
		if errArg != nil { continue }
		var config *SingleValueArg
		for _, argConfig := range configs {
			if argConfig.Name == arg { config = argConfig; break }
		}
		if config == nil { continue }
		config.CurrentValue = ValidateArgumentValue(val, *config)
	}
}

// ProcessMultiValueArgs parses a slice of command-line arguments and validates them 
// against expected options, assigning values to matching configurations.
// All configurations are initialized with their default values. Invalid or unknown 
// arguments are silently ignored, ensuring all configurations have valid values.
// Parameters:
//   - args: slice of arguments to process
//   - configs: slice of argument configurations
//   - prefix: expected argument prefix (e.g., "--") 
//   - sign: separator between argument and value (e.g., "=")
//   - size: length of the argument name
func ProcessMultiValueArgs(args []string, configs MultiArgs, prefix, sign string, size int) {
	for _, config := range configs { config.CurrentValues = config.DefaultValues }
	if len(args) < 1 { return }
	for _, fullArg := range args {
		arg, val, errArg := ValidateAndExtractArgument(fullArg, prefix, sign, size)
		if errArg != nil { continue }
		var config *MultiValueArg
		for _, argConfig := range configs {
			if argConfig.Name == arg { config = argConfig; break }
		}
		if config == nil { continue }
		argValues := ExtractMultipleValues(val, config.Separator)
		config.CurrentValues = ValidateArgumentValues(argValues, *config)
	}
}
