package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:     "host",
	Aliases: []string{"h"},
	Short:   "Host intruder",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Host intruder")
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)
}
