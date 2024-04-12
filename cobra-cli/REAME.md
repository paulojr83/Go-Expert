## Usage:
_cobra-cli [command]_

### _Available Commands:_
* add         Add a command to a Cobra Application
* completion  Generate the autocompletion script for the specified shell
* help        Help about any command
* init        Initialize a Cobra Application

### _Flags:_
* -a, --author string    author name for copyright attribution (default "YOUR NAME")
* --config string    config file (default is $HOME/.cobra.yaml)
* -h, --help             help for cobra-cli
* -l, --license string   name of license for the project
* --viper            use Viper for configuration

<code>Use "cobra-cli [command] --help" for more information about a command.</code>

## Adicionando comandos
 
    $ cobra-cli add category
    $ cobra-cli add create -p 'categoryCmd' 
