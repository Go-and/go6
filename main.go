package main

import (
	"fmt"
	
)

func main() {

ch := make(chan int)

 go say("Hello Go!!!!", ch)

for a := range ch {
    fmt.Println(a)
}
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
}

func say(greet string, ch chan int) {

for i := 0; i <= 5; i++ {
    fmt.Println(i)
    ch <- i
}

close(ch)
ch <- 1
}