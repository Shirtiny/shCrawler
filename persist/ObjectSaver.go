package persist

import (
	"log"
)

func Saver() chan interface{} {
	//输出通道
	out := make(chan interface{})
	go func() {
		count := 0
		for {
			object := <-out

			log.Printf("worker输出的对象为：#%d %v\n", count, object)
			count++
		}
	}()
	return out
}
