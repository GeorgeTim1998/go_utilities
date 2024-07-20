package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Определение флагов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Установка таймаута для подключения
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %v\n", address)
	fmt.Println("Escape is Cntr+D")

	// Канал для сигналов прерывания (Ctrl+C, Ctrl+D)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Канал для завершения программы
	doneChan := make(chan struct{})
	// Создаем объект синхронизации, чтобы удостовериться что канал doneChan закрывается только один раз
	var once sync.Once

	// Горутина для чтения из сокета и записи в STDOUT
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("Connection closed by server")
		}
		once.Do(func() { close(doneChan) })
	}()

	// Горутина для чтения из STDIN и записи в сокет
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Println("Error writing to connection:", err)
		}
		once.Do(func() { close(doneChan) })
	}()

	select {
	case <-doneChan:
	case <-sigChan:
		fmt.Println("Received interrupt signal, closing connection...")
	}

	fmt.Println("Connection closed")
}
