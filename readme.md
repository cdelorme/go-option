
# go-option

A library to abstract cli-option parsing in a sane but posix-familiar way.


## alternatives

There is a supplied golang `flag` package that serves the same purpose but with its own somewhat unusual handling.  There is an extensive third party library as well that has a myriad of additional support features:

- [golang flag](http://golang.org/pkg/flag/)
- [kingpin](https://github.com/alecthomas/kingpin)


## sales pitch

My library aims to provide the simplest implementation, while at the same time providing POSIX compatible and familiar argument parser handling.

It's features include:

- single container to identify your application and its expected arguments
- simple argument registration
- parse method to initiate parsing
- optional automated help handler with example and description support
- traditional parsing with single and double dashed arguments (ex. `--key=value || -k value`)

It does not:

- have a comprehensive testing suite
- contain more than 100 lines of code
- provide wildly extensible abstractions

There are some minor differences in how it handles arguments from standard posix implementations.  It will accept more than one character with single-dash flags.

The code will silently ignore supplied arguments that are invalid.

Finally, the parse returns a `map[string]string` with named keys, which the developer is responsible for casting to needed types.


## usage

To create a new app instance with a description:

    appFlags := option.App{Description: "optional app description for help output"}

Alternatively you can specify not to use the builtin help via:

    appFlags := option.App{NoHelp: true}

Or:

    appFlags.NoHelp = true

You can register flags with an internal name for your map, a description for help output, and one or more named flags (**flags must include dash prefixes**):

    appFlags.Flag("name", "description", "--name", "-n")

_If you do not supply dash prefixes to flags they will be ignored._

You can provide example usage to help output with:

    appFlags.Example("-name value")

_Examples will automatically be prefixes with the application name dynamically._

Finally, you can parse registered flags:

    flags := appFlags.Parse()
