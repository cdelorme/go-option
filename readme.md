
# go-option

A library to abstract cli-option parsing in a sane but posix-familiar way.


## alternatives

There is a builtin [`flag` package](https://golang.org/pkg/flag/), but it does not offer posix-standard implementation, making it somewhat unwieldy to begin with.  However, what really pushed me away was the verbosity when combined with configuration file and environment variable parsing to collect settings.

Another well maintained third-party option is [kingpin](https://github.com/alecthomas/kingpin), and there are plenty more.  The third-party tools tend to offer extended features and end up being rather sizable.


## sales pitch

My library aims to deliver the simplest POSIX compatible implementation.  _If you want a light-weight solution with no room to get lost in confusing errors or misbehavior of complex features then this is the library you are looking for._

It provides a struct that gives you:

- named option registration with one or more accepted flags
- example registration to support help output
- parser to acquire a standard `map[string]interface{}` with results

_It offers help functionality that can easily be disabled, and will automatically print the available options then exit with a code of 0._

**It comes fully tested demonstrating parsing using my [go-maps support library](https://github.com/cdelorme/go-maps), and the implementation is under 115 lines of code.**  It does **not** have any complicated functionality, or confusing abstraction layers.

This package, combined with [go-config](https://github.com/cdelorme/go-config) and [go-env](https://github.com/cdelorme/go-env) and leveraging [go-maps](https://github.com/cdelorme/go-maps), can merge three basic configuration methods into a single `map[string]interface{}`, and optionally cast into an applications custom configuration struct.


## usage

To create a new app instance with a description:

    appFlags := option.App{Description: "optional app description for help output"}

_You can disable help by setting `NoHelp: false` on instantiation, and thus omit the description._

Flags can be registered with a name, description, and one or more accepted **dash-prefixed inputs**:

    appFlags.Flag("name", "description", "--name", "-n")

If you plan to leverage the `help` system, you can also register example strings showing just the flags (_the executable name is automatically prefixed_):

    appFlags.Example("-name <value>")

Finally, you can parse registered flags:

    flags := appFlags.Parse()
