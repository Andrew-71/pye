package find_user

import (
	"fmt"

	"git.a71.su/Andrew71/pye/config"
	"git.a71.su/Andrew71/pye/storage"
	"git.a71.su/Andrew71/pye/storage/sqlite"
)

func FindUser(mode, query string) {
	data := sqlite.MustLoadSQLite(config.Cfg.SQLiteFile)
	var user storage.User
	var ok bool
	if mode == "email" {
		user, ok = data.ByEmail(query)
	} else if mode == "uuid" {
		user, ok = data.ById(query)
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
