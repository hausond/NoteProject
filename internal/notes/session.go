package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Session() {
	var command string
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите комманду: ")
		command, _ = reader.ReadString('\n')
		command = strings.TrimSpace(command)
		TaskManager(command)
		if sessionOver(command) {
			break
		}
	}
}

func sessionOver(command string) bool {
	return command == "exit"
}
