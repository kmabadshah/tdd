package main

import (
	"errors"
	"fmt"
	"math"
)

func Greet(name string, lang string) string {
	greetPrefix := getGreetPrefix(lang)
	if name == "" {
		name = "World"
	}

	return greetPrefix + name
}

func getGreetPrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = "Hola "
	case "French":
		prefix = "Bonjour "
	default:
		prefix = "Hello "
	}
	return
}

// Add takes two integers and returns the sum of them
func Add(a, b int) int {
	return a + b
}

// returns s which contains c n times
func Repeat(c byte, n int) (s string) {
	for i := 0; i < n; i++ {
		s += string(c)
	}

	return
}

// add numbers inside slice
func Sum(s []int) (res int) {
	for _, num := range s {
		res += num
	}

	return
}

// add numbers inside arr
func SumAll(slices ...[]int) (sumSlice []int) {
	for _, slice := range slices {
		sumSlice = append(sumSlice, Sum(slice))
	}
	return
}

func SumAllTails(slices ...[]int) (sumSlice []int) {
	for _, slice := range slices {
		if len(slice) == 1 {
			sumSlice = append(sumSlice, slice[0])
		} else if len(slice) == 0 {
			sumSlice = append(sumSlice, 0)
		} else {
			sumSlice = append(sumSlice, Sum(slice[1:]))
		}
	}

	return
}

func main() {
	fmt.Println(Greet("Adnan", "Spanish"))
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Circle) Perimeter() float64 {
	return r.Radius * 2 * math.Pi
}

func (r Triangle) Perimeter() float64 {
	thirdSide := math.Sqrt(math.Pow(r.Height, 2) + math.Pow(r.Base, 2))
	return r.Height + r.Base + thirdSide
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Circle) Area() float64 {
	return math.Pi * math.Pow(r.Radius, 2)
}

func (r Triangle) Area() float64 {
	return (r.Height * r.Base) / 2
}

//Wallet
//------

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.balance {
		return errors.New("not enough balance")
	}

	w.balance -= b
	return nil
}
