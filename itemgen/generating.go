package itemgen

import (
	"fmt"
	"os"
	"embed"
	"path/filepath"
)

type ItemType string

type InvalidTypeError struct {
	Type ItemType
}
type CreationError struct {
	Name string
	Path string
	Err error
}
type ProcessError struct {
	Name string
	Path string
	Err error
}
type TemplateError struct {
	Path string
	Err error
}

func (e InvalidTypeError) Error() string {
	return fmt.Sprintf("invalid item type: %s", e.Type)
}
func (e CreationError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("failed to create item %s at %s", e.Name, e.Path)
	}
	return fmt.Sprintf("failed to create item %s at %s: %v", e.Name, e.Path, e.Err)
}
func (e ProcessError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("failed to process item %s at %s", e.Name, e.Path)
	}
	return fmt.Sprintf("failed to process item %s at %s: %v", e.Name, e.Path, e.Err)
}
func (e TemplateError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("failed to get content from template at %s", e.Path)
	}
	return fmt.Sprintf("failed to get content from template at %s: %v", e.Path, e.Err)
}

const FILE ItemType = "file"
const DIR ItemType = "directory"

type ItemConfig struct {
	Name string
	Extension string
	Type ItemType
	CreationPath string
	TemplatePath string
	TemplateContent []byte
}

type Items []ItemConfig

func createFullPath(path, name, ext string) string {
	var fullPath string
	if ext == "" { 
		fullPath = filepath.Join(path, name)
	} else if ext[0] != '.' {
		fullPath = filepath.Join(path, name + "." + ext)
	} else {
		fullPath = filepath.Join(path, name + ext)
	}
	return fullPath
}

// CreateFile creates a new file based on the provided configuration.
// The function handles extension formatting automatically, adding a dot prefix if not present.
// Content priority: If template content is provided, it is used directly. If not, but a template 
// path is specified, the file content is read from the template file. Otherwise, an empty file is created.
// Parameters:
//   - item: ItemConfig struct containing all necessary configurations
// Returns nil if successful, or an error of type InvalidTypeError, TemplateError, or CreationError.
func CreateFile(item ItemConfig) error {
	if item.Type != FILE { return InvalidTypeError{ Type: item.Type } }
	fullPath := createFullPath(item.CreationPath, item.Name, item.Extension)
	if item.TemplateContent != nil {
		errFile := os.WriteFile(fullPath, item.TemplateContent, 0644)
		if errFile != nil { return CreationError{ Name: item.Name, Path: fullPath, Err: errFile } }
		return nil
	}
	var template[]byte = nil
	if item.TemplatePath != "" { 
		data, errTemplate := os.ReadFile(item.TemplatePath)
		if errTemplate != nil { return TemplateError{ Path: item.TemplatePath, Err: errTemplate } }
		template = data
	}
	errFile := os.WriteFile(fullPath, template, 0644)  // Filemode: permissions and attributes.
	if errFile != nil { return CreationError{ Name: item.Name, Path: fullPath, Err: errFile } }
	return nil
}

// CreateFileEmbed creates a new file based on the provided configuration using an embedded filesystem.
// The function handles extension formatting automatically, adding a dot prefix if not present.
// Content priority: If template content is provided, it is used directly. If not, but a template 
// path is specified, the file content is read from the embedded filesystem. Otherwise, an empty file is created.
// Parameters:
//   - item: ItemConfig struct containing all necessary configurations
//   - fs: Pointer to embed.FS containing the template files
// Returns nil if successful, or an error of type InvalidTypeError, TemplateError, or CreationError.
func CreateFileEmbed(item ItemConfig, fs *embed.FS) error {
	if item.Type != FILE { return InvalidTypeError{ Type: item.Type } }
	fullPath := createFullPath(item.CreationPath, item.Name, item.Extension)
	if item.TemplateContent != nil {
		errFile := os.WriteFile(fullPath, item.TemplateContent, 0644)
		if errFile != nil { return CreationError{ Name: item.Name, Path: fullPath, Err: errFile } }
		return nil
	}
	var template []byte = nil
	if item.TemplatePath != "" {
		data, errTemplate := fs.ReadFile(item.TemplatePath)
		if errTemplate != nil { return TemplateError{ Path: item.TemplatePath, Err: errTemplate } }
		template = data
	}
	errFile := os.WriteFile(fullPath, template, 0644)
	if errFile != nil { return CreationError{ Name: item.Name, Path: fullPath, Err: errFile } }
	return nil
}

// CreateDirectory creates a new directory and necessary parent directories based on the provided configuration.
// The function uses os.MkdirAll to ensure all parent directories in the path are created if they don't exist.
// Parameters:
//   - item: ItemConfig struct containing directory name, type, and creation path
// Returns nil if successful, or an error of type InvalidTypeError, or CreationError.
func CreateDirectory(item ItemConfig) error {
	if item.Type != DIR { return InvalidTypeError{ Type: item.Type } }
	fullPath := filepath.Join(item.CreationPath, item.Name)
	errDir := os.MkdirAll(fullPath, 0755)
	if errDir != nil { return CreationError{ Name: item.Name, Path: fullPath, Err: errDir } }
	return nil
}

// CreateItems iterates through an Items slice and creates each item based on its type.
// The function processes all items regardless of individual failures, collecting any errors encountered.
// Parameters:
//   - items: Items slice containing all ItemConfig structs to create
// Returns nil if everything is successful, or a slice of all errors encountered.
func CreateItems(items Items) []error {
	var errors []error
	for _, item := range items {
		switch item.Type {
		case FILE:
			errFile := CreateFile(item)
			if errFile != nil { errors = append(errors, errFile) }
		case DIR:
			errDir := CreateDirectory(item)
			if errDir != nil { errors = append(errors, errDir) }
		default:
			errors = append(errors, InvalidTypeError{ Type: item.Type })
		}
	}
	if len(errors) == 0 { return nil }
	return errors
}

// CreateItemsEmbed iterates through an Items slice and creates each item based on its type using an embedded filesystem.
// The function processes all items regardless of individual failures, collecting any errors encountered.
// Parameters:
//   - items: Items slice containing all ItemConfig structs to create
//   - fs: Pointer to embed.FS containing the template files for FILE items
// Returns nil if everything is successful, or a slice of all errors encountered.
func CreateItemsEmbed(items Items, fs *embed.FS) []error {
	var errors []error
	for _, item := range items {
		switch item.Type {
		case FILE:
			errFile := CreateFileEmbed(item, fs)
			if errFile != nil { errors = append(errors, errFile) }
		case DIR:
			errDir := CreateDirectory(item)
			if errDir != nil { errors = append(errors, errDir) }
		default:
			errors = append(errors, InvalidTypeError{ Type: item.Type })
		}
	}
	if len(errors) == 0 { return nil }
	return errors
}
