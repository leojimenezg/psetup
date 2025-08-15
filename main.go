package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"github.com/leojimenezg/psetup/argparse"
	"github.com/leojimenezg/psetup/itemgen"
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
		AllowedValues: []string { argparse.ANY } }
	routeArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "rte",
		DefaultValues: []string{ "./" },
		AllowedValues: []string { argparse.ANY } }
	languageArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "lng",
		DefaultValues: []string{ "go" },
		AllowedValues: []string { "py", "c", "java", "go", "cpp", "lua", "js", "r", "txt" } }
	licenseArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "lic",
		DefaultValues: []string{ "mit" },
		AllowedValues: []string { "mit", "apache" } }
	documentsArg := argparse.MultiValueArg{
		Name: OPTION_PREFIX + "dcs",
		DefaultValues: []string{ "all" },
		AllowedValues: []string { "all", "license", "ignore", "readme" } }
	argumentsConfigs := argparse.MultiArgs{ &nameArg, &routeArg, &languageArg, &licenseArg, &documentsArg }
	argparse.ProcessMultiValueArgs(os.Args[1:], argumentsConfigs, OPTION_PREFIX, OPTION_SIGN, OPTION_SIZE)
	projectDir := filepath.Join(routeArg.CurrentValues[0], nameArg.CurrentValues[0])
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
			Name: "main", Extension: languageArg.CurrentValues[0], Type: itemgen.FILE,
			CreationPath: filepath.Join(projectDir, "src") },
		"license": {
			Name: "LICENSE", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/license/" + licenseArg.CurrentValues[0] + ".txt" },
		"ignore": {
			Name: ".gitignore", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/ignore.txt" },
		"readme": {
			Name: "README", Extension: "md", Type: itemgen.FILE, CreationPath: projectDir,
			TemplatePath: "templates/readme.txt" } }
	var files itemgen.Items
	files = append(files, documentsConfigs["main"])
	all := slices.Contains(documentsArg.CurrentValues, "all")
	for _, document := range documentsArg.AllowedValues[1:] {
		if !all { if !slices.Contains(documentsArg.CurrentValues, document) { continue } }
		files = append(files, documentsConfigs[document])
	}
	errsFiles := itemgen.CreateItemsEmbed(files, &templatesFS)
	if errsFiles != nil { fmt.Printf("could not create all files: %v\n", errsFiles) }
	fmt.Println("psetup: project structure successfully created!")
}
