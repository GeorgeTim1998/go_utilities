package main

import (
	"fmt"
	"sort"
	"strings"
)

// findAnagrams находит все множества анаграмм в заданном словаре
func findAnagrams(words []string) map[string][]string {
	anagramMap := make(map[string][]string)
	seen := make(map[string]bool)

	for _, word := range words {
		// приводим слова к нижнему регистру
		lowerWord := strings.ToLower(word)
		// сортируем слово, чтобы абстрагировать от  реального порядки и получить "суть" слова - все его буква в алфавитном порядке
		sortedWord := sortString(lowerWord)

		// если слово уже было, то переходим к следующей итерации
		if _, found := seen[lowerWord]; found {
			continue
		}

		// если слово не было найдено, то добавляем его к списку найденных анаграм под ключ sortedWord
		anagramMap[sortedWord] = append(anagramMap[sortedWord], lowerWord)
		// помечаем слово как увиденное
		seen[lowerWord] = true
	}

	// Создаем результирующую мапу
	result := make(map[string][]string)
	for _, group := range anagramMap {

		// выкидываем все найденные "анаграммы" через проверку длины слайса найденных "анаграм".
		// по скольку для одной буквы всегда будет найдена только одна анаграмма, то и длина слайса будет 1
		if len(group) > 1 {
			// сортируем найденные анаграммы в порядке убывания
			sort.Strings(group)

			// добавляем рузультат
			result[group[0]] = group
		}
	}

	return result
}

// sortString сортирует символы в строке в алфавитном порядке
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "пятак", "п", "п"}
	anagramSets := findAnagrams(words)

	// выводим содержимое мапы
	for key, set := range anagramSets {
		fmt.Printf("%s: %v\n", key, set)
	}
}
