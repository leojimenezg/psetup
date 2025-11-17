package main

import (
	"os"
	"fmt"
	"embed"
	"slices"
	"path/filepath"
	"github.com/leojimenezg/psetup/itemgen"
	"github.com/leojimenezg/psetup/argparse"
)

//go:embed templates/*.txt templates/license/*.txt
var templatesFS embed.FS

const OPTION_PREFIX = "-"
const OPTION_SIZE = 3
const OPTION_SIGN = "="

func main() {
	nameArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "nme",
		DefaultValues: []string{ "new-project" },
		AllowedValues: []string { argparse.ANY },
	}

	routeArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "rte",
		DefaultValues: []string{ "./" },
		AllowedValues: []string { argparse.ANY },
	}

	languageArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "lng",
		DefaultValues: []string{ "go" },
		AllowedValues: []string { "py", "c", "java", "go", "cpp", "lua", "js", "r", "txt" },
	}

	licenseArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "lic",
		DefaultValues: []string{ "mit" },
		AllowedValues: []string { "mit", "apache" },
	}

	documentsArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "dcs",
		DefaultValues: []string{ "all" },
		AllowedValues: []string { "all", "license", "ignore", "readme" },
	}

	argsConfigs := argparse.MultiArgs{
		&nameArg,
		&routeArg,
		&languageArg,
		&licenseArg,
		&documentsArg,
	}

	// The os.Args[0] corresponds to the program name, so its omitted
	argparse.ProcessMultiValueArgs(os.Args[1:], argsConfigs, OPTION_PREFIX, OPTION_SIGN, OPTION_SIZE)

	projectDirectory := filepath.Join(routeArg.CurrentValues[0], nameArg.CurrentValues[0])

	directories := itemgen.Items{
		{ Name: "src", Type: itemgen.DIR, CreationPath: projectDirectory },
		{ Name:	"data", Type: itemgen.DIR, CreationPath: projectDirectory },
		{ Name: "tests", Type: itemgen.DIR, CreationPath: projectDirectory },
		{ Name:	"public", Type: itemgen.DIR, CreationPath: projectDirectory },
	}
	directoriesErr := itemgen.CreateItems(directories)
	if directoriesErr != nil {
		fmt.Printf("could not create all needed directories: %v\n", directoriesErr)
		return
	}

	documentsConfigs := map[string]itemgen.ItemConfig{
		"main": {
			Name: "main", Extension: languageArg.CurrentValues[0], Type: itemgen.FILE,
			CreationPath: filepath.Join(projectDirectory, "src"),
		},
		"license": {
			Name: "LICENSE", Type: itemgen.FILE, CreationPath: projectDirectory,
			TemplatePath: "templates/license/" + licenseArg.CurrentValues[0] + ".txt",
		},
		"ignore": {
			Name: ".gitignore", Type: itemgen.FILE,
			CreationPath: projectDirectory, TemplatePath: "templates/ignore.txt",
		},
		"readme": {
			Name: "README", Extension: "md", Type: itemgen.FILE,
			CreationPath: projectDirectory, TemplatePath: "templates/readme.txt",
		},
	}

	var files itemgen.Items
	files = append(files, documentsConfigs["main"])
	all := slices.Contains(documentsArg.CurrentValues, "all")
	for _, document := range documentsArg.AllowedValues[1:] {
		if !all {
			if !slices.Contains(documentsArg.CurrentValues, document) { continue }
		}
		files = append(files, documentsConfigs[document])
	}
	filesErr := itemgen.CreateItemsEmbed(files, &templatesFS)
	if filesErr != nil {
		fmt.Printf("could not create all needed files: %v\n", filesErr)
		return
	}

	fmt.Println("psetup: project structure successfully created!")
}
