# argparse

Go package for extracting and validating command-line arguments with flexible configurations. Built around two argument types for both single and multiple value handling.

## Features
* **Structured argument validation** with configurable format: prefix, separator, and fixed name size
* **Dual argument support** for single values and multiple values with customizable separators
* **Flexible value constraints** using a slice of allowed string values or `ANY` constant for unrestricted input
* **In-place modifications** to argument configurations for reusability across operations
* **Automatic fallback** to default values for missing or invalid argument values
* **Custom error types** for detailed validation and parsing error handling
* **Batch processing** of multiple arguments configurations in a single function call

## Quick start
### Single value arguments
```Go
import "github.com/leojimenezg/psetup/argparse"

// Configure expected argument
routeArg := argparse.SingleValueArg{
    Name: "-rte",
    DefaultValue: "./",
    AllowedValues: []string{ argparse.ANY },
}

// Process command-line arguments
args := os.Args[1:]
configs := argparse.SingleArgs{ &routeArg }
argparse.ProcessSingleValueArgs(args, configs, "-", "=", 3)

// Access result
fmt.Println(routeArg.CurrentValue)
```

### Multiple value arguments
```Go
import "github.com/leojimenezg/psetup/argparse"

// Configure expected argument
documentArg := argparse.MultiValueArg{
    Name: "-dcs",
    Separator: ",",
    DefaultValues: []string{ "all" },
    AllowedValues: []string{ "all", "readme", "license", "ignore" },
}

// Process command-line arguments
args := os.Args[1:]
configs := argparse.MultiArgs{ &documentArg }
argparse.ProcessMultiValueArgs(args, configs, "-", "=", 3)

// Access result
fmt.Println(documentArg.CurrentValues)
```

## Core types
### SingleValueArg
```Go
type SingleValueArg struct {
    Name          string   // Argument full identifier (e.g., "-rte")
    DefaultValue  string   // Fallback value for invalid/missing values
    CurrentValue  string   // Processed result value
    AllowedValues []string // Valid values or ANY for unrestricted
}
```

### MultiValueArg
```Go
type MultiValueArg struct {
    Name          string   // Argument full identifier (e.g., "-dcs")
    Separator     string   // Expected value separator (e.g., ",")
    DefaultValues []string // Fallback values for invalid/missing values
    CurrentValues []string // Processed result values
    AllowedValues []string // Valid values or ANY for unrestricted
}
```

### SingleArgs
```Go
type SingleArgs []*SingleValueArg
```

### MultiArgs
```Go
type MultiArgs []*MultiValueArg
```

## Functions
### ProcessSingleValueArgs
Main function for single value arguments batch processing.
```Go
func ProcessSingleValueArgs(args []string, configs SingleArgs, prefix, sign string, size int)
```
**Parameters:**
* `args`: Direct command-line arguments slice
* `configs`: Single value argument configurations to process
* `prefix`: Expected argument prefix (e.g., "-")
* `sign`: Value separator (e.g., "=")
* `size`: Expected argument name length

**Behavior:**
* All configurations initialize with their default value
* Invalid/unknown arguments are silently ignored
* Valid arguments update their `CurrentValue`

### ProcessMultiValueArgs
Main function for multiple value arguments batch processing.
```Go
func ProcessMultiValueArgs(args []string, configs MultiArgs, prefix, sign string, size int)
```
**Parameters:**
* `args`: Direct command-line arguments slice
* `configs`: Multiple value argument configurations to process
* `prefix`: Expected argument prefix (e.g., "-")
* `sign`: Value separator (e.g., "=")
* `size`: Expected argument name length

**Behavior:**
* All configurations initialize with their default values
* Invalid/unknown arguments are silently ignored
* Valid arguments update their `CurrentValues`

### ValidateAndExtractArgument
Low-level argument parsing and validation.
```Go
func ValidateAndExtractArgument(arg, prefix, sign string, size int) (string, string, error)
```
**Returns:** argument full identifier, value, and error.

### ExtractMultipleValues
Extracts multiple values from a single string using a separator.
```Go
func ExtractMultipleValues(fullString, separator string) []string
```
**Returns:** slice of strings containing all extracted values.

### ValidateArgumentValue
Single value constraint checking with automatic fallback.
```Go
func ValidateArgumentValue(value string, argConfig SingleValueArg) string
```
**Returns:** validated value or default if invalid.

### ValidateArgumentValues
Multiple values constraint checking with automatic fallback and deduplication.
```Go
func ValidateArgumentValues(values []string, argConfig MultiValueArg) []string
```
**Returns:** validated values or defaults if none are valid.

## Error types
### InvalidFormatError
Triggered by malformed argument structure.
```Go
type InvalidFormatError struct {
    Argument string
    Reason   string
}
```
### UnknownArgumentError
Triggered by unrecognized arguments.
```Go
type UnknownArgumentError struct {
    Argument string
}
```

## Notes
* All arguments must follow the exact format you specify: `{prefix}{name}{sign}{value}`, otherwise they won't be considered
* Unknown or invalid arguments are silently ignored, so pay attention when configuring each argument
* Both `SingleValueArg` and `MultiValueArg` are guaranteed to have valid values after processing, as long as `DefaultValue` or `DefaultValues` are specified
* I designed this package based on my needs and how I work with command-line arguments, so I tried to make it as flexible and configurable as possible. However, I can't guarantee it will work for everyone's use case
* This package was originally created to complete the [psetup](https://github.com/leojimenezg/psetup) project, but I intend to keep expanding its capabilities as much as I can
