package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new LGTM message",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        err := addMessageToFile(args[0])
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        fmt.Printf("Added new LGTM message: '%s'\n", args[0])
    },
}

func addMessageToFile(message string) error {
    file, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()
    _, err = file.WriteString(message + "\n")
    if err != nil {
        return fmt.Errorf("failed to write to file: %v", err)
    }
    return nil
}

func init() {
    rootCmd.AddCommand(addCmd)
}
