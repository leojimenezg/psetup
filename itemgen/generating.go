package itemgen

import ()

type ItemType string

const FILE ItemType = "file"
const DIR ItemType = "directory"

type ItemConfig struct {
	Name string
	Type ItemType
	CreationPath string
	TemplatePath string
}

type Configs []*ItemConfig

func CreateFile(config *ItemConfig) bool {
	return true
}

func CreateDirectory(config *ItemConfig) bool {
	return true
}

func CreateItem(config *ItemConfig) bool {
	switch (config.Type) {
	case FILE:
		return CreateFile(config)
	case DIR:
		return CreateDirectory(config)
	default:
		return false
	}
}
