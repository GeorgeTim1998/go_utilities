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
	sortLines(lines, *k, *n, *r, *u, *M, *b, *c, *h)

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
func sortLines(lines []string, k int, n, r, u, M, b, _, h bool) {
	sort.Slice(lines, func(i, j int) bool {
		// Extract the key for the specified column and process flags
		keyI := getField(lines[i], k, b)
		keyJ := getField(lines[j], k, b)

		var less bool
		if M {
			// Sort by month name
			less = monthIndex(keyI) < monthIndex(keyJ)
		} else if n {
			// Sort by numeric value
			num1, err1 := strconv.Atoi(keyI)
			num2, err2 := strconv.Atoi(keyJ)

			if err1 == nil && err2 == nil {
				less = num1 < num2
			} else {
				// Fallback to lexicographical order if conversion fails
				less = keyI < keyJ
			}
		} else if h {
			// Sort by numeric value with suffixes
			num1, _ := parseHumanReadableNumber(keyI)
			num2, _ := parseHumanReadableNumber(keyJ)
			less = num1 < num2
		} else {
			// Default: lexicographical order
			less = keyI < keyJ
		}

		if r {
			return !less
		}
		return less
	})

	// Remove duplicates if the -u flag is set
	if u {
		lines = unique(lines)
	}
}

// getField extracts the field for the specified column, ignoring leading/trailing spaces if -b is set
func getField(line string, k int, b bool) string {
	fields := strings.Fields(line)
	if k < 0 || k >= len(fields) {
		return ""
	}
	field := fields[k]
	if b {
		field = strings.TrimSpace(field)
	}
	return field
}

// monthIndex converts a month name to an index (0 for January, 11 for December)
func monthIndex(month string) int {
	months := map[string]int{
		"January": 0, "February": 1, "March": 2, "April": 3,
		"May": 4, "June": 5, "July": 6, "August": 7,
		"September": 8, "October": 9, "November": 10, "December": 11,
	}
	return months[month]
}

// parseHumanReadableNumber converts a human-readable number with suffixes to an integer
func parseHumanReadableNumber(s string) (int64, error) {
	multipliers := map[byte]int64{
		'K': 1 << 10,
		'M': 1 << 20,
		'G': 1 << 30,
		'T': 1 << 40,
	}
	if len(s) == 0 {
		return 0, fmt.Errorf("invalid number")
	}
	n := len(s)
	lastChar := s[n-1]
	multiplier, hasSuffix := multipliers[lastChar]
	if hasSuffix {
		value, err := strconv.ParseInt(s[:n-1], 10, 64)
		if err != nil {
			return 0, err
		}
		return value * multiplier, nil
	}
	return strconv.ParseInt(s, 10, 64)
}

// unique removes duplicate lines
func unique(lines []string) []string {
	uniqueLines := []string{}
	lineMap := map[string]struct{}{}
	for _, line := range lines {
		if _, exists := lineMap[line]; !exists {
			lineMap[line] = struct{}{}
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
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
