# GoPhish CLI
This is a CLI client for the GoPhish API. It is still a work in progress with only things implemented that I have needed. I have plans to implement most of the API but if you want something to be implemented faster, create an issue and I will get to it.

## Install
To install the CLI, GoLang v1.19+ is required. Run the following to install:
```
go install github.com/NoF0rte/gophish-cli@latest
```

## Config
The special config options can be passed via the command line but can also be stored in a `gophish-cli.yaml` config file. The default config file looks like the following:
```
api-key: ""
url: https://localhost:3333
```

Running the following will create a `gophish-cli.yaml` file in the current directory
```
gophish-cli config -s
```

## Usage
```
$ gophish-cli --help

A CLI to interact with the Gophish API

Usage:
  gophish-cli [command]

Available Commands:
  campaigns   List, add, or delete campaigns
  completion  Generate the autocompletion script for the specified shell
  config      Display config information
  help        Help about any command
  login       Login to GoPhish and attempt to retrieve the user's API key
  profiles    List, add, or delete sending profiles
  templates   List, add, or delete e-mail templates

Flags:
  -T, --api-key string        A valid Gophish API key
  -h, --help                  help for gophish-cli
  -u, --url string            The URL to the Gophish server (default "https://localhost:3333")
  -V, --vars stringToString   Variables to use when creating/editing items from files that have replacement variables. Use name=value syntax. (default [])

Use "gophish-cli [command] --help" for more information about a command.

```
### Templates
Currently, the only thing the CLI can do with the API is get, add, and delete e-mail templates.

```
$ gophish-cli templates --help

List, add, or delete templates

Usage:
  gophish-cli templates [flags]
  gophish-cli templates [command]

Available Commands:
  add         Add new e-mail template(s)
  delete      Delete e-mail templates
  export      Export e-mail templates

Flags:
  -h, --help           help for templates
      --id int         Get the template by ID
  -n, --name string    Get the template by name.
  -r, --regex string   List the templates with the name matching the regex.
      --show-content   Show the template content in output

Global Flags:
  -T, --api-key string   A valid Gophish API key
  -u, --url string       The URL to the Gophish server (default "https://localhost:3333")

Use "gophish-cli templates [command] --help" for more information about a command.

```

#### Adding E-mail Templates
Use the example templates ([examples/templates](examples/templates)) for reference on the e-mail template file format.
