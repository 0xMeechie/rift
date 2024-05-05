/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/fdaygon/rift/pkg/spotify"
	"github.com/fdaygon/rift/pkg/terminal"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into spotify",
	Long: `Login into spotify. This will provide rift with the needed access token to make
	api request on your behalf. This will need to be done before starting to ensure that
	a recent access token is being used and it is not expired.`,
	Run: func(cmd *cobra.Command, args []string) {
		spotifyAuthUrl := spotify.UserAuth()
		if err := exec.Command("open", spotifyAuthUrl).Run(); err != nil {
			fmt.Println("Unable to open auth page")
			os.Exit(1)
		}

		for spotify.AuthCode == "" {
			time.Sleep(time.Second * 5)
			fmt.Println("waiting for user to log in")
		}

		spotify.GetToken()

		shell := terminal.CurrentShell()
		terminal.CheckProfileFile(shell, spotify.Token.Token)

		fmt.Println("Successfully Logged In")
		fmt.Println("Closing now")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
