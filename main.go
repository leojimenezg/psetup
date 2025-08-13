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
		AllowedValues: []string { argparse.ANY } }
	routeArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "rte",
		DefaultValue: "./",
		AllowedValues: []string { argparse.ANY } }
	languageArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "lng",
		DefaultValue: "go",
		AllowedValues: []string { "py", "c", "java", "go", "cpp", "lua", "js", "r", "txt" } }
	licenseArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "lic",
		DefaultValue: "mit",
		AllowedValues: []string { "mit", "apache" } }
	documentsArg := argparse.ArgumentConfig{
		Name: OPTION_PREFIX + "dcs",
		DefaultValue: "all",
		AllowedValues: []string { "all", "license", "ignore", "readme" } }
	argumentsConfigs := argparse.Arguments{ &nameArg, &routeArg, &languageArg, &licenseArg, &documentsArg }
	argparse.ProcessArguments(os.Args[1:], argumentsConfigs, OPTION_PREFIX, OPTION_SIGN, OPTION_SIZE)
	projectDir := filepath.Join(routeArg.CurrentValue, nameArg.CurrentValue)
	directories := itemgen.Items{
		{ Name: "src", Type: itemgen.DIR, CreationPath: projectDir },
		{ Name: "tests", Type: itemgen.DIR, CreationPath: projectDir },
		{ Name:	"assets", Type: itemgen.DIR, CreationPath: projectDir },
		{ Name:	"data", Type: itemgen.DIR, CreationPath: filepath.Join(projectDir, "assets") },
		{ Name:	"images", Type: itemgen.DIR, CreationPath: filepath.Join(projectDir, "assets") } }
	errsDir := itemgen.CreateItems(directories)
	if errsDir != nil { fmt.Printf("could not create all directories: %v\n", errsDir); return }
	documentsConfigs := map[string]itemgen.ItemConfig{
		"main": {
			Name: "main", Extension: languageArg.CurrentValue, Type: itemgen.FILE,
			CreationPath: filepath.Join(projectDir, "src") },
		"license": {
			Name: "LICENSE", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/license/" + licenseArg.CurrentValue + ".txt" },
		"ignore": {
			Name: ".gitignore", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/ignore.txt" },
		"readme": {
			Name: "README", Extension: "md", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/readme.txt" } }
	var files itemgen.Items
	files = append(files, documentsConfigs["main"])
	for _, value := range documentsArg.AllowedValues {
		if value == "all" { continue }
		if documentsArg.CurrentValue != "all" && documentsArg.CurrentValue != value { continue }
		files = append(files, documentsConfigs[value])
	}
	errsFiles := itemgen.CreateItemsEmbed(files, &templatesFS)
	if errsFiles != nil { fmt.Printf("could not create all files: %v\n", errsFiles) }
	fmt.Println("psetup: project structure successfully created!")
}
