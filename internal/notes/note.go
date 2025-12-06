// 1. Занести весь проект на гит!

package notes // Тут пишем пакет используемый

import (
	"encoding/json" // Отвечает за обработку JSON файлов
	"fmt"           // Отвечает за форматирование
	"os"            // Системные функции, запись файлов и тд
)

type Note struct { // Создаем класс заметок
	ID      int    `json:"id"`      // Прописываем поле айди, прописываем имя в JSON
	Title   string `json:"title"`   // Прописываем поле название, прописываем имя в JSON
	Content string `json:"content"` // Прописываем поле содержимого, прописываем имя в JSON
}

func NewID(arrNotes []Note) int {
	return len(arrNotes)
}

func NotesLogicManager(title, content string, filename string) error {
	// 2. Надо бы вынести обработчик в отдельный файл

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		/*
		   3. Доделать реализацию создания файла в обработчике логики
		*/
	} else {
		/*
		   4. Доделать реализацию дополнения файла в обработчике логики
		*/
	}
	return nil
}

func LoadFile(filename string) []Note {
	notes := []Note{}
	jsonData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error with read file: ", err)
		return []Note{}
	}
	json.Unmarshal(jsonData, &notes)
	return notes
}

func CreateFileJson(title, content string, filename string) error {
	var arrNotes []Note
	id := NewID(arrNotes)
	data := []Note{Note{id, title, content}}

	fileJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON: ", err)
		return err
	}

	file, err := os.Create(filename) // Придумать, как сохранять файл в медиа директорию
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
	data := LoadFile(fileName) // type: []Note
	id := NewID(data)
	newData := Note{id, title, content}
	data = append(data, newData) // Добавление элемента в конец массива
	fileJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при парсе JSON: ", err)
		return err
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка перезаписи файла:", err)
	}
	defer file.Close()

	_, err = file.Write(fileJson)

	return nil
}
