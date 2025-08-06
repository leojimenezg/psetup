package main

import (
	"os"
	"github.com/leojimenezg/psetup/argparse"
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
			AllowedValues: []string { "py", "c", "java", "go", "cpp", "lua", "js", "r" },
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
}
