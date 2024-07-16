package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Определяем флаги для командной строки
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Читаем данные из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверяем наличие разделителя. Если его нет - сразу выкидываем строку из рассмотрения
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		// Разбиваем строку по разделителю
		columns := strings.Split(line, *delimiter)
		selectedFields := selectFields(columns, *fields)

		// Выводим выбранные поля
		if len(selectedFields) > 0 {
			fmt.Println(strings.Join(selectedFields, *delimiter))
		}
	}

	// Обрабатываем возможны ошибки.
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
	}
}

// Функция для выбора полей на основе флага -f. Функция вернес слайс индексов колонок для cut
func selectFields(columns []string, fields string) []string {
	if fields == "" {
		return columns
	}

	// Разбиваем список полей на массив индексов
	fieldIndexes := strings.Split(fields, ",")
	selectedFields := make([]string, 0, len(fieldIndexes))

	for _, indexStr := range fieldIndexes {
		// Пытаемся привести считанное значение номера колонки к числу
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 1 || index > len(columns) {
			// Если значение невалидно, то пропускаем это вереданное значение
			continue
		}

		// выбираем содержимое колонки
		selectedFields = append(selectedFields, columns[index-1])
	}

	return selectedFields
}
