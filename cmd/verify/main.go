package verify

import (
	"log/slog"
	"os"

	"git.a71.su/Andrew71/pye/auth"
)

func Verify(token, filename string) {
	key, err := os.ReadFile(filename)
	if err != nil {
		slog.Error("error reading file", "error", err, "file", filename)
	}
	t, err := auth.VerifyJWT(token, key)
	slog.Info("result", "token", t, "error", err, "ok", err == nil)
}
