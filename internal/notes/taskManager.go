package notes

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TaskManager(command string) error {
	reader := bufio.NewReader(os.Stdin)
	switch command {

	case "add":
		var title, content, filename string
		fmt.Print("Введите название: ")
		title, _ = reader.ReadString('\n')
		title = strings.TrimSpace(title)
		fmt.Println("Введите название файла: ")
		filename, _ = reader.ReadString('\n')
		filename = strings.TrimSpace(filename)
		fmt.Println("Введите содержание: ")
		content, _ = reader.ReadString('\n')
		content = strings.TrimSpace(content)
		path := MakePath(filename)
		_, err := os.Stat(path)

		if os.IsNotExist(err) {

			CreateFileJson(title, content, filename)
			if err != nil {
				return err
			}

		} else {

			AppendFile(title, content, filename)

		}

	case "list":
		arrList := ListFiles()
		for i := 0; i < len(arrList); i++ {
			fmt.Println(arrList[i])
		}

	case "remove":
		var title string
		fmt.Print("Введите название: ")
		title, _ = reader.ReadString('\n')
		title = strings.TrimSpace(title)
		RemoveFile(title)

	case "read":
		var title string
		fmt.Print("Введите название: ")
		title, _ = reader.ReadString('\n')
		title = strings.TrimSpace(title)
		ReadFile(title)
	}

	return nil
}
