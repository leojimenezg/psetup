package main

import (
	"os"
	"fmt"
	"github.com/leojimenezg/psetup/argparse"
	"github.com/leojimenezg/psetup/itemgen"
)

const OPTION_PREFIX = "-"
const OPTION_SIZE = 3
const OPTION_SIGN = "="

func main() {
	commandLineArgs := os.Args[1:]
	arguments := argparse.Arguments {
		&argparse.ArgumentConfig{
			Name: OPTION_PREFIX + "nme",
			DefaultValue: "new-project",
			AllowedValues: []string { argparse.ANY },
		},
		&argparse.ArgumentConfig{
			Name: OPTION_PREFIX + "rte",
			DefaultValue: "./",
			AllowedValues: []string { argparse.ANY },
		},
		&argparse.ArgumentConfig{
			Name: OPTION_PREFIX + "lng",
			DefaultValue: "go",
			AllowedValues: []string { "py", "c", "java", "go", "cpp", "lua", "js", "r", "txt" },
		},
		&argparse.ArgumentConfig{
			Name: OPTION_PREFIX + "lic",
			DefaultValue: "mit",
			AllowedValues: []string { "mit", "apache" },
		},
		&argparse.ArgumentConfig{
			Name: OPTION_PREFIX + "dcs",
			DefaultValue: "all",
			AllowedValues: []string { "all", "license", "ignore", "readme" },
		},
	}
	argparse.ProcessArguments(commandLineArgs, arguments, OPTION_PREFIX, OPTION_SIGN, OPTION_SIZE)
	var rteArg argparse.ArgumentConfig
	var lngArg argparse.ArgumentConfig
	for _, arg := range arguments {
		if arg.Name == "-rte" { rteArg = *arg }
		if arg.Name == "-lng" { lngArg = *arg }
	}
	directories := itemgen.Configs{
		{ Name: "dir1", Type: itemgen.DIR, CreationPath: rteArg.CurrentValue },
		{ Name: "dir2", Type: itemgen.DIR, CreationPath: rteArg.CurrentValue },
		{ Name: "dir3", Type: itemgen.DIR, CreationPath: rteArg.CurrentValue },
	}
	files := itemgen.Configs{
		{ 
			Name: "file1", Extension: lngArg.CurrentValue, Type: itemgen.FILE, 
			CreationPath: rteArg.CurrentValue + "/dir1", TemplatePath: "./templates/ignore.txt",
		},
		{ 
			Name: "file2", Extension: lngArg.CurrentValue, Type: itemgen.FILE, 
			CreationPath: rteArg.CurrentValue + "/dir2", TemplatePath: "./templates/readme.txt",
		},
		{ 
			Name: "file3", Extension: lngArg.CurrentValue, Type: itemgen.FILE, 
			CreationPath: rteArg.CurrentValue + "/dir3", TemplatePath: "./templates/license/mit.txt",
		},
	}
	errDirs := itemgen.CreateItems(directories)
	if errDirs != nil {
		fmt.Printf("could not create all directories: %v", errDirs)
		return
	}
	fmt.Println("all directories successfully created")
	errFiles := itemgen.CreateItems(files)
	if errFiles != nil {
		fmt.Printf("could not create all files: %v", errFiles)
		return
	}
	fmt.Println("all files successfully created")
}
