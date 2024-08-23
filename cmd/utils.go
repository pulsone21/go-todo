package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var w *tabwriter.Writer

func NewTabWriter(c Config) {
	w = tabwriter.NewWriter(os.Stdout, c.minWidth, c.tabWidth, c.padding, c.padChar, tabwriter.TabIndent)
}

func Print(msg string) {
	fmt.Fprintln(w, msg)
}
