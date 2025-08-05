package main

import (
	"fmt"
	"os"
	"strings"
)

const OPTION_PREFIX = "-"
const OPTION_SIZE = 3
const OPTION_SIGN = "="

var options map[string]string = map[string]string {
	OPTION_PREFIX + "nme": "new-project",
	OPTION_PREFIX + "rte": "./",
	OPTION_PREFIX + "lng": "go",
	OPTION_PREFIX + "lic": "mit",
	OPTION_PREFIX + "dcs": "all",
}
var languages map[string]string = map[string]string {
	"python": ".py", "c": ".c", "golang": ".go", "lua": ".lua",
	"java": ".java", "javascript": ".js",  "cpp": ".cpp", "r": ".r",
}
var licences map[string]string = map[string]string {
	"mit": "MIT", "apache": "APACHE",
}
var documents map[string]string = map[string]string {
	"ignore": ".gitignore", "license": "LICENSE", "readme": "README", "all": "",
}

// ParseFullArgument parses and validates a command-line argument with the specified format.
// Parameters:
//   - arg: the full command-line argument to parse
//   - prefix: expected argument prefix (e.g., "--")
//   - sign: separator between argument and value (e.g., "=")
//   - size: length of the argument name
// Returns the argument name and its value if valid, empty strings otherwise.
func ParseFullArgument(arg, prefix, sign string, size int) (string, string) {
	expected_size := len(prefix) + size
	if len(arg) < expected_size { return "", "" }
	if !strings.HasPrefix(arg, prefix) { return "", "" }
	assignator_idx := strings.Index(arg, sign)
	if assignator_idx != expected_size { return "", "" }
	return arg[:assignator_idx], arg[assignator_idx + 1:]
}

// ValidateArguments goes through each argument from os.Args, compares them with the
// received map and assigns its value to the validated option.
// Parameters:
//   - opts: pointer to a map to compare and assing.
// Returns nothing.
func ValidateArguments(opts *map[string]string) {
	arguments := os.Args[1:]
	if len(arguments) < 2 { return }
	for _, full_arg := range(arguments) {
		arg, val := ParseFullArgument(full_arg, OPTION_PREFIX, OPTION_SIGN, OPTION_SIZE)
		fmt.Printf("Argument: %s, Value: %s\n", arg, val)
	}
}

func main() {
	// TODO: Validate options' values.
	ValidateArguments(&options)
}
