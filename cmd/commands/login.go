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
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		fmt.Println("Successfully Logged In")
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
