package error

import (
	"testing"
	"time"
)

func TestGo(t *testing.T) {

	Go(func() {
		t.Log("Go")
		panic("Hello world")
	})
	time.Sleep(time.Second)
}

func TestFailed(t *testing.T) {
	ch := make(chan bool)
	go func() {
		defer func() {
			ch <- true
		}()
		panic("Failed Go")

	}()
	<-ch

}
