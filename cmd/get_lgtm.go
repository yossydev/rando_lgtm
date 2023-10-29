package cmd

import (
    "fmt"
    "math/rand"
    "os"
    "strings"
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

var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Get a random LGTM message",
    Args:  cobra.NoArgs,
    Run: func(cmd *cobra.Command, args []string) {
        message, err := getRandomMessageFromFile()
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        fmt.Println(message)
    },
}

func surroundWithEmoji(s string) string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    emoji := emojis[r.Intn(len(emojis))]
    return emoji + s + emoji
}

var defaultMessages = []string{
    "Looks good to me!",
    "Great job!",
    "Well done!",
    "Keep up the good work!",
}

func getRandomMessageFromFile() (string, error) {
    var messages []string
    if _, err := os.Stat("messages.txt"); err == nil {
        content, err := os.ReadFile("messages.txt")
        if err != nil {
            return "", fmt.Errorf("failed to read file: %v", err)
        }
        lines := strings.Split(string(content), "\n")
        for _, line := range lines {
            if line != "" {
                messages = append(messages, line)
            }
        }
    }
    for _, message := range defaultMessages {
        if message != "" {
            messages = append(messages, message)
        }
    }
    if len(messages) == 0 {
        return "", fmt.Errorf("no messages available")
    }
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomMessage := messages[r.Intn(len(messages))]
    return surroundWithEmoji(randomMessage), nil
}

func init() {
    rootCmd.AddCommand(getCmd)
}

