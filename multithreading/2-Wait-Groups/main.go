package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}

}

/*
* Adicionar quantidade de tarefas/operações
* Informar que você terminou uma operação
* Esperar até que as operações sejam finalizadas
 */
// Thread 1
func main() {

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	// Thread 2
	go task("A", &waitGroup)

	// Thread 3
	go task("B", &waitGroup)

	// Thread
	go func() {
		for i := 1; i < 6; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
