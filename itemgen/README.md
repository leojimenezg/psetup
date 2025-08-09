# itemgen

Go package for generating directories and files with configurable template support. Built around the `ItemConfig` structure for flexible file system item creation.

## Features
* **Dual item support** for both files and directories with type-safe operations
* **Template-based file generation** with automatic content injection from template files
* **Intelligent extension handling** with automatic dot prefix management
* **Recursive directory creation** ensuring parent directories exist
* **Batch processing** of multiple items in a single operation
* **Comprehensive error handling** with detailed failure information

## Quick start
```Go
import "github.com/leojimenezg/psetup/itemgen"

// Configure a file with template
mainFile := itemgen.ItemConfig{
    Name: "main",
    Extension: "go", 
    Type: itemgen.FILE,
    CreationPath: "./src",
    TemplatePath: "./templates/main.txt",
}

// Create the file
err := itemgen.CreateFile(mainFile)
if err != nil {
    fmt.Println(err)
}
```

## Core types
### ItemType
```Go
type ItemType string

const FILE ItemType = "file"
const DIR ItemType = "directory"
```

### ItemConfig
```Go
type ItemConfig struct {
    Name         string   // Item name without extension
    Extension    string   // File extension (files only)
    Type         ItemType // FILE or DIR
    CreationPath string   // Absolute creation path
    TemplatePath string   // Template file path (files only)
}
```

### Configs
```Go
type Configs []ItemConfig
```

## Functions
### CreateFile
Creates files with optional template content.
```Go
func CreateFile(item ItemConfig) error
```
**Behavior:**
* Validates item type is `FILE`
* Handles extension formatting automatically
* Reads template content if `TemplatePath` specified
* Creates empty file if no template provided
* Uses `0644` permissions

### CreateDirectory
Creates directories with parent path support.
```Go
func CreateDirectory(item ItemConfig) error
```
**Behavior:**
* Validates item type is `DIR`
* Creates all necessary parent directories
* Uses `0755` permissions

### CreateItems
Batch processes multiple items regardless of individual failures.
```Go
func CreateItems(items Configs) []error
```
**Returns:** nil if all successful, or slice of all errors encountered.

## Error types
### InvalidTypeError
Triggered by invalid ItemConfig Type.
```Go
type InvalidTypeError struct {
    Type ItemType
}
```

### CreationError
Triggered when the item could not be created.
```Go
type CreationError struct {
    Name string
    Path string
    Err  error
}
```

### TemplateError
Triggered when the template could not be retrieved.
```Go
type TemplateError struct {
    Path string
    Err  error
}
```

## Notes
* File permissions are set to `0644` (read/write for owner, read for group/others)
* Directory permissions are set to `0755` (full access for owner, read/execute for others)
* Extension handling automatically adds dot prefix if missing
* CreateItems continues processing even if individual items fail
* This package was originally created to complete the [psetup](https://github.com/leojimenezg/psetup) project, however, I intend to expand its capabilities as much as I can in the future
