# psetup

Go CLI tool for rapidly setting up simple project structures with customizable templates, automating directory creation and basic file generation for streamlined development workflows.

## Features
* **Quick project scaffolding** with predefined directory structures
* **Multi-language support** for different project types (Go, Python, C++, JavaScript, etc.)
* **Template-based file generation** with customizable content
* **Embedded filesystem** to remove any external file dependencies
* **License integration** with MIT and Apache options
* **Multiple value arguments** for selecting multiple documentation types in a single command
* **Flexible documentation setup** with selective generation of README, .gitignore, and license files
* **Configurable project** location and naming

## Installation
**Note:** To use this project, you can either clone the entire repository, or download the executable binary. However this binary is exclusively compiled for macOS, for now.

### Option 1: Download binary (recommended)
1. **Download the binary:** Visit the release [psetup v2.0.0](https://github.com/leojimenezg/psetup/releases/tag/v2.0.0) and download the `psetup-macos-arm64` binary.
2. **Make it executable in your system:**
    ```bash
    chmod +x psetup
    ```
3. **Run the program:**
    ```bash
    ./psetup -nme=my-project -rte=./ -lng=go -lic=mit -dcs=all
    ```

### Option 2: Build from source
#### Prerequisites
Before anything else, make sure you have installed **Go 1.24.X** or a newer version in your system.
#### Steps
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
The program may create the following project structure:
```
your-project/
├── src/
│   └── main.{lng}          # Main file with specified language/extension
├── tests/                  # Tests directory
├── assets/
│   ├── data/               # Data files directory
│   └── images/             # Image assets directory
├── README.md               # Project documentation (if requested)
├── .gitignore              # Git ignore file (if requested)
└── LICENSE                 # License file (if requested)
```

## Templates
**Embedded templates:** Virtual filesystem to embed the templates in the resulting binary for complete portability.

psetup uses an embedded filesystem containing template files for consistent content generation:
* **License templates:** Pre-configured MIT and Apache 2.0 licenses
* **README template:** Basic project documentation structure
* **Gitignore template:** Common ignore patterns for various environments

Templates are embedded directly into the binary using Go's `embed` package, eliminating external dependencies and ensuring the tool works anywhere without additional setup.

## Notes
* All configurations have sensible defaults for immediate usage
* Invalid arguments fall back to default values rather than causing errors
* The tool ensures all parent directories exist before creating files
* This project has a previous version made in Lua [project-basic-setup](https://github.com/leojimenezg/project-basic-setup), where the problem was very strictly solved.
* This new version was designed to be as flexible and modular as possible. Also, this was my first project made in Go, so I don't expect it to be perfect

## Useful Resources

* [Go Documentation](https://golang.org/doc/) - Official Go programming language documentation
* [Go embed Package](https://pkg.go.dev/embed) - Official Go embed package documentation
* [argparse Documentation](./argparse/README.md) - Detailed argparse package documentation
* [itemgen Documentation](./itemgen/README.md) - Detailed itemgen package documentation
