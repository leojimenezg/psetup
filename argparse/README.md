# argparse

Go package for validating and extracting command-line arguments with customizable configurations. Built around the `ArgumentConfig` struct for flexible and detailed argument handling.

## Features
* **Structured argument validation** with configurable format: prefix, separator, and fixed name size
* **Flexible value constraints** using an allowed values slice or ANY constant for unrestricted input
* **In-place modifications** to ArgumentConfig structs for reusability across operations
* **Automatic fallback** to default values for missing or invalid argument values
* **Custom error types** for detailed validation and parsing error handling
* **Batch processing** of multiple arguments in a single function call

## Quick start
```Go
import "github.com/leojimenezg/psetup/argparse"

// Configure expected argument
routeArg := argparse.ArgumentConfig{
    Name: "-rte",
    DefaultValue: "./",
    AllowedValues: []string{ argparse.ANY },
}

// Process command-line arguments
args := os.Args[1:]
configs := argparse.Arguments{ &routeArg }
argparse.ProcessArguments(args, configs, "-", "=", 3)

// Access result
fmt.Println(routeArg.CurrentValue)
```

## Core types
### ArgumentConfig
```Go
type ArgumentConfig struct {
    Name          string   // Argument full identifier (e.g., "-rte")
    DefaultValue  string   // Fallback value for invalid/missing values
    CurrentValue  string   // Processed result value
    AllowedValues []string // Valid values or [ANY] for unrestricted
}
```
### Arguments
```Go
type Arguments []*ArgumentConfig
```

## Functions
### ProcessArguments
Main function for batch argument processing.
```Go
func ProcessArguments(args []string, configs Arguments, prefix, sign string, size int)
```
**Parameters:**
* `args`: Direct command-line arguments slice
* `configs`: Arguments configurations to process
* `prefix`: Expected argument prefix (e.g., "-")
* `sign`: Value separator (e.g., "=")
* `size`: Expected argument name length

**Behavior:**
* All configurations initialize with their default value
* Invalid/unknown arguments are silently ignored
* Valid arguments update their `CurrentValue`

### ValidateAndExtractArgument
Low-level argument parsing and validation.
```Go
func ValidateAndExtractArgument(arg, prefix, sign string, size int) (string, string, error)
```
**Returns:** argument identifier, value, and error.

### ValidateArgumentValue
Value constraint checking with automatic fallback.
```Go
func ValidateArgumentValue(value string, argConfig ArgumentConfig) string
```
**Returns:** validated value or default if invalid.

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
* Arguments must follow exact format: `{prefix}{name}{sign}{value}`, otherwise, they wont be considered
* The `size` parameter enforces consistent argument name lengths, however, in the future, it is planned to allow more flexibility for each argument
* Unknown or invalid arguments are silently ignored, therefore, pay attention when configuring each argument
* All `ArgumentConfig`s are guaranteed to have valid values after processing, as long as a `DefaultValue` is specified
* This package was orignially created in order to complete the [psetup](https://github.com/leojimenezg/psetup) project, however, I intend to expand its capabilities as much as I can in the future
