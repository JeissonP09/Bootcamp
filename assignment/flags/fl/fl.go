package fl

import "os"

type Flag struct {
	value       bool   // define type
	description string // define type
}

var flags = map[string]*Flag{} // define value to control here all program flags

func Parse() {
	args := os.Args[1:]
	// Evaluate if the defined flags should change their default value
	for _, arg := range args {
		if flag, exits := flags[arg]; exits {
			flag.value = true
		}
	}
}

func Bool(cmd string, value bool, description string) *bool {
	// Add logic to create boolean flags
	flag := &Flag{
		value:       value,
		description: description,
	}
	flags[cmd] = flag
	return &flag.value
}
