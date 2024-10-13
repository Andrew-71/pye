package cmd

import (
	"fmt"

	"git.a71.su/Andrew71/pye/storage"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findUserCmd)
}

var findUserCmd = &cobra.Command{
	Use:   "find <uuid/email> <query>",
	Short: "Find a user",
	Long:  `Find information about a user from their UUID or email`,
	Args:  cobra.ExactArgs(2),
	Run:   findUser,
}

// TODO: Better name.
func findUser(cmd *cobra.Command, args []string) {
	var user storage.User
	var ok bool
	if args[0] == "email" {
		user, ok = storage.Data.ByEmail(args[1])
	} else if args[0] == "uuid" {
		user, ok = storage.Data.ById(args[1])
	} else {
		fmt.Println("expected email or uuid")
		return
	}
	if !ok {
		fmt.Println("User not found")
	} else {
		fmt.Printf("Information for user:\nuuid\t- %s\nemail\t- %s\nhash\t- %s\n",
			user.Uuid, user.Email, user.Hash)
	}
}
