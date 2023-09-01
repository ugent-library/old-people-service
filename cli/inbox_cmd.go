package cli

import "github.com/spf13/cobra"

var inboxCmd = &cobra.Command{
	Use: "inbox",
}

func init() {
	rootCmd.AddCommand(inboxCmd)
}
