package app

import (
	"fmt"

	"git.a71.su/Andrew71/pye/internal/models/user"
	"git.a71.su/Andrew71/pye/internal/storage"
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

func findUser(cmd *cobra.Command, args []string) {
	var user user.User
	var ok bool
	if args[0] == "email" {
		user, ok = storage.Data.ByEmail(args[1])
	} else if args[0] == "uuid" {
		user, ok = storage.Data.ById(args[1])
	} else {
		fmt.Println("Expected email or uuid")
		return
	}
	if !ok {
		fmt.Println("User not found")
	} else {
		fmt.Printf("Information for user:\nuuid\t- %s\nemail\t- %s\nhash\t- %s\n",
			user.Uuid, user.Email, user.Hash)
	}
}
