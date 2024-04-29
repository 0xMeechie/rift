package terminal

import (
	"fmt"
	"os/exec"
)

// This is for checking to see if there is a Entry in the zshrc
func CheckZSH() bool {
	command := "cat ~/.zshrc | grep Spotify_Token"
	cmd := exec.Command("bash", "-c", command)

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
		return false
	}
	// otherwise, print the output from running the command

	if len(string(out)) != 0 {
		fmt.Println("True")
		fmt.Println(string(out))
		return true
	}

	return false

}

func CheckBash() {

}

func CheckKitty() {
}
