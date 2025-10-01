package cmds

import (
	"fmt"
	"os"
)

func commandExit(c *Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
