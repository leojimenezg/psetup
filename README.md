# psetup

Go CLI tool for rapidly setting up simple project structures with customizable templates, automating directory creation and basic file generation for streamlined development workflows.

## Features
* **Quick project scaffolding** with predefined directory structures
* **Multi-language support** for different project types (Go, Python, C++, JavaScript, etc.)
* **Template-based file generation** with customizable content
* **License integration** with MIT and Apache options
* **Flexible documentation setup** with README, .gitignore, and license files
* **Configurable project** location and naming

## Installation
**Note:** Currently, the only available way to use this project is by clonning the repository.

To get this project up and runnig on your local machine, follow the next instructions.
### Prerequisites
Before anything else, make sure you have installed **Go 1.24.X** or a newer version in your system.
### Steps
1. **Clone the repository:**
Open your prefered terminal and clone the project to your local machine.
    ```bash
    git clone https://github.com/leojimenezg/psetup.git
    ```
2. **Navigate into the project directory:**
    ```bash
    cd psetup
    ```
3. **Compile the project:**
    ```bash
    go build .
    ```
4. **Run the Application:**
Finally, execute the binary to launch the psetup program.
    ```bash
    ./psetup -nme=my-project -rte=./ -lng=go -lic=mit -dcs=all
    ```

## Command-line options
| Option | Description | Default | Allowed Values |
|--------|-------------|---------|----------------|
| `-nme` | Project name | `new-project` | Any string |
| `-rte` | Creation path | `./` | Any valid path |
| `-lng` | Language/Extension | `go` | `py`, `c`, `java`, `go`, `cpp`, `lua`, `js`, `r`, `txt` |
| `-lic` | License type | `mit` | `mit`, `apache` |
| `-dcs` | Documents to include | `all` | `all`, `license`, `ignore`, `readme` |

## Generated structure
psetup may create the following project structure:
```
your-project/
├── src/
│   └── main.{lng}          # Main file with specified language/extension
├── tests/                  # Tests directory
├── assets/
│   ├── data/              # Data files directory
│   └── images/            # Image assets directory
├── README.md              # Project documentation (if requested)
├── .gitignore            # Git ignore file (if requested)
└── LICENSE               # License file (if requested)
```

## Templates
psetup uses template files for consistent content generation:
* **License templates**: Pre-configured MIT and Apache 2.0 licenses
* **README template**: Basic project documentation structure
* **Gitignore template**: Common ignore patterns for various environments

Template files are located in the `templates/` directory and can be customized as needed.

## Notes
* All configurations have sensible defaults for immediate usage
* Invalid arguments fall back to default values rather than causing errors
* The tool ensures all parent directories exist before creating files
* This project has a previous version made in Lua [project-basic-setup](https://github.com/leojimenezg/project-basic-setup), where the problem was very strictly solved.
* This new version was designed to be as flexible and modular as possible. Also, this was my first project made in Go, so I don't expect it to be perfect

## Useful Resources

* [Go Documentation](https://golang.org/doc/) - Official Go programming language documentation
* [argparse Documentation](./argparse/README.md) - Detailed argparse package documentation
* [itemgen Documentation](./itemgen/README.md) - Detailed itemgen package documentation
