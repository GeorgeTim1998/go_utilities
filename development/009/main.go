package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// Функция для загрузки URL и сохранения содержимого в файл
func downloadURL(url, outputPath string) error {
	// Выполняем HTTP-запрос
	resp, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Создаем файл для записи содержимого
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer out.Close()

	// Копируем данные из ответа в файл
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка при записи данных в файл: %v", err)
	}

	return nil
}

// Главная функция для обработки входных данных и вызова загрузки
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: wget <URL>")
		return
	}

	url := os.Args[1]

	// Определяем имя файла для сохранения
	urlParts := strings.Split(url, "/")
	fileName := urlParts[len(urlParts)-1]
	if fileName == "" {
		fileName = "index.html"
	}

	// Создаем директорию для сохранения, если ее нет
	outputDir := "downloads"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Ошибка при создании директории: %v\n", err)
		return
	}

	outputPath := path.Join(outputDir, fileName)
	err = downloadURL(url, outputPath)
	if err != nil {
		fmt.Printf("Ошибка при загрузке URL: %v\n", err)
		return
	}

	fmt.Printf("Файл сохранен по адресу: %s\n", outputPath)
}
