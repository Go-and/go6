package main

import (
	"fmt"
	"sync"
	
)

func main() {
var money int32   
var donationsCount int32
 
mutex := &sync.RWMutex{}


go func()  {
 for {
    mutex.RLock()
    m := money
    dc := donationsCount
    mutex.RUnlock()
 if m != dc {
    fmt.Println("money =", m, "donations =", dc)
    break
 }
    
 }
}()

wg := &sync.WaitGroup{}

wg.Add(1000)
for range 1000 {
    go func() {
        defer wg.Done()
        mutex.Lock()
        money ++
        donationsCount ++
        mutex.Unlock()
    }()
}
     wg.Wait()
 fmt.Println(money)
}


