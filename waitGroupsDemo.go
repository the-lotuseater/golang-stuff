// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	fmt.Println("initializing 2 threads for Golang, this will initialize the counter to 2, the main function will wait for the counter to become 0 before terminating. The defer method reduces the counter of the waitGroup by 1.")
	waitGroup.Add(2)
	go sayHello(&waitGroup)
	go sayBye(&waitGroup)
	waitGroup.Wait()
}

func sayHello(wg *sync.WaitGroup) {
	fmt.Println("Hello!")
	defer wg.Done()
}

func sayBye(wg *sync.WaitGroup) {
	fmt.Println("Bye!!!")
	defer wg.Done()

}
