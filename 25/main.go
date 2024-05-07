package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	fmt.Println("Старт")
	sleep(5 * time.Second)
	fmt.Println("Финиш")
}
