package commands

import (
	"os"
)

type ExitCommand struct {
	Name string
	Description string
	Callback    func() error
}


func CommandExit() error {
	os.Exit(0)
	return nil
}



