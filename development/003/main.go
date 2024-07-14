package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// определяем допустимые при вызове программы флаги с их описанием
	k := flag.Int("k", 0, "column for sorting")
	n := flag.Bool("n", false, "sort by numeric value")
	r := flag.Bool("r", false, "reverse the result of comparisons")
	u := flag.Bool("u", false, "suppress duplicate lines")
	M := flag.Bool("M", false, "sort by month name")
	b := flag.Bool("b", false, "ignore leading blanks")
	c := flag.Bool("c", false, "check if sorted")
	h := flag.Bool("h", false, "sort by numeric value with suffixes")

	// парсим флаги. информация о них попадет в флаги определенные выше
	flag.Parse()

	// проверяем что передали input файл при вызове программы
	if flag.NArg() == 0 {
		fmt.Println("Usage: sort -k <column> -n -r -u -M -b -c -h <filename>")
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
	sortLines(lines, *k, *n, *r, *u, *M, *b, *h)

	// выводим отсортированные строки
	for _, line := range lines {
		fmt.Println(line)
	}

	// записываем отсортированные строки в файл
	if err := writeLines("done"+filename, lines); err != nil {
		fmt.Printf("Error writing sorted lines to file: %v\n", err)
		os.Exit(1)
	}
}

// readLines читает файл и возвращает строки  виде слайса []string
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

// sortLines sorts lines according to specified options
func sortLines(lines []string, k int, n, r, u, M, b, h bool) {
	sort.Slice(lines, func(i, j int) bool {
		// Implement sorting logic based on flags
		if M {
			// Sort by month name
			return monthIndex(strings.Fields(lines[i])[k]) < monthIndex(strings.Fields(lines[j])[k])
		} else if n {
			// Sort by numeric value
			num1, err1 := strconv.Atoi(getField(lines[i], k, b))
			num2, err2 := strconv.Atoi(getField(lines[j], k, b))

			if err1 == nil && err2 == nil {
				return num1 < num2
			}
			// Fallback to lexicographical order if conversion fails
			return lines[i] < lines[j]
		} else {
			// Default: lexicographical order
			return lines[i] < lines[j]
		}
		if r {
			return !r
		}
	})
}

// getField returns the k-th field from a line based on delimiter and ignoring leading spaces if specified
func getField(line string, k int, ignoreLeadingSpaces bool) string {
	fields := strings.Fields(line)
	if k >= len(fields) {
		return ""
	}
	return fields[k]
}

// monthIndex returns the index of a month name in time.Month
func monthIndex(monthName string) int {
	for i := 1; i <= 12; i++ {
		if monthName == time.Month(i).String() {
			return i
		}
	}
	return 0
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
