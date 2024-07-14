package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// определяем допустимые при вызове программы флаги с их описанием
	k := flag.Int("k", 0, "column for sorting (default is entire line)")
	n := flag.Bool("n", false, "sort by numeric value")
	r := flag.Bool("r", false, "reverse the result of comparisons")
	u := flag.Bool("u", false, "suppress duplicate lines")

	// парсим флаги. информация о них попадет в флаги определенные выше
	flag.Parse()

	// проверяем что передали input файл при вызове программы
	if flag.NArg() == 0 {
		fmt.Println("Usage: sort -k <column> -n -r -u <filename>")
		os.Exit(1)
	}
	filename := flag.Args()[0]

	// считываем строчки из переданного файла
	lines, err := readLines(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// производим сортировку
	sortLines(lines, *k, *n, *r)

	// убираем дубликаты если -u передан
	if *u {
		lines = removeDuplicates(lines)
	}

	// выводим отсортированные строки
	for _, line := range lines {
		fmt.Println(line)
	}

	// записываем отсортированные строки в файл
	if err := writeLines(filename, lines); err != nil {
		fmt.Printf("Error writing sorted lines to file: %v\n", err)
		os.Exit(1)
	}
}

// readLines читает файл и возвращает строки виде слайса []string
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	// скинируем файл построчно
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// проверяем наличие ошибок
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// бирает дубликаты строк с использоавнием карты
func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			result = append(result, line)
		}
	}

	return result
}

// sortLines сортирует строки с учетом переданных флагов
func sortLines(lines []string, k int, n, r bool) {
	sort.Slice(lines, func(i, j int) bool {
		// Split lines into fields if column k is specified
		if k > 0 {
			fields1 := strings.Fields(lines[i]) // разбиваем строки на слова, используя пробелы как разделители
			fields2 := strings.Fields(lines[j])

			if len(fields1) >= k && len(fields2) >= k {
				line1 := fields1[k-1]
				line2 := fields2[k-1]

				// Compare numeric values if -n flag is set
				if n {
					num1, err1 := strconv.Atoi(line1)
					num2, err2 := strconv.Atoi(line2)
					if err1 == nil && err2 == nil {
						line1 = strconv.Itoa(num1)
						line2 = strconv.Itoa(num2)
					}
				}

				// Compare in reverse order if -r flag is set
				if r {
					return line1 > line2
				}
				return line1 < line2
			}
		}

		// Default: lexicographical order
		// Compare entire lines if no column is specified
		if r {
			return lines[i] > lines[j]
		}
		return lines[i] < lines[j]
	})
}

// writeLines записывает итоговые данные в файл
func writeLines(filename string, lines []string) error {
	// создаем новый файл для записи в него
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//записываем данные в файл и проверяем наличие ошибок
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
