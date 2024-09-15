package test

import (
	"fmt"
)

func InvalidUseCase() {
	msgCh := make(chan int)

	for {
		select {
		case msg := <-msgCh:
			fmt.Println("this should be identified by linter", msg)
		}
	}
}
