package task2

import (
	"fmt"
	"sync/atomic"
	"time"
)

func TestCounter() {

	var num int32

	for i := 0; i < 10; i++ {

		go func() {

			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&num, 1)
			}
		}()

	}
	time.Sleep(time.Second)
	fmt.Println(num)

}
