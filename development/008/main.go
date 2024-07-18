package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to GoShell. Type \\quit to exit.")
	for {
		fmt.Print("> ")
		// Считываем введенную пользователем строку до символа возврата каретки.
		input, _ := reader.ReadString('\n')
		// Отрезаем лишние пробелы в начале и в конце
		input = strings.TrimSpace(input)

		if input == "\\quit" {
			break
		}

		// Разделяем по символу |
		commands := strings.Split(input, "|")
		if len(commands) > 1 {
			runPipedCommands(commands)
		} else {
			runCommand(input)
		}
	}
}

func runCommand(input string) {
	// Разбиваем строку по пробелам
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	// Берем комманду и аргументы
	cmd := parts[0]
	args := parts[1:]

	switch cmd {
	case "cd":
		if len(args) == 0 {
			fmt.Println("cd requires an argument")
			return
		}

		// Изменяет текущий рабочий каталог на каталог, указанный в path.
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Println("Error changing directory:", err)
		}
	case "pwd":
		// Возвращает текущий рабочий каталог.
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
		} else {
			fmt.Println(dir)
		}
	case "echo":
		// Выводим все аргументы разделенные пробелом
		fmt.Println(strings.Join(args, " "))
	case "kill":
		if len(args) == 0 {
			fmt.Println("kill requires a PID")
			return
		}

		// Проверяем что можно привести строку к числу PID
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid PID:", args[0])
			return
		}

		// Убиваем процесс
		err = syscall.Kill(pid, syscall.SIGKILL)
		if err != nil {
			fmt.Println("Error killing process:", err)
		}
	case "ps":
		// Этот фрагмент кода создает команду ps, перенаправляет её стандартный вывод
		// и поток ошибок в стандартные вывод и поток ошибок текущего процесса, а затем запускает команду.
		// В результате, когда команда ps выполняется, её вывод отображается в консоли текущего процесса
		ps := exec.Command("ps")
		ps.Stdout = os.Stdout
		ps.Stderr = os.Stderr
		ps.Run()
	default:
		// Выполняем external команды
		executeExternalCommand(parts)
	}
}

func executeExternalCommand(parts []string) {
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}

func runPipedCommands(commands []string) {
	// Переменная для хранения последней выполненной команды
	var lastCmd *exec.Cmd

	// Проходим по всем командам
	for i, cmdStr := range commands {
		// Разбиваем строку команды на части (команда и её аргументы)
		parts := strings.Fields(strings.TrimSpace(cmdStr))
		if len(parts) == 0 {
			// Если строка пустая, переходим к следующей команде
			continue
		}

		// Создаем новую команду
		cmd := exec.Command(parts[0], parts[1:]...)
		if i == 0 {
			// Если это первая команда, её стандартный ввод (stdin) - это стандартный ввод программы (os.Stdin)
			cmd.Stdin = os.Stdin
		} else {
			// Если это не первая команда, её стандартный ввод (stdin) - это стандартный вывод (stdout) предыдущей команды
			cmd.Stdin, _ = lastCmd.StdoutPipe()
		}

		if i == len(commands)-1 {
			// Если это последняя команда, её стандартный вывод (stdout) - это стандартный вывод программы (os.Stdout)
			cmd.Stdout = os.Stdout
		} else {
			// Если это не последняя команда, её стандартный вывод (stdout) - это стандартный вывод программы (os.Stdout)
			// Здесь ошибка: должно быть подключение к следующей команде, а не os.Stdout. Нужно изменить на `cmd.Stdout, _ = lastCmd.StdoutPipe()`
			cmd.Stdout = os.Stdout
		}

		// Стандартный поток ошибок (stderr) команды - это стандартный поток ошибок программы (os.Stderr)
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			// Если произошла ошибка при запуске команды, выводим ошибку и выходим
			fmt.Println("Error starting command:", err)
			return
		}

		// Сохраняем текущую команду как последнюю выполненную
		lastCmd = cmd
	}

	// Ожидаем завершения последней команды
	if lastCmd != nil {
		lastCmd.Wait()
	}
}
