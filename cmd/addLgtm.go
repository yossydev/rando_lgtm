package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var lgtms = []string{
    "Looks good to me!",
    "Perfect!",
    "Great work!",
    "Nice job!",
    "Keep it up!",
}

var addLgtmCmd = &cobra.Command{
    Use:   "addLgtm",
    Short: "Add a new LGTM message",
    Long:  `This command allows you to add a new LGTM message to the list of messages.`,
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        lgtms = append(lgtms, args[0])
        fmt.Printf("Added new LGTM message: '%s'\n", args[0])
    },
}

func init() {
 rootCmd.AddCommand(addLgtmCmd)
}
    
