package option

import (
	"fmt"
	"os"
	"strings"
)

type Option struct {
	Name        string
	Description string
	Flags       []string
}

type App struct {
	NoHelp      bool
	Description string
	examples    []string
	options     []Option
}

func (self *App) help() {
	if len(self.Description) > 0 {
		fmt.Printf("[%s]: %s\n\n", os.Args[0], self.Description)
	}
	if self.options != nil && len(self.options) > 0 {
		fmt.Printf("Flags:\n")
		fmt.Printf("%-30s\t%s\n", "help, -h, --help", "display help information")
		for _, option := range self.options {
			fmt.Printf("%-30s\t%s\n", strings.Join(option.Flags, ", "), option.Description)
		}
	}
	if len(self.examples) > 0 {
		fmt.Printf("\nUsage:\n")
		for _, e := range self.examples {
			fmt.Printf("%s %s\n", os.Args[0], e)
		}
	}
	os.Exit(0)
}

func (self *App) Flag(name string, description string, flags ...string) {
	if len(flags) == 0 {
		return
	}
	if self.options == nil {
		self.options = make([]Option, 0)
	}
	o := Option{Name: name, Description: description, Flags: make([]string, 0)}
	for _, flag := range flags {
		if strings.HasPrefix(flag, "-") {
			o.Flags = append(o.Flags, flag)
		}
	}
	if len(o.Flags) > 0 {
		self.options = append(self.options, o)
	}
}

func (self *App) Example(example string) {
	if len(example) == 0 {
		return
	}
	if self.examples == nil {
		self.examples = make([]string, 0)
	}
	self.examples = append(self.examples, example)
}

func (self *App) Parse() map[string]interface{} {
	if !self.NoHelp {
		for _, v := range os.Args {
			if v == "help" || v == "--help" || v == "-h" {
				self.help()
			}
		}
	}
	options := make(map[string]interface{})
	for idx, arg := range os.Args {
		for _, option := range self.options {
			for _, flag := range option.Flags {
				if strings.HasPrefix(arg, flag) {
					if strings.HasPrefix(flag, "--") {
						parsed := strings.Split(arg, "=")
						if len(parsed) > 1 {
							options[option.Name] = parsed[1]
						} else {
							options[option.Name] = true
						}
					} else {
						if idx+1 < len(os.Args) && !strings.HasPrefix(os.Args[idx+1], "-") {
							options[option.Name] = os.Args[idx+1]
						} else {
							options[option.Name] = true
						}
					}
				}
			}
		}
	}
	return options
}
