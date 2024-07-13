package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Unpack выполняет распаковку строки по описанному в задании алгоритму.
func Unpack(input string) (string, error) {
	var result []rune      // создаем слайс рун под результат
	runes := []rune(input) // представляем строку как слайс рун, чтобы правильно найти количество символов
	length := len(runes)

	for i := 0; i < length; i++ {
		current := runes[i]

		if unicode.IsDigit(current) { // если мы натыкаемся на число в строке, после переступания через предыдущее число, то это значит что формат строки не верный и надо вернуть ошибку
			return "", errors.New("invalid string format")
		}

		if current == '\\' { // наличие escape символа с последующим считается одним символом
			if i+1 < length && (unicode.IsDigit(runes[i+1]) || runes[i+1] == '\\') { // проверяем что далее идет либо чило либо еще escape последовательность
				current = runes[i+1] // если это верно, то вереходим к обработке следующего символа
				i++
			} else {
				return "", errors.New("invalid escape sequence")
			}
		}

		if i+1 < length && unicode.IsDigit(runes[i+1]) { // проверяем является ли слудующий символ - цифрой (при условии, что мы не вылезли за длинну слайса)
			count, _ := strconv.Atoi(string(runes[i+1])) // представляем слудующее найденное число в виде числа
			for j := 0; j < count; j++ {
				result = append(result, current) // добавлям текущий символ к результату count количество раз
			}
			i++ // переступаем через найденное число и переходим к следующей букве (предположительно. если это не так, то программа вывовет ошибку проверкой if unicode.IsDigit(current))
		} else {
			result = append(result, current) // добавляем найденную букву или число (если было escape) к результату, если не нашли последующее число
		}
	}

	return string(result), nil
}

func main() {
	testCases := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}

	for _, testCase := range testCases {
		unpacked, err := Unpack(testCase)
		if err != nil {
			fmt.Printf("Unpakced %q: %q (%v)\n", testCase, unpacked, err)
		} else {
			fmt.Printf("Unpacked %q: %q\n", testCase, unpacked)
		}
	}
}
