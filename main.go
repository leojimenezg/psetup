package main

import (
	"fmt"
)

var options map[string]string = map[string]string {
	"--nme": "new-project", "--rte": "./", "--lng": "go", "--lic": "mit", "--dcs": "all",
}
var languages map[string]string = map[string]string {
	"python": ".py", "c": ".c", "golang": ".go", "lua": ".lua",
	"java": ".java", "javascript": ".js",  "cpp": ".cpp", "r": ".r",
	"": "",
}
var licences map[string]string = map[string]string {
	"mit": "MIT", "apache": "APACHE",
}
var documents map[string]string = map[string]string {
	"ignore": ".gitignore", "license": "LICENSE", "readme": "README", "all": "",
}

func main() {
	// TODO: Parse command-line argumnents.
	// TODO: Validate options' values.
	fmt.Println("Hello from main package!")
}
