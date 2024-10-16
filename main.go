// Создайте программу, которая запускает 5 рутин, каждая из которых печатает
// свой порядковый номер 10 раз. Добиться такого результата за один вызов wg.Add.

package main

import (
	"fmt"
	"sync"
)

const numGoroutines = 5

func main() {
	var wg sync.WaitGroup

	wg.Add(numGoroutines)
	for idx := 0; idx < numGoroutines; idx++ {
		go func(idx int) {
			for j := 0; j < 10; j++ {
				fmt.Print(idx, " ")
			}
			wg.Done()
		}(idx)
	}

	wg.Wait()
}
