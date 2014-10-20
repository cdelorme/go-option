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

func (app *App) help() {
    if len(app.Description) > 0 {
    	fmt.Printf("[%s]: %s\n\n", os.Args[0], app.Description)
    }
    if app.options != nil && len(app.options) > 0 {
    	fmt.Printf("Flags:\n")
    	fmt.Printf("%-30s\t%s\n", "help, -h, --help", "display help information")
    	for _, option := range app.options {
    		fmt.Printf("%-30s\t%s\n", strings.Join(option.Flags, ", "), option.Description)
    	}
    }
    if len(app.examples) > 0 {
    	fmt.Printf("\nUsage:\n")
    	for _, e := range app.examples {
    		fmt.Printf("%s %s\n", os.Args[0], e)
    	}
    }
    os.Exit(0)
}

func (app *App) Flag(name string, description string, flags ...string) {
    if len(flags) == 0 {
    	return
    }
    if app.options == nil {
    	app.options = make([]Option, 0)
    }
    o := Option{Name: name, Description: description, Flags: make([]string, 0)}
    for _, flag := range flags {
    	if strings.HasPrefix(flag, "-") {
    		o.Flags = append(o.Flags, flag)
    	}
    }
    if len(o.Flags) > 0 {
    	app.options = append(app.options, o)
    }
}

func (app *App) Example(example string) {
    if len(example) == 0 {
    	return
    }
    if app.examples == nil {
    	app.examples = make([]string, 0)
    }
    app.examples = append(app.examples, example)
}

func (app *App) Parse() map[string]interface{} {
    if !app.NoHelp {
    	for _, v := range os.Args {
    		if v == "help" || v == "--help" || v == "-h" {
    			app.help()
    		}
    	}
    }
    options := make(map[string]interface{})
    for idx, arg := range os.Args {
    	for _, option := range app.options {
    		for _, flag := range option.Flags {
    			if strings.HasPrefix(arg, flag) {
    				if strings.HasPrefix(flag, "--") {
    					parsed := strings.Split(arg, "=")
    					if len(parsed) > 1 {
    						options[option.Name] = parsed[1]
    					}
    				} else {
    					if idx+1 < len(os.Args) {
    						options[option.Name] = os.Args[idx+1]
    					} else {
    						options[option.Name] = "true"
    					}
    				}
    			}
    		}
    	}
    }
    return options
}
