# itemgen

Go package for generating directories and files with configurable template support. Built around the `ItemConfig` structure for flexible item creation with embedded filesystem support for portable binaries.

## Features
* **Dual item support** for both files and directories with type-safe operations
* **Template-based file generation** with automatic content injection from template files or embedded filesystems
* **Embedded filesystem support** for portable binaries with no external dependencies
* **Intelligent extension handling** with automatic dot prefix management
* **Recursive directory creation** ensuring parent directories exist
* **Batch processing** of multiple items in a single operation
* **Comprehensive error handling** with detailed failure information
* **Multiple content sources** supporting direct content, file paths, and embedded templates

## Quick start
### Using embedded templates (recommended)
```Go
import (
    "embed"
    "github.com/leojimenezg/psetup/itemgen"
)

//go:embed templates/*.txt
var templatesFS embed.FS

// Configure a file with embedded template
mainFile := itemgen.ItemConfig{
    Name: "main",
    Extension: "go", 
    Type: itemgen.FILE,
    CreationPath: "./src",
    TemplatePath: "templates/main.txt",
}

// Create the file using embedded filesystem
err := itemgen.CreateFileEmbed(mainFile, &templatesFS)
if err != nil {
    fmt.Println(err)
}
```
### Using direct content
```Go
import "github.com/leojimenezg/psetup/itemgen"

// Configure a file with direct content
configFile := itemgen.ItemConfig{
    Name: "config",
    Extension: "json", 
    Type: itemgen.FILE,
    CreationPath: "./",
    TemplateContent: []byte(`{"version": "1.0"}`),
}

// Create the file
err := itemgen.CreateFile(configFile)
if err != nil {
    fmt.Println(err)
}
```
### Using direct path
```Go
import "github.com/leojimenezg/psetup/itemgen"

// Configure a file with direct path
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
    Name            string   // Item name without extension
    Extension       string   // File extension (files only)
    Type            ItemType // FILE or DIR
    CreationPath    string   // Parent directory path
    TemplatePath    string   // Template file path (optional)
    TemplateContent []byte   // Template content (optional)
}
```

### Items
```Go
type Items []ItemConfig
```

## Functions
### CreateFile
Creates files using the provided content or path following certain priority.
```Go
func CreateFile(item ItemConfig) error
```
**Content Priority:**
* `TemplateContent` (if provided) - uses direct byte content
* `TemplatePath` (if provided) - reads from direct path
* `Empty file` (if neither provided) - empty also when reading from path causes error or no content

**Behavior:**
* Validates item type is `FILE`
* Handles extension formatting automatically
* Uses `0644` permissions

### CreateFileEmbed
Creates files using the provided content or path from the embedded filesystem, following certain priority.
```Go
func CreateFileEmbed(item ItemConfig, fs *embed.FS) error
```
**Content Priority:**
* `TemplateContent` (if provided) - uses direct byte content
* `TemplatePath` (if provided) - reads from embedded filesystem
* `Empty file` (if neither provided) - empty also when reading from path causes error or no content

**Behavior:**
* Validates item type is `FILE`
* Handles extension formatting automatically
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
func CreateItems(items Items) []error
```
**Returns:** nil if all successful, or slice of all errors encountered.

### CreateItemsEmbed
Batch processes multiple items regardless of individual failures using embedded filesystems.
```Go
func CreateItemsEmbed(items Items, fs *embed.FS) []error
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
* The embedded filesystem approach was a major improvement over file system dependencies, as it makes distribution much simpler
* The decision to maintain both embedded and legacy functions was driven by backward compatibility, though I highly recommend using the embedded versions, as they can handle the same cases as the normal versions but using paths from the embedded filesystems
* Extension handling was implemented because, users (me) sometimes forget the dot, sometimes include it, so automatic handling saves over thinking
* This package was originally created to complete the [psetup](https://github.com/leojimenezg/psetup) project, however, I intend to expand its capabilities as much as I can in the future
