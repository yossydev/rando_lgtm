package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "lgtmapp",
    Short: "LGTM message generator",
    Long:  `This is a CLI tool for generating and managing LGTM messages.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    Execute()
}
