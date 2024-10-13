package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"git.a71.su/Andrew71/pye/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

var (
	verifyToken string
	verifyFile  string
)

func init() {
	verifyCmd.Flags().StringVarP(&verifyToken, "token", "t", "", "token to verify")
	verifyCmd.MarkFlagRequired("token")
	verifyCmd.Flags().StringVarP(&verifyFile, "file", "f", "", "PEM file to use")
	rootCmd.AddCommand(verifyCmd)
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a JWT token",
	Long: `Pass a JWT token (and optionally a path to a PEM-formatted file with the public key)
		to verify whether it is valid.`,
	Run: verifyFunc,
}

// TODO: Needs a better name?
func verifyFunc(cmd *cobra.Command, args []string) {
	if verifyToken == "" {
		fmt.Println("Empty token supplied!")
		return
	}

	var t *jwt.Token
	var err error
	if verifyFile == "" {
		fmt.Println("No PEM file supplied, assuming local")
		t, err = auth.VerifyLocalJWT(verifyToken)
	} else {
		key, err_k := os.ReadFile(verifyFile)
		if err_k != nil {
			slog.Error("error reading file", "error", err, "file", verifyFile)
			return
		}
		t, err = auth.VerifyJWT(verifyToken, key)
	}
	slog.Debug("result", "token", t, "error", err, "ok", err == nil)
	if err == nil {
		fmt.Println("Token valid!")
	} else {
		fmt.Println("Token invalid!", err)
	}
}
