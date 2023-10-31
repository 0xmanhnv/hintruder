package cmd

import (
	"context"
	"fmt"
	"hintruder/pkg/hintruder"
	"os"

	"github.com/spf13/cobra"
)

var proxy string
var fileRequest string
var rootCmd = &cobra.Command{
	Use:   "hintruder",
	Short: "hintruder - Like burpsuite intruder, but it's cli",
	Long: `hintruder is host intruder
	Or, hand intruder`,
	Run: handlerRun,
}

func init() {
	rootCmd.Flags().StringVar(&proxy, "proxy", "", "Add proxy")
	rootCmd.Flags().StringVarP(&fileRequest, "file", "f", "Request file", "")

	// Make requied
	rootCmd.MarkFlagRequired("file")
}

func handlerRun(cmd *cobra.Command, args []string) {
	h := hintruder.Hintruder{}
	if proxy != "" {
		h.ProxyUrl = proxy
	}

	h.Protocol = "https"
	h.TlsVerify = true

	ctx := context.Background()
	h.Run(ctx, fileRequest)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
