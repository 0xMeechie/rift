package terminal

import (
	"fmt"
	"os/exec"
)

func ReplaceToken(oldToken, newToken string) {
	command := fmt.Sprintf(`sed 's/%s/%s/' ~/.zshrc > ~/.zshrc_tmp && mv ~/.zshrc_tmp ~/.zshrc
`, oldToken, newToken)
	cmd := exec.Command("bash", "-c", command)

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
		return
	}
	// otherwise, print the output from running the command

	fmt.Println(string(out))
}

func AddToken() {

}
