# go-wordcloud

A CLI tool to generate word clouds using the Cobra and Viper packages.

## Installation

To install the application, clone the repository and build the project:

```bash
git clone https://github.com/jbrinkman/go-wordcloud.git
cd go-wordcloud
go build
```

## Usage

After building the project, you can run the CLI tool with the following command:

```bash
./go-wordcloud
```

### Commands

- `wordcloud`: The main command to generate word clouds.

## Configuration

The application supports configuration through a configuration file or environment variables. You can specify settings such as input files, output formats, and other options.

### Configuration File

Create a configuration file named `config.yaml` in the root directory of the project. Here is an example configuration:

```yaml
input: "path/to/input.txt"
output: "path/to/output.png"
```

## Contributing

Feel free to submit issues or pull requests to improve the project. 

## License

This project is licensed under the MIT License. See the LICENSE file for more details.