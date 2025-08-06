package main

import (
	"fmt"
	"os"
	"github.com/leojimenezg/psetup/parsing"
)

const OPTION_PREFIX = "-"
const OPTION_SIZE = 3
const OPTION_SIGN = "="

func main() {
	commandLineArgs := os.Args[1:]
	arguments := parsing.Arguments {
		&parsing.ArgumentConfig{
			Name: "-nme",
			DefaultValue: "new-project",
			AllowedValues: []string { "any" },
		},
		&parsing.ArgumentConfig{
			Name: "-rte",
			DefaultValue: "new-project",
			AllowedValues: []string { "any" },
		},
		&parsing.ArgumentConfig{
			Name: "-lng",
			DefaultValue: "go",
			AllowedValues: []string { "py", "c", "java", "go", "cpp", "lua", "js", "r" },
		},
		&parsing.ArgumentConfig{
			Name: "-lic",
			DefaultValue: "mit",
			AllowedValues: []string { "mit", "apache" },
		},
		&parsing.ArgumentConfig{
			Name: "-dcs",
			DefaultValue: "all",
			AllowedValues: []string { "all", "license", "ignore", "readme" },
		},
	}
	fmt.Println("Before Processing")
	for _, argConfig := range arguments {
		fmt.Printf("%s = %s\n", argConfig.Name, argConfig.CurrentValue)
	}
	parsing.ProcessArguments(commandLineArgs, arguments, OPTION_PREFIX, OPTION_SIGN, OPTION_SIZE)
	fmt.Println("After Processing")
	for _, argConfig := range arguments {
		fmt.Printf("%s = %s\n", argConfig.Name, argConfig.CurrentValue)
	}
}
