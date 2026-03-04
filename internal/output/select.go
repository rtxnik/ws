package output

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

// SelectOption represents a selectable item with a display label and value.
type SelectOption struct {
	Label string
	Value string
}

// Select displays an interactive selector and returns the chosen value.
// Exits the process if no options are provided or the user cancels.
func Select(title string, options []SelectOption) string {
	if len(options) == 0 {
		Die("no items to select")
	}

	huhOpts := make([]huh.Option[string], 0, len(options))
	for _, o := range options {
		huhOpts = append(huhOpts, huh.NewOption(o.Label, o.Value))
	}

	var selected string
	err := huh.NewSelect[string]().
		Title(title).
		Options(huhOpts...).
		Value(&selected).
		Run()

	if err != nil {
		fmt.Fprintln(os.Stderr)
		os.Exit(0)
	}
	return selected
}

// StatusLabel formats a workspace status for display in the selector.
func StatusLabel(name, status string) string {
	return fmt.Sprintf("%s  %s", name, StatusText(status))
}
