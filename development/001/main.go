package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// Получаем текущее точное время через NTP
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		// В случае ошибки выводим её в STDERR и возвращаем ненулевой код выхода
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Печатаем текущее точное время
	fmt.Printf("Current time: %s\n", ntpTime.Format(time.RFC1123))
}
