package task2

import (
	"fmt"
	"sync"
	"time"
)

// 指针
func Add(num *int) {
	*num += 10
}

func Slice(arr []int) {
	for i, num := range arr {
		arr[i] = num * 2
	}

}

// goroutine
func Goroutine1() {

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("routine1", i)
			}
		}

	}()

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("routine2", i)
			}
		}
	}()

}

// 面向对象
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct{}

func (r *Rectangle) Area() float64 {

	return 180.0
}

func (r *Rectangle) Perimeter() float64 {
	return 20.0
}

type Circle struct{}

func (c *Circle) Area() float64 {
	return 0
}

func (c *Circle) Perimeter() float64 {
	return 20.0
}

func Perimeter(s Shape) float64 {

	s.Area()
	s.Perimeter()
	return s.Perimeter() * s.Perimeter()

}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) Printlnfo() {
	fmt.Println("Employee", e.Name, e.Age, e.EmployeeID)
}

func Channel1() {

	ch := make(chan int, 10)
	fmt.Println("channel1", &ch)
	go func(cnt chan int) {

		for i := 0; i < 100; i++ {
			ch <- i
		}

	}(ch)

	go func(cnt chan int) {
		for {
			i, flag := <-ch

			if flag {
				fmt.Println(i)
			}
		}

	}(ch)
}

func Lock1() {
	num := 0
	mutex := sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				num++
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(time.Second)

	fmt.Println(num)

}
