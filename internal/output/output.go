package output

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

var (
	SectionStyle = StyleHeader

	successStyle = StyleSuccess
	errorStyle   = StyleError
	warnStyle    = StyleWarning
	infoStyle    = StyleInfo
	detailStyle  = StyleDim
)

func Info(msg string) {
	fmt.Fprintln(os.Stderr, infoStyle.Render("ℹ "+msg))
}

func Success(msg string) {
	fmt.Fprintln(os.Stderr, successStyle.Render("✓ "+msg))
}

func Warn(msg string) {
	fmt.Fprintln(os.Stderr, warnStyle.Render("⚠ "+msg))
}

func Detail(msg string) {
	fmt.Fprintln(os.Stderr, detailStyle.Render("  "+msg))
}

func Die(msg string) {
	fmt.Fprintln(os.Stderr, errorStyle.Render("✗ "+msg))
	os.Exit(1)
}

// Confirm shows an interactive confirmation dialog. Returns true only if
// the user explicitly confirms. Default is No (safe default).
func Confirm(title string, description string) bool {
	var confirmed bool
	err := huh.NewConfirm().
		Title("⚠ " + title).
		Description(description).
		Affirmative("Yes").
		Negative("No").
		Value(&confirmed).
		Run()

	if err != nil {
		return false
	}
	return confirmed
}
