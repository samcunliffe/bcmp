package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tidyCmd = &cobra.Command{
	Use:     "tidy",
	Short:   "Not implemented :(",
	Aliases: []string{"t"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("not implemented")
	},
}

func init() {
	rootCmd.AddCommand(tidyCmd)
}
