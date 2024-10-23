package main

import (
	"fmt"
	"sync"
)

func task[T string | int](i T, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(i)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}
	wg.Wait()

	// time.Sleep(time.Second * 1)
}
