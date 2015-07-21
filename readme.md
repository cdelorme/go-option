
# go-option

A library to abstract cli-option parsing in a sane but posix-familiar way.


## alternatives

There is a supplied golang `flag` package that serves the same purpose but with its own somewhat unusual handling.  There is an extensive third party library as well that has a myriad of additional support features:

- [golang flag](http://golang.org/pkg/flag/)
- [kingpin](https://github.com/alecthomas/kingpin)


## sales pitch

My library aims to deliver the simplest POSIX compatible implementation.

It provides you with a container to process your applications expected arguments, and automation of the help command including example usage.  It supports multiple argument keys.

Invalid arguments will be silently ignored.  Required arguments must be enforced by your own implementation.

The results are provided to you in the form of a `map[string]interface{}`.  A [map tool](https://github.com/cdelorme/go-maps) is available for casting and merging assistance.

It does not:

- include unit tests
- use abstractions
- contain more than 110 lines of code

It's size makes it possible to effortlessly grasp the implementation, and have the confidence in using it.

This package can be combined with [go-config](https://github.com/cdelorme/go-config) and [go-env](https://github.com/cdelorme/go-env) to produce a single `map[string]interface{}` of application settings.


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

