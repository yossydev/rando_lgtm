package cmd

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/spf13/cobra"
)

var emojis = []string{
    "😀", "😃", "😄", "😁", "😆", "😅", "😂", "🤣", "😊", "😇",
    "🙂", "🙃", "😉", "😌", "😍", "🥰", "😘", "😗", "😙", "😚",
    "😋", "😛", "😝", "😜", "🤪", "🤨", "🧐", "🤓", "😎", "🥸",
    "😏", "😒", "😫", "😩", "🥺", "😢", "😭", "😤", "🤯", "😳",
    "🤗", "🤭", "🤫", "😬", "😴", "🤤", "🥴", "🤑", "🤠", "😈",
    "👹", "👺", "🤡", "👻", "💀", "👽", "👾", "🤖", "🎃", "😺",
    "😸", "😹", "😻", "😼", "😽",
}

var getLgtmCmd = &cobra.Command{
    Use:   "getLgtm",
    Short: "Get a random LGTM message",
    Long:  `This command allows you to get a random LGTM message from the list of messages.`,
    Run: func(cmd *cobra.Command, args []string) {
        rand.Seed(time.Now().UnixNano())
        selectedEmoji := emojis[rand.Intn(len(emojis))]
        fmt.Printf("%s %s %s\n", selectedEmoji, lgtms[rand.Intn(len(lgtms))], selectedEmoji)
    },
}

func init() {
    rootCmd.AddCommand(getLgtmCmd)
}
