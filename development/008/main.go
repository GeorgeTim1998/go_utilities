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
	var lastCmd *exec.Cmd

	for i, cmdStr := range commands {
		parts := strings.Fields(strings.TrimSpace(cmdStr))
		if len(parts) == 0 {
			continue
		}

		cmd := exec.Command(parts[0], parts[1:]...)
		if i == 0 {
			cmd.Stdin = os.Stdin
		} else {
			cmd.Stdin, _ = lastCmd.StdoutPipe()
		}

		if i == len(commands)-1 {
			cmd.Stdout = os.Stdout
		} else {
			cmd.Stdout = os.Stdout
		}

		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		lastCmd = cmd
	}

	if lastCmd != nil {
		lastCmd.Wait()
	}
}
