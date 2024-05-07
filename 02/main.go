package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2,4,6,8,10}

	// Создаем группу ожидания для корректной отработки программы
	var wg sync.WaitGroup

	for i := range arr {
		// Увеличиваем счетчик группы на 1 на каждой итерации
		wg.Add(1)
		// Запускаем горутину с выводом квадрата элемента массива
		go func(x int) {
			// После выполнения отправляется сигнал о завершении
			defer wg.Done()
			fmt.Println(x*x)
		}(arr[i])
	}
	// Ожидаем выполнения всех горутин
	wg.Wait()
}

// func main() {
// 	arr := []int{2,4,6,8,10}

// 	var wg sync.WaitGroup

// 	for i := range arr {
// 		wg.Add(1)
// 		go func(x int, wg *sync.WaitGroup) {
// 			defer wg.Done()
// 			fmt.Println(x*x)
// 		}(arr[i], &wg)
// 	}
// 	wg.Wait()
// }