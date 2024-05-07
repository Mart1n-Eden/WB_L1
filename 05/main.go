package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	dur := flag.Int("d", 5, "время завершения программы")
	flag.Parse()

	ch := make(chan int)

	// Непрерывная запись данных в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(20 * time.Millisecond)
		}
	}()

	// Запускаем функцию чтения из канала
	go worker(ch)

	// Запускаем таймер до завершения программы
	time.Sleep(time.Duration(*dur) * time.Second)

	close(ch)
}

func worker(ch <-chan int) {
	for data := range ch {
		fmt.Println(data)
	}
}
