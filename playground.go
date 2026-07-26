package main

import(
	"fmt"
	"sync"
)

func Worker(id int,wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("Worker", id, "started")

	fmt.Println("Worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go Worker(1, &wg)
	go Worker(2, &wg)

	wg.Wait()

	fmt.Println("All workers finished!")
}