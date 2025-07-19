package main

import ("fmt" 
)
type Numbers struct {
    num1 int
    num2 int
}

type NumbersInterface interface {
    Sum() int
    Multiply() int
    Division() float64
    Substract() int
}

 func (n Numbers) Sum() int {
    return n.num1 + n.num2
 }
 func (n Numbers) Multiply() int {
    return n.num1 * n.num2
 }
 func (n Numbers) Division() float64 {
    return float64(n.num1) / float64(n.num2)
 }
 func (n Numbers) Substract() int {
    return n.num1 - n.num2
 }
func main() {

var i NumbersInterface

numbers := Numbers{64546, 7576}

i = numbers

fmt.Println("Суммировать: %d\n", i.Sum())
fmt.Println("Умножить : %d\n", i.Multiply())
fmt.Println("Поделить: %d\n", i.Division())
fmt.Println("Отнять: %d\n", i.Substract())
}  