package cmd

import (
	"os"
	"text/tabwriter"
)

var w *tabwriter.Writer

func NewTabWriter(c Config) {
	b := []byte(c.PadChar)
	w = tabwriter.NewWriter(os.Stdout, c.MinWidth, c.TabWidth, c.Padding, b[0], tabwriter.TabIndent)
}
