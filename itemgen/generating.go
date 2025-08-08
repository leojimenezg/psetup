package itemgen

import (
	"fmt"
	"os"
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
}

type Configs []ItemConfig

// CreateFile creates a new file based on the provided configuration.
// The function handles extension formatting automatically, adding a dot prefix if not present.
// If a template path is specified, the file content is read from the template file;
// otherwise, an empty file is created.
// Parameters:
//   - item: ItemConfig struct containing file name, extension, type, creation path, and template path
// Returns nil if successful, or an error of type InvalidTypeError, TemplateError, or CreationError.
func CreateFile(item ItemConfig) error {
	if item.Type != FILE { return InvalidTypeError{ Type: item.Type } }
	var fullPath string
	if item.Extension == "" { 
		fullPath = filepath.Join(item.CreationPath, item.Name)
	} else if item.Extension[0] != '.' {
		fullPath = filepath.Join(item.CreationPath, item.Name + "." + item.Extension)
	} else {
		fullPath = filepath.Join(item.CreationPath, item.Name + item.Extension)
	}
	var fileContent[]byte
	if item.TemplatePath != "" { 
		template, errTemplate := os.ReadFile(item.TemplatePath)
		if errTemplate != nil {
			return TemplateError{ Path: item.TemplatePath, Err: errTemplate }
		}
		fileContent = template
	} else {
		fileContent = nil
	}
	errFile := os.WriteFile(fullPath, fileContent, 0644)  // Filemode: permissions and attributes.
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

// CreateItems iterates through a Configs slice and creates each item based on its type.
// The function processes all items regardless of individual failures, collecting any errors encountered.
// Parameters:
//   - items: Configs slice containing all ItemConfig structs to create
// Returns nil if everything is successful, or a slice of all errors encountered.
func CreateItems(items Configs) []error {
	var errors []error
	for _, item := range items {
		switch item.Type {
		case FILE:
			errFile := CreateFile(item)
			if errFile != nil {
				errors = append(errors, errFile)
			}
		case DIR:
			errDir := CreateDirectory(item)
			if errDir != nil {
				errors = append(errors, errDir)
			}
		default:
			errors = append(errors, InvalidTypeError{ Type: item.Type })
		}
	}
	if len(errors) == 0 { return nil }
	return errors
}
