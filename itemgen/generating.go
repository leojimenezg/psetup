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

func CreateDirectory(item ItemConfig) error {
	if item.Type != DIR { return InvalidTypeError{ Type: item.Type } }
	fullPath := filepath.Join(item.CreationPath, item.Name)
	errDir := os.MkdirAll(fullPath, 0755)
	if errDir != nil { return CreationError{ Name: item.Name, Path: fullPath, Err: errDir } }
	return nil
}

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
