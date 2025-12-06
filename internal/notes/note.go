package notes // Тут пишем пакет используемый

import (
	"encoding/json" // Отвечает за обработку JSON файлов
	"fmt"           // Отвечает за форматирование
	"os"            // Системные функции, запись файлов и тд
	"path/filepath"
)

type Note struct { // Создаем класс заметок
	ID      int    `json:"id"`      // Прописываем поле айди, прописываем имя в JSON
	Title   string `json:"title"`   // Прописываем поле название, прописываем имя в JSON
	Content string `json:"content"` // Прописываем поле содержимого, прописываем имя в JSON
}

func NewID(arrNotes []Note) int {
	return len(arrNotes)
}

func LoadFile(filename string) []Note {
	notes := []Note{}
	absolutePath := MakePath(filename)
	jsonData, err := os.ReadFile(absolutePath)
	if err != nil {
		fmt.Println("Error with read file: ", err)
		return []Note{}
	}
	json.Unmarshal(jsonData, &notes)
	return notes
}

func MakePath(fileName string) string {
	absolutePath := filepath.Join("./media", fileName+".json")
	return absolutePath
}

func CreateFileJson(title, content string, filename string) error {
	var arrNotes []Note
	id := NewID(arrNotes)
	data := []Note{Note{id, title, content}}

	absolutePath := MakePath(filename)

	fileJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON: ", err)
		return err
	}

	file, err := os.Create(absolutePath) // Придумать, как сохранять файл в медиа директорию
	if err != nil {
		fmt.Println("Error create file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(fileJson)
	if err != nil {
		fmt.Println("Error write file:", err)
		return err
	}
	return nil
}

func AppendFile(title, content string, fileName string) error {
	absolutePath := MakePath(fileName)
	data := LoadFile(fileName) // type: []Note
	id := NewID(data)
	newData := Note{id, title, content}
	data = append(data, newData) // Добавление элемента в конец массива
	fileJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при парсе JSON: ", err)
		return err
	}
	file, err := os.Create(absolutePath)
	if err != nil {
		fmt.Println("Ошибка перезаписи файла:", err)
	}
	defer file.Close()

	_, err = file.Write(fileJson)

	return nil
}

func ListFiles() []string {
	dir, err := os.Open("./media")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer dir.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return files
}

func RemoveFile(fileName string) error {
	files := ListFiles()
	for _, file := range files {
		if fileName+".json" == file {
			path := MakePath(fileName)
			fmt.Println(path)
			os.Remove(path)
		}
	}
	return nil
}

func ReadFile(fileName string) string {
	files := ListFiles()
	for _, file := range files {
		if fileName+".json" == file {
			data := LoadFile(fileName)
			fmt.Println(data)
		}
	}
	return ""
}
