package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RefreshSource() {

	shell := CurrentShell()

	//command := fmt.Sprintf("source ~/.%src", shell)
	cmd := exec.Command("bash", "-c", shell)

	sr, err := cmd.Output()
	if err != nil {
		fmt.Println("Couldn't refresh shell")
		return
	}

	fmt.Println(string(sr))

}

func ReplaceToken(oldToken, newToken, shell string) {
	newEntry := fmt.Sprintf("export Spotify_Token=%s", newToken)
	command := fmt.Sprintf("sed 's/%s/%s/' ~/.%src > ~/.%src_tmp && mv ~/.%src_tmp ~/.%src", oldToken, newEntry, shell, shell, shell, shell)
	newCommand := strings.ReplaceAll(command, "\n", "")
	cmd := exec.Command("bash", "-c", newCommand)

	// The `Output` method executes the command and
	// collects the output, returning its value
	_, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not replace Token ")
		return
	}

}

func AddToken(token, shell string) {

	fileEntry := fmt.Sprintf(`echo export Spotify_Token='%s >> ~/.%src'`, token, shell)

	cmd := exec.Command("bash", "-c", fileEntry)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error Adding Token")
		return
	}

	fmt.Println("Successfully Added Token")

}

func CurrentShell() string {
	// getting the current shell being used.
	shell := os.Getenv("SHELL")
	// this trims the /bin from the string
	shelltrimmed := strings.TrimPrefix(shell, "/bin/")
	return shelltrimmed
}

// This function is to check the shell config files to see if the required env var are in it.
func CheckProfileFile(shell, token string) {
	switch shell {
	case "zsh":
		oldToken := CheckZSH()
		if oldToken != "" {
			ReplaceToken(oldToken, token, shell)
		} else {
			AddToken(token, shell)
		}

	case "bash":
		oldToken := CheckBash()
		if oldToken != "" {
			ReplaceToken(oldToken, token, shell)
		} else {
			AddToken(token, shell)
		}
	}
}

func CheckZSH() string {
	//Check to see if the spotify token is in the zshrc file
	command := "cat ~/.zshrc | grep Spotify_Token"
	cmd := exec.Command("bash", "-c", command)

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not find token in zshrc file ")
		return ""
	}
	// otherwise, print the output from running the command

	if string(out) != "" {
		return string(out)
	}

	return ""

}
func CheckBash() string {
	//Check to see if the spotify token is in the bashrc file
	command := "cat ~/.bashrc | grep Spotify_Token"
	cmd := exec.Command("bash", "-c", command)

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not find token in bashrc file")
		return ""
	}
	// otherwise, print the output from running the command

	if len(string(out)) != 0 {
		return string(out)
	}

	return ""

}

// This is needed to refresh the profile to get the newest profile
func RefreshProfile() {

}

func CheckPS() bool {
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
