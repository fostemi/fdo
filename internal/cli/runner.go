package cli

import (
  "github.com/spf13/cobra"
)

var (
  rootCmd = &cobra.Command{
    Use: "fdo-cli",
    Short: "A cli tool for requesting opinionated Infrastructure",
    Long: `Foster DevOps is a CLI and API for Google Cloud to empower application developers to create specific cloud resources.`,
  }
)

func Execute() error {
  return rootCmd.Execute()
}

