package main

import (
	"os"
	"fmt"
	"embed"
	"path/filepath"
	"github.com/leojimenezg/psetup/argparse"
	"github.com/leojimenezg/psetup/itemgen"
)

//go:embed templates/*.txt templates/license/*.txt
var templatesFS embed.FS

const OPTION_PREFIX = "-"
const OPTION_SIZE = 3
const OPTION_SIGN = "="

func main() {
	nameArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "nme",
		DefaultValue: "new-project",
		AllowedValues: []string { argparse.ANY },
	}
	routeArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "rte",
		DefaultValue: "./",
		AllowedValues: []string { argparse.ANY },
	}
	languageArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "lng",
		DefaultValue: "go",
		AllowedValues: []string { "py", "c", "java", "go", "cpp", "lua", "js", "r", "txt" },
	}
	licenseArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "lic",
		DefaultValue: "mit",
		AllowedValues: []string { "mit", "apache" },
	}
	documentsArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "dcs",
		DefaultValue: "all",
		AllowedValues: []string { "all", "license", "ignore", "readme" },
	}
	commandLineArgs := os.Args[1:]
	argumentsConfigs := argparse.Arguments{ &nameArg, &routeArg, &languageArg, &licenseArg, &documentsArg }
	argparse.ProcessArguments(commandLineArgs, argumentsConfigs, OPTION_PREFIX, OPTION_SIGN, OPTION_SIZE)
	projectDir := filepath.Join(routeArg.CurrentValue, nameArg.CurrentValue)
	directories := itemgen.Items{
		{ Name: "src", Type: itemgen.DIR, CreationPath: projectDir },
		{ Name: "tests", Type: itemgen.DIR, CreationPath: projectDir },
		{ Name:	"assets", Type: itemgen.DIR, CreationPath: projectDir },
		{ Name:	"data", Type: itemgen.DIR, CreationPath: filepath.Join(projectDir, "assets") },
		{ Name:	"images", Type: itemgen.DIR, CreationPath: filepath.Join(projectDir, "assets") },
	}
	errsDir := itemgen.CreateItems(directories)
	if errsDir != nil { fmt.Printf("could not create all directories: %v\n", errsDir); return }
	// Files are gotten from the Virtual File System (embed.FS)
	var files itemgen.Items
	mainFile := itemgen.ItemConfig{
		Name: "main", Extension: languageArg.CurrentValue, Type: itemgen.FILE,
		CreationPath: filepath.Join(projectDir, "src") }
	files = append(files, mainFile)
	switch documentsArg.CurrentValue {
	case "license":
		licenceFile := itemgen.ItemConfig{
			Name: "LICENSE", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/license/" + licenseArg.CurrentValue + ".txt" }
		files = append(files, licenceFile)
	case "ignore":
		ignoreFile := itemgen.ItemConfig{
			Name: ".gitignore", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/ignore.txt" }
		files = append(files, ignoreFile)
	case "readme":
		readmeFile := itemgen.ItemConfig{ 
			Name: "README", Extension: "md", Type: itemgen.FILE,
			CreationPath: projectDir, TemplatePath: "templates/readme.txt" }
		files = append(files, readmeFile)
	default:
		licenceFile := itemgen.ItemConfig{
			Name: "LICENSE", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/license/" + licenseArg.CurrentValue + ".txt" }
		files = append(files, licenceFile)
		ignoreFile := itemgen.ItemConfig{
			Name: ".gitignore", Type: itemgen.FILE, CreationPath: projectDir, TemplatePath: "templates/ignore.txt" }
		files = append(files, ignoreFile)
		readmeFile := itemgen.ItemConfig{ 
			Name: "README", Extension: "md", Type: itemgen.FILE,
			CreationPath: projectDir, TemplatePath: "templates/readme.txt" }
		files = append(files, readmeFile)
	}
	errsFiles := itemgen.CreateItemsEmbed(files, &templatesFS)
	if errsFiles != nil { fmt.Printf("could not create all files: %v\n", errsFiles) }
	fmt.Println("psetup: project structure successfully created!")
}
