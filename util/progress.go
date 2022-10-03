package util

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
)

// ProgressBar returns the default styled progress bar
func ProgressBar(description string, size int) *progressbar.ProgressBar {
	return progressbar.NewOptions(size,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription(
			fmt.Sprintf("[cyan][%s][reset]", description),
		),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]█[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
}
