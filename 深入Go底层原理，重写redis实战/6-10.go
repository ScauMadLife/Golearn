package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	salary int
	level  int
	mu     sync.RWMutex
}

func (p *Person) promote() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.salary += 1000
	fmt.Println(p.salary)
	p.level += 1
	fmt.Println(p.level)
}

func (p *Person) printPerson(w *sync.WaitGroup) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Println(p.salary)
	fmt.Println(p.level)
	w.Done()
}
func main() {
	p := Person{level: 1, salary: 10000}
	wg := sync.WaitGroup{}
	wg.Add(3)
	once := sync.Once{} //sync.once一段代码只执行一次：使用标志位和Mutex实现
	go once.Do(p.promote)
	go once.Do(p.promote)
	go once.Do(p.promote)

	//wg.Wait()

	time.Sleep(time.Second)
}
