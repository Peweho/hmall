package test

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"log"
	"testing"
)

type Err struct {
}

func (e *Err) Error() string {
	return "123"
}

func TestA(t *testing.T) {
	a := 1

	v, err := mr.MapReduce[int, int, int](func(source chan<- int) {
		source <- a
	}, func(item int, writer mr.Writer[int], cancel func(error)) {
		writer.Write(item)
		//cancel(&Err{})
	}, func(pipe <-chan int, writer mr.Writer[int], cancel func(error)) {
		for i := range pipe {
			writer.Write(i)
		}
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("v:", v)
}
