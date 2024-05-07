package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0

	// Создаем группу ожидания для корректной отработки программы
	var wg sync.WaitGroup
	// Создаем канал в который будут записываться квадраты элементов
	ch := make(chan int)

	for i := range arr {
		// Увеличиваем счетчик группы на 1 на каждой итерации
		wg.Add(1)
		// Запускаем горутину в которой отправляется квадрат элемента в канал
		go func(x int) {
			// После выполнения отправляется сигнал о завершении
			defer wg.Done()
			ch <- x * x
		}(arr[i])
	}

	go func() {
		// Ожидание занесения всех значений в канал
		wg.Wait()
		// Закрытие канала, чтобы чтение из канала с помощью цикла могло завершиться
		close(ch)
	}()

	for squaredNum := range ch {
		// Получение и сложение всех данных из канала
		sum += squaredNum
	}

	fmt.Println(sum)
}

// func main() {
// 	arr := []int{2, 4, 6, 8, 10}
// 	sum := 0

// 	// Создаем мьютекс для блокировки ресурса переменной
// 	var mu sync.Mutex
// 	// Создаем группу ожидания для корректной отработки программы
// 	var wg sync.WaitGroup

// 	for i := range arr {
// 		// Увеличиваем счетчик группы на 1 на каждой итерации
// 		wg.Add(1)
// 		go func(x int) {
// 			// После выполнения отправляется сигнал о завершении
// 			defer wg.Done()

// 			// Блокируем ресурс
// 			mu.Lock()
// 			// Разблокировка произойдет после выхода из функции
// 			defer mu.Unlock()

// 			sum += x * x
// 		}(arr[i])
// 	}

// 	// Ожидаем выполнения всех горутин
// 	wg.Wait()

// 	fmt.Println(sum)
// }
