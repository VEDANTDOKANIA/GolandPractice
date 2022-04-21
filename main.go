package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go print(&wg)
	wg.Wait()
	fmt.Println("hello from main")
}
func print(wg *sync.WaitGroup) {
	fmt.Print("hello from go routine")
	wg.Done() // decrease one count
	wg.Done()

}
